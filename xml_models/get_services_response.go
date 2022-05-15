package xml_models

type Services struct {
	Body struct {
		GetServicesResponse struct {
			Text    string `xml:",chardata"`
			Service []struct {
				Text      string `xml:",chardata"`
				Namespace string `xml:"Namespace"`
				XAddr     string `xml:"XAddr"`
				Version   struct {
					Text  string `xml:",chardata"`
					Major string `xml:"Major"`
					Minor string `xml:"Minor"`
				} `xml:"Version"`
			} `xml:"Service"`
		} `xml:"GetServicesResponse"`
	} `xml:"Body"`
}
