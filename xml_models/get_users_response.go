package xml_models

type Users struct {
	Body struct {
		Text             string `xml:",chardata"`
		GetUsersResponse struct {
			Text string `xml:",chardata"`
			User []struct {
				Text      string `xml:",chardata"`
				Username  string `xml:"Username"`
				UserLevel string `xml:"UserLevel"`
			} `xml:"User"`
		} `xml:"GetUsersResponse"`
	} `xml:"Body"`
}
