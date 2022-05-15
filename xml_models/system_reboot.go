package xml_models

type SystemReboot struct {
	Body struct {
		Text                 string `xml:",chardata"`
		SystemRebootResponse struct {
			Text    string `xml:",chardata"`
			Message string `xml:"Message"`
		} `xml:"SystemRebootResponse"`
	} `xml:"Body"`
}
