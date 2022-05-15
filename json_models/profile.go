package json_models

type Profile struct {
	Fixed                    string `json:"Fixed"`
	Token                    string `json:"Token"`
	Name                     string `json:"Name"`
	VideoSourceConfiguration struct {
		Token       string `json:"Token"`
		Name        string `json:"Name"`
		UseCount    string `json:"UseCount"`
		SourceToken string `json:"SourceToken"`
		Bounds      struct {
			Height string `json:"Height"`
			Width  string `json:"Width"`
			Y      string `json:"Y"`
			X      string `json:"X"`
		} `json:"Bounds"`
	} `json:"VideoSourceConfiguration"`
	AudioSourceConfiguration struct {
		Token       string `json:"Token"`
		Name        string `json:"Name"`
		UseCount    string `json:"UseCount"`
		SourceToken string `json:"SourceToken"`
	} `json:"AudioSourceConfiguration"`
	VideoEncoderConfiguration struct {
		Token      string `json:"Token"`
		Name       string `json:"Name"`
		UseCount   string `json:"UseCount"`
		Encoding   string `json:"Encoding"`
		Resolution struct {
			Width  string `json:"Width"`
			Height string `json:"Height"`
		} `json:"Resolution"`
		Quality     string `json:"Quality"`
		RateControl struct {
			FrameRateLimit   string `json:"FrameRateLimit"`
			EncodingInterval string `json:"EncodingInterval"`
			BitrateLimit     string `json:"BitrateLimit"`
		} `json:"RateControl"`
		H264 struct {
			GovLength   string `json:"GovLength"`
			H264Profile string `json:"H264Profile"`
		} `json:"H264"`
		Multicast struct {
			Address struct {
				Type        string `json:"Type"`
				IPv4Address string `json:"IPv4Address"`
			} `json:"Address"`
			Port      string `json:"Port"`
			TTL       string `json:"TTL"`
			AutoStart string `json:"AutoStart"`
		} `json:"Multicast"`
		SessionTimeout string `json:"SessionTimeout"`
	} `json:"VideoEncoderConfiguration"`
	AudioEncoderConfiguration struct {
		Token      string `json:"Token"`
		Name       string `json:"Name"`
		UseCount   string `json:"UseCount"`
		Encoding   string `json:"Encoding"`
		Bitrate    string `json:"Bitrate"`
		SampleRate string `json:"SampleRate"`
		Multicast  struct {
			Address struct {
				Type        string `json:"Type"`
				IPv4Address string `json:"IPv4Address"`
			} `json:"Address"`
			Port      string `json:"Port"`
			TTL       string `json:"TTL"`
			AutoStart string `json:"AutoStart"`
		} `json:"Multicast"`
		SessionTimeout string `json:"SessionTimeout"`
	} `json:"AudioEncoderConfiguration"`
	VideoAnalyticsConfiguration struct {
		Token                        string `json:"Token"`
		Name                         string `json:"Name"`
		UseCount                     string `json:"UseCount"`
		AnalyticsEngineConfiguration struct {
			AnalyticsModule struct {
				Name       string `json:"Name"`
				Type       string `json:"Type"`
				Parameters struct {
					SimpleItem struct {
						Name  string `json:"Name"`
						Value string `json:"Value"`
					} `json:"SimpleItem"`
					ElementItem struct {
						Name       string `json:"Name"`
						CellLayout struct {
							Columns        string `json:"Columns"`
							Rows           string `json:"Rows"`
							Transformation struct {
								Translate struct {
									X string `json:"X"`
									Y string `json:"Y"`
								} `json:"Translate"`
								Scale struct {
									X string `json:"X"`
									Y string `json:"Y"`
								} `json:"Scale"`
							} `json:"Transformation"`
						} `json:"CellLayout"`
					} `json:"ElementItem"`
				} `json:"Parameters"`
			} `json:"AnalyticsModule"`
		} `json:"AnalyticsEngineConfiguration"`
		RuleEngineConfiguration struct {
			Rule struct {
				Name       string `json:"Name"`
				Type       string `json:"Type"`
				Parameters struct {
					SimpleItem []struct {
						Name  string `json:"Name"`
						Value string `json:"Value"`
					} `json:"SimpleItem"`
				} `json:"Parameters"`
			} `json:"Rule"`
		} `json:"RuleEngineConfiguration"`
	} `json:"VideoAnalyticsConfiguration"`
	PTZConfiguration struct {
		Token                                  string `json:"Token"`
		Name                                   string `json:"Name"`
		UseCount                               string `json:"UseCount"`
		NodeToken                              string `json:"NodeToken"`
		DefaultAbsolutePantTiltPositionSpace   string `json:"DefaultAbsolutePantTiltPositionSpace"`
		DefaultAbsoluteZoomPositionSpace       string `json:"DefaultAbsoluteZoomPositionSpace"`
		DefaultRelativePanTiltTranslationSpace string `json:"DefaultRelativePanTiltTranslationSpace"`
		DefaultRelativeZoomTranslationSpace    string `json:"DefaultRelativeZoomTranslationSpace"`
		DefaultContinuousPanTiltVelocitySpace  string `json:"DefaultContinuousPanTiltVelocitySpace"`
		DefaultContinuousZoomVelocitySpace     string `json:"DefaultContinuousZoomVelocitySpace"`
		DefaultPTZSpeed                        struct {
			PanTilt struct {
				X     string `json:"X"`
				Y     string `json:"Y"`
				Space string `json:"Space"`
			} `json:"PanTilt"`
			Zoom struct {
				X     string `json:"X"`
				Space string `json:"Space"`
			} `json:"Zoom"`
		} `json:"DefaultPTZSpeed"`
		DefaultPTZTimeout string `json:"DefaultPTZTimeout"`
		PanTiltLimits     struct {
			Range struct {
				URI    string `json:"URI"`
				XRange struct {
					Min string `json:"Min"`
					Max string `json:"Max"`
				} `json:"XRange"`
				YRange struct {
					Min string `json:"Min"`
					Max string `json:"Max"`
				} `json:"YRange"`
			} `json:"Range"`
		} `json:"PanTiltLimits"`
		ZoomLimits struct {
			Range struct {
				URI    string `json:"URI"`
				XRange struct {
					Min string `json:"Min"`
					Max string `json:"Max"`
				} `json:"XRange"`
			} `json:"Range"`
		} `json:"ZoomLimits"`
	} `json:"PTZConfiguration"`
}
