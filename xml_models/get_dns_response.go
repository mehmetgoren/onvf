package xml_models

type DNSInformation struct {
	Body struct {
		Text           string `xml:",chardata"`
		GetDNSResponse struct {
			Text           string `xml:",chardata"`
			DNSInformation struct {
				Text        string `xml:",chardata"`
				FromDHCP    string `xml:"FromDHCP"`
				DNSFromDHCP []struct {
					Text        string `xml:",chardata"`
					Type        string `xml:"Type"`
					IPv4Address string `xml:"IPv4Address"`
				} `xml:"DNSFromDHCP"`
			} `xml:"DNSInformation"`
		} `xml:"GetDNSResponse"`
	} `xml:"Body"`
}
