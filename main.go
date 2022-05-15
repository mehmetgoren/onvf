package main

import (
	"github.com/docker/docker/client"
	"log"
	"onvf/eb"
	"onvf/hack"
	"onvf/reps"
	"onvf/utils"
)

func main() {
	defer utils.HandlePanic()

	clnt, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer func() {
		err := clnt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	dm := hack.DockerManager{Client: clnt}
	utils.RemoveContainers(&dm)
	defer func() {
		utils.RemoveContainers(&dm)
	}()

	rb := reps.RepoBucket{}
	rb.Init()

	WhoAreYou(&rb)

	oh := eb.OneHandler{Rb: &rb}
	oh.Init()

	eventBus := eb.EventBus{Rb: &rb, Channel: "onvif_request"}
	err = eventBus.Subscribe(&oh)
	if err != nil {
		log.Println(err.Error())
	}
}
