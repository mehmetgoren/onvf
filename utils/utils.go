package utils

import (
	"fmt"
	"log"
	"net"
	"onvf/hack"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println("RECOVER", r)
		debug.PrintStack()
	}
}

var sep = "_"

func TimeToString(time time.Time, includeNanoSec bool) string {
	arr := make([]string, 0)
	arr = append(arr, strconv.Itoa(time.Year()))
	arr = append(arr, fixZero(int(time.Month())))
	arr = append(arr, fixZero(time.Day()))
	arr = append(arr, fixZero(time.Hour()))
	arr = append(arr, fixZero(time.Minute()))
	arr = append(arr, fixZero(time.Second()))
	if includeNanoSec {
		arr = append(arr, fixZero(time.Nanosecond()))
	}

	return strings.Join(arr, sep)
}

func fixZero(val int) string {
	if val < 9 {
		return "0" + strconv.Itoa(val)
	}
	return strconv.Itoa(val)
}

func RemoveContainers(dm *hack.DockerManager) {
	_, err := dm.RemoveContainers()
	if err != nil {
		log.Println(err.Error())
	}
}

func GetIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			s, _ := ipnet.Mask.Size()
			if s == 24 {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", nil
}

var re = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

func ParseIp(address string) string {
	return re.FindString(address)
}
