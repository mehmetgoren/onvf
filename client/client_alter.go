package client

import (
	"encoding/xml"
	"github.com/use-go/onvif/device"
	"onvf/xml_models"
)

func (o *OnvifClient) Reboot() (bool, error) {
	req := device.SystemReboot{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return false, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return false, err
	}
	m := &xml_models.SystemReboot{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (o *OnvifClient) SetSystemFactoryDefault() (bool, error) {
	req := device.SetSystemFactoryDefault{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return false, err
	}
	_, err = readResponse(resp)

	return err != nil, err
}

func (o *OnvifClient) StartFirmwareUpgrade() (bool, error) {
	req := device.StartFirmwareUpgrade{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return false, err
	}
	_, err = readResponse(resp)

	return err != nil, err
}
