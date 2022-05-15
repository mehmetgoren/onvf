package client

import (
	"encoding/xml"
	"github.com/use-go/onvif/media"
	"github.com/use-go/onvif/xsd/onvif"
	"onvf/json_models"
	"onvf/xml_models"
)

func (o *OnvifClient) GetProfiles() ([]*json_models.Profile, error) {
	req := media.GetProfiles{}

	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	x := &xml_models.Profiles{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}

	ret := make([]*json_models.Profile, 0)
	for _, profile := range x.Body.GetProfilesResponse.Profiles {
		item := &json_models.Profile{}
		err := convert(&profile, item)
		if err != nil {
			return nil, err
		}
		ret = append(ret, item)
	}

	return ret, nil
}

func (o *OnvifClient) GetStreamUri() (*json_models.StreamUri, error) {
	profiles, err := o.GetProfiles()
	if err != nil {
		return nil, err
	}
	if len(profiles) == 0 {
		return nil, nil
	}
	profileToken := profiles[0].Token
	req := media.GetStreamUri{ProfileToken: onvif.ReferenceToken(profileToken)}

	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	x := &xml_models.StreamUri{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}
	ret := &json_models.StreamUri{}
	err = convert(&x.Body.GetStreamUriResponse.MediaUri, ret)
	return ret, nil
}
