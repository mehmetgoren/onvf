package hack

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/lithammer/shortuuid/v4"
	"io"
	"log"
	"onvf/models"
	"os"
	"strings"
)

type DockerManager struct {
	Client *client.Client
}

var imageName = "gokalpgoren/cameradar"

func (d *DockerManager) getImage() (*types.ImageSummary, error) {
	results, err := d.Client.ImageList(context.Background(), types.ImageListOptions{All: false})
	if err != nil {
		log.Println("an error occurred while searching a docker image, err: ", err.Error())
		return nil, err
	}

	name := imageName + ":latest"
	for _, result := range results {
		for _, tag := range result.RepoTags {
			if tag == name {
				return &result, nil
			}
		}
	}

	return nil, nil
}

func (d *DockerManager) pullImage() error {
	name := imageName + ":latest"
	out, err := d.Client.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		log.Println("an error occurred while pulling the image, err: ", err.Error())
		return err
	}
	defer func(out io.ReadCloser) {
		err := out.Close()
		if err != nil {
			log.Println("an error occurred while closing pull image reader, err: ", err.Error())
		}
	}(out)
	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		log.Println("an error occurred while copying reader, err: ", err.Error())
		return err
	}

	return nil
}

func (d *DockerManager) InitImage() error {
	img, _ := d.getImage()
	if img == nil {
		err := d.pullImage()
		if err != nil {
			log.Println("an error occurred while pulling the image, err: ", err.Error())
			return err
		}
	}

	return nil
}

func (d *DockerManager) RemoveContainers() (int, error) {
	ctx := context.Background()
	containers, err := d.Client.ContainerList(ctx, types.ContainerListOptions{All: false})
	if err != nil {
		log.Println("an error occurred while getting the container, err: ", err.Error())
		return 0, err
	}

	count := 0
	startWith := "/cameradar"
	for _, cntr := range containers {
		for _, cname := range cntr.Names {
			if strings.HasPrefix(cname, startWith) {
				err := d.Client.ContainerRemove(ctx, cntr.ID, types.ContainerRemoveOptions{Force: true})
				if err != nil {
					log.Println("an error occurred while removing the container, err: ", err.Error())
				} else {
					log.Println("a container has been removed: ", cntr.Names[0])
					count++
				}
			}
		}
	}

	return count, nil
}

func (d *DockerManager) Scan(target string) ([]*models.ExecResultView, error) {
	name := "cameradar_" + shortuuid.New()[0:11]
	ctx := context.Background()

	config := container.Config{}
	config.Image = imageName
	config.OpenStdin = true
	config.AttachStderr = true
	config.AttachStdout = true
	config.AttachStdin = true
	config.Cmd = []string{"-t", target}

	cntr, err := d.Client.ContainerCreate(ctx, &config, &container.HostConfig{}, nil, nil, name)
	if err != nil {
		log.Println("an error occurred during the creating the container, err: ", err.Error())
		return nil, err
	}
	defer func() {
		err = d.Client.ContainerRemove(ctx, cntr.ID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			log.Println("an error occurred while removing the container, err: ", err.Error())
		} else {
			log.Println("a container has been removed: ", cntr.ID)
		}
	}()

	err = d.Client.ContainerStart(ctx, cntr.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Println("an error occurred during the creating exec the container, err: ", err.Error())
		return nil, err
	}

	statusCh, errCh := d.Client.ContainerWait(ctx, cntr.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Println("an error occurred during the waiting the container, err: ", err.Error())
			return nil, err
		}
	case <-statusCh:
	}

	out, err := d.Client.ContainerLogs(ctx, cntr.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		log.Println("an error occurred while getting the container' s output, err: ", err.Error())
		return nil, err
	}
	if b, err := io.ReadAll(out); err == nil {
		erv := make([]*models.ExecResult, 0)
		err := json.Unmarshal(b[8:], &erv)
		if err != nil {
			log.Println("an error occurred during the deserializing json result, err: ", err.Error())
			return nil, err
		}
		ret := make([]*models.ExecResultView, 0)
		for _, stream := range erv {
			er := &models.ExecResultView{}
			er.Username = stream.Username
			er.AuthenticationType = stream.GetAuthenticationTypeString()
			er.Device = stream.Device
			er.Password = stream.Password
			er.Port = stream.Port
			er.Address = stream.Address
			er.Available = stream.Available
			er.CredentialsFound = stream.CredentialsFound
			er.RouteFound = stream.RouteFound
			er.Routes = stream.Routes
			ret = append(ret, er)
		}
		return ret, nil
	} else {
		log.Println("an error occurred while reading the buffer, err: ", err.Error())
		return nil, err
	}
}
