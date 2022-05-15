package xml_models

type Profiles struct {
	Body struct {
		Text                string `xml:",chardata"`
		GetProfilesResponse struct {
			Text     string `xml:",chardata"`
			Profiles []struct {
				Text                     string `xml:",chardata"`
				Fixed                    string `xml:"fixed,attr"`
				Token                    string `xml:"token,attr"`
				Name                     string `xml:"Name"`
				VideoSourceConfiguration struct {
					Text        string `xml:",chardata"`
					Token       string `xml:"token,attr"`
					Name        string `xml:"Name"`
					UseCount    string `xml:"UseCount"`
					SourceToken string `xml:"SourceToken"`
					Bounds      struct {
						Text   string `xml:",chardata"`
						Height string `xml:"height,attr"`
						Width  string `xml:"width,attr"`
						Y      string `xml:"y,attr"`
						X      string `xml:"x,attr"`
					} `xml:"Bounds"`
				} `xml:"VideoSourceConfiguration"`
				AudioSourceConfiguration struct {
					Text        string `xml:",chardata"`
					Token       string `xml:"token,attr"`
					Name        string `xml:"Name"`
					UseCount    string `xml:"UseCount"`
					SourceToken string `xml:"SourceToken"`
				} `xml:"AudioSourceConfiguration"`
				VideoEncoderConfiguration struct {
					Text       string `xml:",chardata"`
					Token      string `xml:"token,attr"`
					Name       string `xml:"Name"`
					UseCount   string `xml:"UseCount"`
					Encoding   string `xml:"Encoding"`
					Resolution struct {
						Text   string `xml:",chardata"`
						Width  string `xml:"Width"`
						Height string `xml:"Height"`
					} `xml:"Resolution"`
					Quality     string `xml:"Quality"`
					RateControl struct {
						Text             string `xml:",chardata"`
						FrameRateLimit   string `xml:"FrameRateLimit"`
						EncodingInterval string `xml:"EncodingInterval"`
						BitrateLimit     string `xml:"BitrateLimit"`
					} `xml:"RateControl"`
					H264 struct {
						Text        string `xml:",chardata"`
						GovLength   string `xml:"GovLength"`
						H264Profile string `xml:"H264Profile"`
					} `xml:"H264"`
					Multicast struct {
						Text    string `xml:",chardata"`
						Address struct {
							Text        string `xml:",chardata"`
							Type        string `xml:"Type"`
							IPv4Address string `xml:"IPv4Address"`
						} `xml:"Address"`
						Port      string `xml:"Port"`
						TTL       string `xml:"TTL"`
						AutoStart string `xml:"AutoStart"`
					} `xml:"Multicast"`
					SessionTimeout string `xml:"SessionTimeout"`
				} `xml:"VideoEncoderConfiguration"`
				AudioEncoderConfiguration struct {
					Text       string `xml:",chardata"`
					Token      string `xml:"token,attr"`
					Name       string `xml:"Name"`
					UseCount   string `xml:"UseCount"`
					Encoding   string `xml:"Encoding"`
					Bitrate    string `xml:"Bitrate"`
					SampleRate string `xml:"SampleRate"`
					Multicast  struct {
						Text    string `xml:",chardata"`
						Address struct {
							Text        string `xml:",chardata"`
							Type        string `xml:"Type"`
							IPv4Address string `xml:"IPv4Address"`
						} `xml:"Address"`
						Port      string `xml:"Port"`
						TTL       string `xml:"TTL"`
						AutoStart string `xml:"AutoStart"`
					} `xml:"Multicast"`
					SessionTimeout string `xml:"SessionTimeout"`
				} `xml:"AudioEncoderConfiguration"`
				VideoAnalyticsConfiguration struct {
					Text                         string `xml:",chardata"`
					Token                        string `xml:"token,attr"`
					Name                         string `xml:"Name"`
					UseCount                     string `xml:"UseCount"`
					AnalyticsEngineConfiguration struct {
						Text            string `xml:",chardata"`
						AnalyticsModule struct {
							Text       string `xml:",chardata"`
							Name       string `xml:"Name,attr"`
							Type       string `xml:"Type,attr"`
							Parameters struct {
								Text       string `xml:",chardata"`
								SimpleItem struct {
									Text  string `xml:",chardata"`
									Name  string `xml:"Name,attr"`
									Value string `xml:"Value,attr"`
								} `xml:"SimpleItem"`
								ElementItem struct {
									Text       string `xml:",chardata"`
									Name       string `xml:"Name,attr"`
									CellLayout struct {
										Text           string `xml:",chardata"`
										Columns        string `xml:"Columns,attr"`
										Rows           string `xml:"Rows,attr"`
										Transformation struct {
											Text      string `xml:",chardata"`
											Translate struct {
												Text string `xml:",chardata"`
												X    string `xml:"x,attr"`
												Y    string `xml:"y,attr"`
											} `xml:"Translate"`
											Scale struct {
												Text string `xml:",chardata"`
												X    string `xml:"x,attr"`
												Y    string `xml:"y,attr"`
											} `xml:"Scale"`
										} `xml:"Transformation"`
									} `xml:"CellLayout"`
								} `xml:"ElementItem"`
							} `xml:"Parameters"`
						} `xml:"AnalyticsModule"`
					} `xml:"AnalyticsEngineConfiguration"`
					RuleEngineConfiguration struct {
						Text string `xml:",chardata"`
						Rule struct {
							Text       string `xml:",chardata"`
							Name       string `xml:"Name,attr"`
							Type       string `xml:"Type,attr"`
							Parameters struct {
								Text       string `xml:",chardata"`
								SimpleItem []struct {
									Text  string `xml:",chardata"`
									Name  string `xml:"Name,attr"`
									Value string `xml:"Value,attr"`
								} `xml:"SimpleItem"`
							} `xml:"Parameters"`
						} `xml:"Rule"`
					} `xml:"RuleEngineConfiguration"`
				} `xml:"VideoAnalyticsConfiguration"`
				PTZConfiguration struct {
					Text                                   string `xml:",chardata"`
					Token                                  string `xml:"token,attr"`
					Name                                   string `xml:"Name"`
					UseCount                               string `xml:"UseCount"`
					NodeToken                              string `xml:"NodeToken"`
					DefaultAbsolutePantTiltPositionSpace   string `xml:"DefaultAbsolutePantTiltPositionSpace"`
					DefaultAbsoluteZoomPositionSpace       string `xml:"DefaultAbsoluteZoomPositionSpace"`
					DefaultRelativePanTiltTranslationSpace string `xml:"DefaultRelativePanTiltTranslationSpace"`
					DefaultRelativeZoomTranslationSpace    string `xml:"DefaultRelativeZoomTranslationSpace"`
					DefaultContinuousPanTiltVelocitySpace  string `xml:"DefaultContinuousPanTiltVelocitySpace"`
					DefaultContinuousZoomVelocitySpace     string `xml:"DefaultContinuousZoomVelocitySpace"`
					DefaultPTZSpeed                        struct {
						Text    string `xml:",chardata"`
						PanTilt struct {
							Text  string `xml:",chardata"`
							X     string `xml:"x,attr"`
							Y     string `xml:"y,attr"`
							Space string `xml:"space,attr"`
						} `xml:"PanTilt"`
						Zoom struct {
							Text  string `xml:",chardata"`
							X     string `xml:"x,attr"`
							Space string `xml:"space,attr"`
						} `xml:"Zoom"`
					} `xml:"DefaultPTZSpeed"`
					DefaultPTZTimeout string `xml:"DefaultPTZTimeout"`
					PanTiltLimits     struct {
						Text  string `xml:",chardata"`
						Range struct {
							Text   string `xml:",chardata"`
							URI    string `xml:"URI"`
							XRange struct {
								Text string `xml:",chardata"`
								Min  string `xml:"Min"`
								Max  string `xml:"Max"`
							} `xml:"XRange"`
							YRange struct {
								Text string `xml:",chardata"`
								Min  string `xml:"Min"`
								Max  string `xml:"Max"`
							} `xml:"YRange"`
						} `xml:"Range"`
					} `xml:"PanTiltLimits"`
					ZoomLimits struct {
						Text  string `xml:",chardata"`
						Range struct {
							Text   string `xml:",chardata"`
							URI    string `xml:"URI"`
							XRange struct {
								Text string `xml:",chardata"`
								Min  string `xml:"Min"`
								Max  string `xml:"Max"`
							} `xml:"XRange"`
						} `xml:"Range"`
					} `xml:"ZoomLimits"`
				} `xml:"PTZConfiguration"`
			} `xml:"Profiles"`
		} `xml:"GetProfilesResponse"`
	} `xml:"Body"`
}
