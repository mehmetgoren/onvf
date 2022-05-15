package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/docker/docker/client"
	"log"
	onvfclnt "onvf/client"
	"onvf/hack"
)

func getConcord() *onvfclnt.OnvifClient {
	c := &onvfclnt.OnvifClient{Address: "192.168.0.23", Port: 6688, Username: "admin", Password: "admin123456"}
	err := c.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	return c
}

func getTapo() *onvfclnt.OnvifClient {
	c := &onvfclnt.OnvifClient{Address: "192.168.0.29", Port: 2020, Username: "admin12", Password: "admin12"}
	err := c.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	return c
}

func removeContainers(dm *hack.DockerManager) {
	_, err := dm.RemoveContainers()
	if err != nil {
		log.Println("an error occurred while removing a container, err: ", err.Error())
	}
}

func test() {

	//rb := reps.RepoBucket{}
	//rb.Init()
	//
	//c := getTapo() //getTapo() //getConcord()
	//oh := eb.OneHandler{Rb: &rb, Oc: c}
	//oh.Init()
	//showAsJson(oh.GetTargetInfo())

	clnt, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Println("docker client couldn't be created, err: " + err.Error())
		return
	}
	defer func() {
		err := clnt.Close()
		if err != nil {
			log.Println("an error occurred during the closing docker client, err: ", err.Error())
			return
		}
	}()
	dm := hack.DockerManager{Client: clnt}
	removeContainers(&dm)
	defer func() {
		removeContainers(&dm)
	}()

	err = dm.InitImage()
	if err != nil {
		log.Println("an error occurred while initializing the docker client, the program is now exiting, err: ", err.Error())
		return
	}
	result, _ := dm.Scan("192.168.0.1/24")
	js, _ := json.Marshal(&result)
	fmt.Println(string(js))

	//fmt.Println("\n")
	//capabilities, _ := c.GetCapabilities()
	//showAsJson(capabilities)

	//fmt.Println("\n")
	//dv, _ := c.GetDeviceInfo()
	//showAsJson(dv)

	//fmt.Println("\n")
	//s, _ := c.GetServices()
	//showArrAsJson(s)

	//fmt.Println("\n")
	//w, _ := c.IsDiscoverable()
	//fmt.Println(w)

	//fmt.Println("\n")
	//dns, _ := c.GetDNSInfo()
	//showAsJson(dns)

	//fmt.Println("\n")
	//ep, _ := c.GetEndpointReference()
	//fmt.Println(ep)

	//fmt.Println("\n")
	//host, _ := c.GetHostname()
	//showAsJson(host)

	//fmt.Println("\n")
	//dng, _ := c.GetNetworkGateway()
	//fmt.Println(dng)

	//fmt.Println("\n")
	//ni, _ := c.GetNetworkInterfaces()
	//showAsJson(ni)

	//fmt.Println("\n")
	//np, _ := c.GetNetworkProtocols()
	//showArrAsJson(np)

	//fmt.Println("\n")
	//sc, _ := c.GetScopes()
	//showArrAsJson(sc)

	//fmt.Println("\n")
	//scap, _ := c.GetServiceCapabilities()
	//showAsJson(scap)

	//fmt.Println("\n")
	//dt, _ := c.GetSystemDateAndTime()
	//showAsJson(dt)

	//fmt.Println("\n")
	//l, _ := c.GetSystemLog()
	//fmt.Println(l)

	//fmt.Println("\n")
	//u, _ := c.GetUsers()
	//showArrAsJson(u)

	//fmt.Println("\n")
	//p, _ := c.GetProfiles()
	//showArrAsJson(p)

	//fmt.Println("\n")
	//uri, _ := c.GetStreamUri()
	//showAsJson(uri)
}

func showAsJson[T any](model *T) {
	js, _ := json.Marshal(model)
	s := string(js)
	fmt.Println(s)
}

func showArrAsJson[T any](model []*T) {
	js, _ := json.Marshal(model)
	s := string(js)
	fmt.Println(s)
}

func showAsXml[T any](model *T) {
	js, _ := xml.Marshal(model)
	s := string(js)
	fmt.Println(s)
}
