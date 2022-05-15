package eb

import (
	"encoding/json"
	"github.com/docker/docker/client"
	"log"
	onvfclnt "onvf/client"
	"onvf/hack"
	"onvf/json_models"
	"onvf/models"
	"onvf/reps"
	"onvf/utils"
	"strconv"
	"strings"
)

func createOnvifClient(js string) (*models.OnvifParams, *onvfclnt.OnvifClient) {
	params := &models.OnvifParams{}
	err := json.Unmarshal([]byte(js), params)
	if err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	oc := &onvfclnt.OnvifClient{Address: params.Address, Port: params.Port, Username: params.Username, Password: params.Password}
	err = oc.Init()
	if err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	return params, oc
}

func getTargetInfo(rb *reps.RepoBucket, js string) interface{} {
	cam, c := createOnvifClient(js)
	if c == nil {
		return nil
	}
	t := &json_models.TargetInfo{}
	t.DeviceInfo, _ = c.GetDeviceInfo()
	t.IsDiscoverable, _ = c.IsDiscoverable()
	hostName, err := c.GetHostname()
	if err == nil {
		t.HostName = hostName.Name
	}

	di, _ := c.GetDNSInfo()
	t.IPAddresses = make([]string, 0)
	if di != nil && di.DNSFromDHCP != nil && len(di.DNSFromDHCP) > 0 {
		for _, d := range di.DNSFromDHCP {
			if len(d.IPv4Address) > 0 {
				t.IPAddresses = append(t.IPAddresses, d.IPv4Address)
			}
		}
	}
	ni, _ := c.GetNetworkInterfaces()
	if ni != nil {
		if len(ni.IPv4.Config.FromDHCP.Address) > 0 {
			t.IPAddresses = append(t.IPAddresses, ni.IPv4.Config.FromDHCP.Address)
		}
		if len(ni.Info.HwAddress) > 0 {
			t.HwAddress = ni.Info.HwAddress
		}
	}

	np, err := c.GetNetworkProtocols()
	if err == nil {
		for _, j := range np {
			switch j.Name {
			case "HTTP":
				t.HttpPort = j.Port
				break
			case "HTTPS":
				t.HttpsPort = j.Port
				break
			case "RTSP":
				t.RtspPort = j.Port
				break
			}
		}
	}

	dt, err := c.GetSystemDateAndTime()
	if err == nil {
		sb := strings.Builder{}
		l := dt.LocalDateTime
		sb.WriteString(strconv.Itoa(l.Date.Year))
		sb.WriteString("_")
		sb.WriteString(strconv.Itoa(l.Date.Month))
		sb.WriteString("_")
		sb.WriteString(strconv.Itoa(l.Date.Day))
		sb.WriteString("_")
		sb.WriteString(strconv.Itoa(l.Time.Hour))
		sb.WriteString("_")
		sb.WriteString(strconv.Itoa(l.Time.Minute))
		sb.WriteString("_")
		sb.WriteString(strconv.Itoa(l.Time.Second))
		t.LocalDatetime = sb.String()
	}

	t.Logs, _ = c.GetSystemLog()

	t.Users, _ = c.GetUsers()
	su, err := c.GetStreamUri()
	if err == nil {
		t.StreamUri = su.URI
	}

	or := reps.OnvifRepository{Client: rb.GetMainConnection()}
	or.AddWithOnvif(cam, t)

	return t
}

func reboot(rb *reps.RepoBucket, js string) interface{} {
	_, c := createOnvifClient(js)
	if c == nil {
		return nil
	}
	result, _ := c.Reboot()
	return result
}

func setSystemFactoryDefault(rb *reps.RepoBucket, js string) interface{} {
	_, c := createOnvifClient(js)
	if c == nil {
		return nil
	}
	result, _ := c.SetSystemFactoryDefault()
	return result
}

func startFirmwareUpgrade(rb *reps.RepoBucket, js string) interface{} {
	_, c := createOnvifClient(js)
	if c == nil {
		return nil
	}
	result, _ := c.StartFirmwareUpgrade()
	return result
}

func scan(target string) []*models.ExecResultView {
	clnt, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer func() {
		err := clnt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	dm := hack.DockerManager{Client: clnt}
	err = dm.InitImage()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	result, _ := dm.Scan(target)
	return result
}

func networkDiscovery(rb *reps.RepoBucket, js string) interface{} {
	addr, _ := utils.GetIPAddress()
	if len(addr) == 0 {
		addr = "127.0.0.1"
	}
	results := scan(addr + "/24")
	//file, _ := ioutil.ReadFile("/home/gokalp/Documents/scan_json_results/2.json")
	//results := make([]*models.ExecResultView, 0)
	//json.Unmarshal(file, &results)
	if results != nil {
		ndr := reps.NetworkDiscoveryRepository{Client: rb.GetMainConnection()}
		_, err := ndr.Add(results)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return results
}

func hackTarget(rb *reps.RepoBucket, js string) interface{} {
	params := &models.OnvifParams{}
	err := json.Unmarshal([]byte(js), params)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	results := scan(params.Address)
	//file, _ := ioutil.ReadFile("/home/gokalp/Documents/scan_json_results/1.json")
	//results := make([]*models.ExecResultView, 0)
	//json.Unmarshal(file, &results)
	if len(results) > 0 {
		or := reps.OnvifRepository{Client: rb.GetMainConnection()}
		or.AddWithExecResult(params, results[0])
		return results[0]
	}
	return nil
}
