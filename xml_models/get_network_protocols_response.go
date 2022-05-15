package xml_models

type NetworkProtocol struct {
	Body struct {
		Text                        string `xml:",chardata"`
		GetNetworkProtocolsResponse struct {
			Text             string `xml:",chardata"`
			NetworkProtocols []struct {
				Text    string `xml:",chardata"`
				Name    string `xml:"Name"`
				Enabled string `xml:"Enabled"`
				Port    string `xml:"Port"`
			} `xml:"NetworkProtocols"`
		} `xml:"GetNetworkProtocolsResponse"`
	} `xml:"Body"`
}
