package xml_models

type StreamUri struct {
	Body struct {
		Text                 string `xml:",chardata"`
		GetStreamUriResponse struct {
			Text     string `xml:",chardata"`
			MediaUri struct {
				Text                string `xml:",chardata"`
				URI                 string `xml:"Uri"`
				InvalidAfterConnect string `xml:"InvalidAfterConnect"`
				InvalidAfterReboot  string `xml:"InvalidAfterReboot"`
				Timeout             string `xml:"Timeout"`
			} `xml:"MediaUri"`
		} `xml:"GetStreamUriResponse"`
	} `xml:"Body"`
}
