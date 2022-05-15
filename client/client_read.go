package client

import (
	"encoding/json"
	"encoding/xml"
	"github.com/use-go/onvif/device"
	"io/ioutil"
	"net/http"
	"onvf/json_models"
	"onvf/xml_models"
)

func (o *OnvifClient) GetCapabilities() (*json_models.Capabilities, error) {
	req := device.GetCapabilities{Category: "All"}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	c := &xml_models.Capabilities{}
	err = xml.Unmarshal(r, c)
	if err != nil {
		return nil, err
	}

	temp := c.Body.GetCapabilitiesResponse.Capabilities
	js, err := json.Marshal(temp)
	if err != nil {
		return nil, err
	}

	ret := &json_models.Capabilities{}
	err = json.Unmarshal(js, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OnvifClient) GetDeviceInfo() (*json_models.DeviceInfo, error) {
	req := device.GetDeviceInformation{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}

	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	d := &xml_models.DeviceInfo{}
	err = xml.Unmarshal(r, d)
	if err != nil {
		return nil, err
	}

	temp := d.Body.GetDeviceInformationResponse
	ret := &json_models.DeviceInfo{
		Manufacturer:    temp.Manufacturer,
		Model:           temp.Model,
		FirmwareVersion: temp.FirmwareVersion,
		SerialNumber:    temp.SerialNumber,
		HardwareId:      temp.HardwareId,
	}

	return ret, nil
}

func (o *OnvifClient) GetServices() ([]*json_models.Service, error) {
	req := device.GetServices{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}

	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	services := &xml_models.Services{}
	err = xml.Unmarshal(r, services)
	if err != nil {
		return nil, err
	}

	ret := make([]*json_models.Service, 0)
	for _, service := range services.Body.GetServicesResponse.Service {
		item := &json_models.Service{}
		err = convert(&service, item)
		if err != nil {
			return nil, err
		}
		ret = append(ret, item)
	}

	return ret, nil
}

func (o *OnvifClient) IsDiscoverable() (bool, error) {
	req := device.GetDiscoveryMode{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return false, err
	}

	r, err := readResponse(resp)
	if err != nil {
		return false, err
	}

	d := &xml_models.DiscoverMode{}
	err = xml.Unmarshal(r, d)
	if err != nil {
		return false, err
	}

	return d.Body.GetDiscoveryModeResponse.DiscoveryMode == "Discoverable", nil
}

func (o *OnvifClient) GetDNSInfo() (*json_models.DNS, error) {
	req := device.GetDNS{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	d := &xml_models.DNSInformation{}
	err = xml.Unmarshal(r, d)
	if err != nil {
		return nil, err
	}

	temp := d.Body.GetDNSResponse.DNSInformation
	items := make([]json_models.DNSFromDHCP, 0)
	for _, dhcp := range temp.DNSFromDHCP {
		item := json_models.DNSFromDHCP{}
		item.Type, item.IPv4Address = dhcp.Type, dhcp.IPv4Address
		items = append(items, item)
	}
	ret := &json_models.DNS{
		FromDHCP:    temp.FromDHCP,
		DNSFromDHCP: items,
	}

	return ret, nil
}

func (o *OnvifClient) GetEndpointReference() (string, error) {
	req := device.GetEndpointReference{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return "", err
	}
	r, err := readResponse(resp)
	if err != nil {
		return "", err
	}
	m := &xml_models.EndpointReference{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return "", err
	}
	return m.Body.GetEndpointReferenceResponse.GUID, nil
}

func (o *OnvifClient) GetHostname() (*json_models.Hostname, error) {
	req := device.GetHostname{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	m := &xml_models.HostName{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return nil, err
	}
	ret := &json_models.Hostname{}
	ret.FromDHCP = m.Body.GetHostnameResponse.HostnameInformation.FromDHCP == "true"
	ret.Name = m.Body.GetHostnameResponse.HostnameInformation.Name
	return ret, nil
}

func (o *OnvifClient) GetNetworkGateway() (string, error) {
	req := device.GetNetworkDefaultGateway{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return "", err
	}
	r, err := readResponse(resp)
	if err != nil {
		return "", err
	}
	m := &xml_models.NetworkDefaultGateway{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return "", err
	}

	return m.Body.GetNetworkDefaultGatewayResponse.NetworkGateway, nil
}

func (o *OnvifClient) GetNetworkInterfaces() (*json_models.NetworkInterface, error) {
	req := device.GetNetworkInterfaces{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	m := &xml_models.NetworkInterfaces{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return nil, err
	}

	ret := &json_models.NetworkInterface{}
	err = convert(&m.Body.GetNetworkInterfacesResponse.NetworkInterfaces, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OnvifClient) GetNetworkProtocols() ([]*json_models.NetworkProtocol, error) {
	req := device.GetNetworkProtocols{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	m := &xml_models.NetworkProtocol{}
	err = xml.Unmarshal(r, m)
	if err != nil {
		return nil, err
	}

	ret := make([]*json_models.NetworkProtocol, 0)
	for _, item := range m.Body.GetNetworkProtocolsResponse.NetworkProtocols {
		n := &json_models.NetworkProtocol{}
		err = convert(&item, n)
		if err != nil {
			return nil, err
		}
		ret = append(ret, n)
	}
	return ret, nil
}

func (o *OnvifClient) GetScopes() ([]*json_models.Scope, error) {
	req := device.GetScopes{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	x := &xml_models.Scope{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}
	items := make([]*json_models.Scope, 0)
	for _, item := range x.Body.GetScopesResponse.Scopes {
		s := &json_models.Scope{}
		err = convert(&item, s)
		if err != nil {
			return nil, err
		}
		items = append(items, s)
	}

	return items, nil
}

func (o *OnvifClient) GetServiceCapabilities() (*json_models.ServiceCapabilities, error) {
	req := device.GetServiceCapabilities{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	x := &xml_models.ServiceCapabilities{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}

	ret := &json_models.ServiceCapabilities{}
	err = convert(&x.Body.GetServiceCapabilitiesResponse.Capabilities, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OnvifClient) GetSystemDateAndTime() (*json_models.SystemDateAndTime, error) {
	req := device.GetSystemDateAndTime{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}
	x := &xml_models.SystemDateAndTime{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}
	ret := &json_models.SystemDateAndTime{}
	err = convert(&x.Body.GetSystemDateAndTimeResponse.SystemDateAndTime, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (o *OnvifClient) GetSystemLog() (string, error) {
	req := device.GetSystemLog{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return "", err
	}
	r, err := readResponse(resp)
	if err != nil {
		return "", err
	}
	x := &xml_models.SystemLog{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return "", err
	}

	return x.Body.GetSystemLogResponse.SystemLog.String, nil
}

func (o *OnvifClient) GetUsers() ([]*json_models.User, error) {
	req := device.GetUsers{}
	resp, err := o.d.CallMethod(req)
	if err != nil {
		return nil, err
	}
	r, err := readResponse(resp)
	if err != nil {
		return nil, err
	}

	x := &xml_models.Users{}
	err = xml.Unmarshal(r, x)
	if err != nil {
		return nil, err
	}

	items := make([]*json_models.User, 0)
	for _, i := range x.Body.GetUsersResponse.User {
		item := &json_models.User{}
		item.UserLevel = i.UserLevel
		item.Username = i.Username
		items = append(items, item)
	}

	return items, nil
}

func convert[TSource any, TDest any](source *TSource, dest *TDest) error {
	js, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(js, dest)
	if err != nil {
		return err
	}
	return nil
}

func readResponse(resp *http.Response) ([]byte, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
