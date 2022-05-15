package xml_models

type NetworkDefaultGateway struct {
	Body struct {
		Text                             string `xml:",chardata"`
		GetNetworkDefaultGatewayResponse struct {
			Text           string `xml:",chardata"`
			NetworkGateway string `xml:"NetworkGateway"`
		} `xml:"GetNetworkDefaultGatewayResponse"`
	} `xml:"Body"`
}
