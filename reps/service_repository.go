package reps

import (
	"context"
	"github.com/go-redis/redis/v8"
	"onvf/utils"
	"os"
	"runtime"
	"time"
)

type InstanceType int

const (
	systemd   InstanceType = 0
	container InstanceType = 1
)

type ServiceModel struct {
	Name            string       `json:"name" redis:"name"`
	Description     string       `json:"description" redis:"description"`
	Platform        string       `json:"platform" redis:"platform"`
	PlatformVersion string       `json:"platform_version" redis:"platform_version"`
	HostName        string       `json:"hostname" redis:"hostname"`
	IpAddress       string       `json:"ip_address" redis:"ip_address"`
	MacAddress      string       `json:"mac_address" redis:"mac_address"`
	Processor       string       `json:"processor" redis:"processor"`
	CpuCount        int          `json:"cpu_count" redis:"cpu_count"`
	Ram             string       `json:"ram" redis:"ram"`
	Pid             int          `json:"pid" redis:"pid"`
	InstanceType    InstanceType `json:"instance_type" redis:"instance_type"`
	InstanceName    string       `json:"instance_name" redis:"instance_name"`
	CreatedAt       string       `json:"created_at" redis:"created_at"`
	Heartbeat       string       `json:"heartbeat" redis:"heartbeat"`
}

type ServiceRepository struct {
	Client *redis.Client
}

func getKey(serviceName string) string {
	return "services:" + serviceName
}

func (r *ServiceRepository) Add(serviceName string) (int64, error) {
	now := time.Now()
	host, _ := os.Hostname()
	sm := ServiceModel{
		Name:            serviceName,
		Description:     "The Onvif and Hacking Service®",
		Platform:        runtime.GOOS,
		PlatformVersion: runtime.GOARCH,
		HostName:        host,
		IpAddress:       "127.0.0.1",
		MacAddress:      "00:00:00:00:00:00",
		Processor:       "unknown",
		CpuCount:        runtime.NumCPU(),
		Ram:             "unknown",
		Pid:             os.Getpid(),
		InstanceType:    container,
		InstanceName:    "onvf-instance",
		CreatedAt:       utils.TimeToString(now, true),
		Heartbeat:       "",
	}

	return r.Client.HSet(context.Background(), getKey(serviceName), Map(sm)).Result()
}
