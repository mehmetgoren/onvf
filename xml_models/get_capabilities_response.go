package xml_models

type Capabilities struct {
	Body struct {
		Text                    string `xml:",chardata"`
		GetCapabilitiesResponse struct {
			Text         string `xml:",chardata"`
			Capabilities struct {
				Text      string `xml:",chardata"`
				Analytics struct {
					Text                   string `xml:",chardata"`
					XAddr                  string `xml:"XAddr"`
					RuleSupport            string `xml:"RuleSupport"`
					AnalyticsModuleSupport string `xml:"AnalyticsModuleSupport"`
				} `xml:"Analytics"`
				Device struct {
					Text    string `xml:",chardata"`
					XAddr   string `xml:"XAddr"`
					Network struct {
						Text              string `xml:",chardata"`
						IPFilter          string `xml:"IPFilter"`
						ZeroConfiguration string `xml:"ZeroConfiguration"`
						IPVersion6        string `xml:"IPVersion6"`
						DynDNS            string `xml:"DynDNS"`
						Extension         struct {
							Text               string `xml:",chardata"`
							Dot11Configuration string `xml:"Dot11Configuration"`
						} `xml:"Extension"`
					} `xml:"Network"`
					System struct {
						Text              string `xml:",chardata"`
						DiscoveryResolve  string `xml:"DiscoveryResolve"`
						DiscoveryBye      string `xml:"DiscoveryBye"`
						RemoteDiscovery   string `xml:"RemoteDiscovery"`
						SystemBackup      string `xml:"SystemBackup"`
						SystemLogging     string `xml:"SystemLogging"`
						FirmwareUpgrade   string `xml:"FirmwareUpgrade"`
						SupportedVersions []struct {
							Text  string `xml:",chardata"`
							Major string `xml:"Major"`
							Minor string `xml:"Minor"`
						} `xml:"SupportedVersions"`
						Extension struct {
							Text                   string `xml:",chardata"`
							HttpFirmwareUpgrade    string `xml:"HttpFirmwareUpgrade"`
							HttpSystemBackup       string `xml:"HttpSystemBackup"`
							HttpSystemLogging      string `xml:"HttpSystemLogging"`
							HttpSupportInformation string `xml:"HttpSupportInformation"`
						} `xml:"Extension"`
					} `xml:"System"`
					IO struct {
						Text            string `xml:",chardata"`
						InputConnectors string `xml:"InputConnectors"`
						RelayOutputs    string `xml:"RelayOutputs"`
						Extension       struct {
							Text      string `xml:",chardata"`
							Auxiliary string `xml:"Auxiliary"`
							Extension string `xml:"Extension"`
						} `xml:"Extension"`
					} `xml:"IO"`
					Security struct {
						Text                 string `xml:",chardata"`
						TLS11                string `xml:"TLS1.1"`
						TLS12                string `xml:"TLS1.2"`
						OnboardKeyGeneration string `xml:"OnboardKeyGeneration"`
						AccessPolicyConfig   string `xml:"AccessPolicyConfig"`
						X509Token            string `xml:"X.509Token"`
						SAMLToken            string `xml:"SAMLToken"`
						KerberosToken        string `xml:"KerberosToken"`
						RELToken             string `xml:"RELToken"`
						Extension            struct {
							Text      string `xml:",chardata"`
							TLS10     string `xml:"TLS1.0"`
							Extension struct {
								Text               string `xml:",chardata"`
								Dot1X              string `xml:"Dot1X"`
								SupportedEAPMethod string `xml:"SupportedEAPMethod"`
								RemoteUserHandling string `xml:"RemoteUserHandling"`
							} `xml:"Extension"`
						} `xml:"Extension"`
					} `xml:"Security"`
				} `xml:"Device"`
				Events struct {
					Text                                          string `xml:",chardata"`
					XAddr                                         string `xml:"XAddr"`
					WSSubscriptionPolicySupport                   string `xml:"WSSubscriptionPolicySupport"`
					WSPullPointSupport                            string `xml:"WSPullPointSupport"`
					WSPausableSubscriptionManagerInterfaceSupport string `xml:"WSPausableSubscriptionManagerInterfaceSupport"`
				} `xml:"Events"`
				Imaging struct {
					Text  string `xml:",chardata"`
					XAddr string `xml:"XAddr"`
				} `xml:"Imaging"`
				Media struct {
					Text                  string `xml:",chardata"`
					XAddr                 string `xml:"XAddr"`
					StreamingCapabilities struct {
						Text         string `xml:",chardata"`
						RTPMulticast string `xml:"RTPMulticast"`
						RTPTCP       string `xml:"RTP_TCP"`
						RTPRTSPTCP   string `xml:"RTP_RTSP_TCP"`
					} `xml:"StreamingCapabilities"`
					Extension struct {
						Text                string `xml:",chardata"`
						ProfileCapabilities struct {
							Text                    string `xml:",chardata"`
							MaximumNumberOfProfiles string `xml:"MaximumNumberOfProfiles"`
						} `xml:"ProfileCapabilities"`
					} `xml:"Extension"`
				} `xml:"Media"`
				PTZ struct {
					Text  string `xml:",chardata"`
					XAddr string `xml:"XAddr"`
				} `xml:"PTZ"`
				Extension struct {
					Text     string `xml:",chardata"`
					DeviceIO struct {
						Text         string `xml:",chardata"`
						XAddr        string `xml:"XAddr"`
						VideoSources string `xml:"VideoSources"`
						VideoOutputs string `xml:"VideoOutputs"`
						AudioSources string `xml:"AudioSources"`
						AudioOutputs string `xml:"AudioOutputs"`
						RelayOutputs string `xml:"RelayOutputs"`
					} `xml:"DeviceIO"`
				} `xml:"Extension"`
			} `xml:"Capabilities"`
		} `xml:"GetCapabilitiesResponse"`
	} `xml:"Body"`
}
