package xml_models

type NetworkInterfaces struct {
	Body struct {
		Text                         string `xml:",chardata"`
		GetNetworkInterfacesResponse struct {
			Text              string `xml:",chardata"`
			NetworkInterfaces struct {
				Text    string `xml:",chardata"`
				Token   string `xml:"token,attr"`
				Enabled string `xml:"Enabled"`
				Info    struct {
					Text      string `xml:",chardata"`
					Name      string `xml:"Name"`
					HwAddress string `xml:"HwAddress"`
					MTU       string `xml:"MTU"`
				} `xml:"Info"`
				IPv4 struct {
					Text    string `xml:",chardata"`
					Enabled string `xml:"Enabled"`
					Config  struct {
						Text     string `xml:",chardata"`
						FromDHCP struct {
							Text         string `xml:",chardata"`
							Address      string `xml:"Address"`
							PrefixLength string `xml:"PrefixLength"`
						} `xml:"FromDHCP"`
						DHCP string `xml:"DHCP"`
					} `xml:"Config"`
				} `xml:"IPv4"`
			} `xml:"NetworkInterfaces"`
		} `xml:"GetNetworkInterfacesResponse"`
	} `xml:"Body"`
}
