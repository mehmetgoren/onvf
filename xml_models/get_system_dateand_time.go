package xml_models

type SystemDateAndTime struct {
	Body struct {
		Text                         string `xml:",chardata"`
		GetSystemDateAndTimeResponse struct {
			Text              string `xml:",chardata"`
			SystemDateAndTime struct {
				Text            string `xml:",chardata"`
				DateTimeType    string `xml:"DateTimeType"`
				DaylightSavings string `xml:"DaylightSavings"`
				TimeZone        struct {
					Text string `xml:",chardata"`
					TZ   string `xml:"TZ"`
				} `xml:"TimeZone"`
				UTCDateTime struct {
					Text string `xml:",chardata"`
					Time struct {
						Text   string `xml:",chardata"`
						Hour   string `xml:"Hour"`
						Minute string `xml:"Minute"`
						Second string `xml:"Second"`
					} `xml:"Time"`
					Date struct {
						Text  string `xml:",chardata"`
						Year  string `xml:"Year"`
						Month string `xml:"Month"`
						Day   string `xml:"Day"`
					} `xml:"Date"`
				} `xml:"UTCDateTime"`
				LocalDateTime struct {
					Text string `xml:",chardata"`
					Time struct {
						Text   string `xml:",chardata"`
						Hour   string `xml:"Hour"`
						Minute string `xml:"Minute"`
						Second string `xml:"Second"`
					} `xml:"Time"`
					Date struct {
						Text  string `xml:",chardata"`
						Year  string `xml:"Year"`
						Month string `xml:"Month"`
						Day   string `xml:"Day"`
					} `xml:"Date"`
				} `xml:"LocalDateTime"`
			} `xml:"SystemDateAndTime"`
		} `xml:"GetSystemDateAndTimeResponse"`
	} `xml:"Body"`
}
