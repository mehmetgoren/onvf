package xml_models

type DiscoverMode struct {
	Body struct {
		Text                     string `xml:",chardata"`
		GetDiscoveryModeResponse struct {
			Text          string `xml:",chardata"`
			DiscoveryMode string `xml:"DiscoveryMode"`
		} `xml:"GetDiscoveryModeResponse"`
	} `xml:"Body"`
}
