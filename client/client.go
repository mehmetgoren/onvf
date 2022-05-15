package client

import (
	"github.com/use-go/onvif"
	"strconv"
)

type OnvifClient struct {
	Address  string
	Port     int
	Username string
	Password string
	d        *onvif.Device
}

func (o *OnvifClient) Init() error {
	addr := o.Address
	if o.Port > 0 {
		addr += ":" + strconv.Itoa(o.Port)
	}
	d, err := onvif.NewDevice(addr)
	if err != nil {
		return err
	}
	o.d = d
	d.Authenticate(o.Username, o.Password)

	return nil
}
