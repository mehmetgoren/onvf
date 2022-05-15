package xml_models

type SystemLog struct {
	Body struct {
		Text                 string `xml:",chardata"`
		GetSystemLogResponse struct {
			Text      string `xml:",chardata"`
			SystemLog struct {
				Text   string `xml:",chardata"`
				String string `xml:"String"`
			} `xml:"SystemLog"`
		} `xml:"GetSystemLogResponse"`
	} `xml:"Body"`
}
