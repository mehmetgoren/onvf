package xml_models

type DeviceInfo struct {
	Body struct {
		Text                         string `xml:",chardata"`
		GetDeviceInformationResponse struct {
			Text            string `xml:",chardata"`
			Manufacturer    string `xml:"Manufacturer"`
			Model           string `xml:"Model"`
			FirmwareVersion string `xml:"FirmwareVersion"`
			SerialNumber    string `xml:"SerialNumber"`
			HardwareId      string `xml:"HardwareId"`
		} `xml:"GetDeviceInformationResponse"`
	} `xml:"Body"`
}
