package xml_models

type EndpointReference struct {
	Body struct {
		Text                         string `xml:",chardata"`
		GetEndpointReferenceResponse struct {
			Text string `xml:",chardata"`
			GUID string `xml:"GUID"`
		} `xml:"GetEndpointReferenceResponse"`
	} `xml:"Body"`
}
