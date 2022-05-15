package json_models

type SystemDateAndTime struct {
	DateTimeType    string `json:"DateTimeType"`
	DaylightSavings bool   `json:"DaylightSavings,string"`
	TimeZone        struct {
		Tz string `json:"TZ"`
	} `json:"TimeZone"`
	UTCDateTime struct {
		Time struct {
			Hour   int `json:"Hour,string"`
			Minute int `json:"Minute,string"`
			Second int `json:"Second,string"`
		} `json:"Time"`
		Date struct {
			Year  int `json:"Year,string"`
			Month int `json:"Month,string"`
			Day   int `json:"Day,string"`
		} `json:"Date"`
	} `json:"UTCDateTime"`
	LocalDateTime struct {
		Time struct {
			Hour   int `json:"Hour,string"`
			Minute int `json:"Minute,string"`
			Second int `json:"Second,string"`
		} `json:"Time"`
		Date struct {
			Year  int `json:"Year,string"`
			Month int `json:"Month,string"`
			Day   int `json:"Day,string"`
		} `json:"Date"`
	} `json:"LocalDateTime"`
}
