package xml_models

type HostName struct {
	Body struct {
		Text                string `xml:",chardata"`
		GetHostnameResponse struct {
			Text                string `xml:",chardata"`
			HostnameInformation struct {
				Text     string `xml:",chardata"`
				FromDHCP string `xml:"FromDHCP"`
				Name     string `xml:"Name"`
			} `xml:"HostnameInformation"`
		} `xml:"GetHostnameResponse"`
	} `xml:"Body"`
}
