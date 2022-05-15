package xml_models

type ServiceCapabilities struct {
	Body struct {
		Text                           string `xml:",chardata"`
		GetServiceCapabilitiesResponse struct {
			Text         string `xml:",chardata"`
			Capabilities struct {
				Text    string `xml:",chardata"`
				Network struct {
					Text                string `xml:",chardata"`
					IPFilter            string `xml:"IPFilter,attr"`
					ZeroConfiguration   string `xml:"ZeroConfiguration,attr"`
					IPVersion6          string `xml:"IPVersion6,attr"`
					DynDNS              string `xml:"DynDNS,attr"`
					Dot11Configuration  string `xml:"Dot11Configuration,attr"`
					Dot1XConfigurations string `xml:"Dot1XConfigurations,attr"`
					HostnameFromDHCP    string `xml:"HostnameFromDHCP,attr"`
					NTP                 string `xml:"NTP,attr"`
					DHCPv6              string `xml:"DHCPv6,attr"`
				} `xml:"Network"`
				Security struct {
					Text                 string `xml:",chardata"`
					TLS10                string `xml:"TLS1.0,attr"`
					TLS11                string `xml:"TLS1.1,attr"`
					TLS12                string `xml:"TLS1.2,attr"`
					OnboardKeyGeneration string `xml:"OnboardKeyGeneration,attr"`
					AccessPolicyConfig   string `xml:"AccessPolicyConfig,attr"`
					DefaultAccessPolicy  string `xml:"DefaultAccessPolicy,attr"`
					Dot1X                string `xml:"Dot1X,attr"`
					RemoteUserHandling   string `xml:"RemoteUserHandling,attr"`
					X509Token            string `xml:"X.509Token,attr"`
					SAMLToken            string `xml:"SAMLToken,attr"`
					KerberosToken        string `xml:"KerberosToken,attr"`
					UsernameToken        string `xml:"UsernameToken,attr"`
					HttpDigest           string `xml:"HttpDigest,attr"`
					RELToken             string `xml:"RELToken,attr"`
					SupportedEAPMethods  string `xml:"SupportedEAPMethods,attr"`
					MaxUsers             string `xml:"MaxUsers,attr"`
					MaxUserNameLength    string `xml:"MaxUserNameLength,attr"`
					MaxPasswordLength    string `xml:"MaxPasswordLength,attr"`
				} `xml:"Security"`
				System struct {
					Text                   string `xml:",chardata"`
					DiscoveryResolve       string `xml:"DiscoveryResolve,attr"`
					DiscoveryBye           string `xml:"DiscoveryBye,attr"`
					RemoteDiscovery        string `xml:"RemoteDiscovery,attr"`
					SystemBackup           string `xml:"SystemBackup,attr"`
					SystemLogging          string `xml:"SystemLogging,attr"`
					FirmwareUpgrade        string `xml:"FirmwareUpgrade,attr"`
					HttpFirmwareUpgrade    string `xml:"HttpFirmwareUpgrade,attr"`
					HttpSystemBackup       string `xml:"HttpSystemBackup,attr"`
					HttpSystemLogging      string `xml:"HttpSystemLogging,attr"`
					HttpSupportInformation string `xml:"HttpSupportInformation,attr"`
					StorageConfiguration   string `xml:"StorageConfiguration,attr"`
				} `xml:"System"`
				Misc struct {
					Text              string `xml:",chardata"`
					AuxiliaryCommands string `xml:"AuxiliaryCommands,attr"`
				} `xml:"Misc"`
			} `xml:"Capabilities"`
		} `xml:"GetServiceCapabilitiesResponse"`
	} `xml:"Body"`
}
