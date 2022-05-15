package json_models

type Capabilities struct {
	Analytics struct {
		XAddr                  string `json:"XAddr"`
		RuleSupport            string `json:"RuleSupport"`
		AnalyticsModuleSupport string `json:"AnalyticsModuleSupport"`
	} `json:"Analytics"`
	Device struct {
		XAddr   string `json:"XAddr"`
		Network struct {
			IPFilter          string `json:"IPFilter"`
			ZeroConfiguration string `json:"ZeroConfiguration"`
			IPVersion6        string `json:"IPVersion6"`
			DynDNS            string `json:"DynDNS"`
			Extension         struct {
				Text               string `json:"Text"`
				Dot11Configuration string `json:"Dot11Configuration"`
			} `json:"Extension"`
		} `json:"Network"`
		System struct {
			DiscoveryResolve  string `json:"DiscoveryResolve"`
			DiscoveryBye      string `json:"DiscoveryBye"`
			RemoteDiscovery   string `json:"RemoteDiscovery"`
			SystemBackup      string `json:"SystemBackup"`
			SystemLogging     string `json:"SystemLogging"`
			FirmwareUpgrade   string `json:"FirmwareUpgrade"`
			SupportedVersions []struct {
				Major string `json:"Major"`
				Minor string `json:"Minor"`
			} `json:"SupportedVersions"`
			Extension struct {
				HTTPFirmwareUpgrade    string `json:"HttpFirmwareUpgrade"`
				HTTPSystemBackup       string `json:"HttpSystemBackup"`
				HTTPSystemLogging      string `json:"HttpSystemLogging"`
				HTTPSupportInformation string `json:"HttpSupportInformation"`
			} `json:"Extension"`
		} `json:"System"`
		Io struct {
			InputConnectors string `json:"InputConnectors"`
			RelayOutputs    string `json:"RelayOutputs"`
			Extension       struct {
				Auxiliary string `json:"Auxiliary"`
				Extension string `json:"Extension"`
			} `json:"Extension"`
		} `json:"IO"`
		Security struct {
			TLS11                string `json:"TLS11"`
			TLS12                string `json:"TLS12"`
			OnboardKeyGeneration string `json:"OnboardKeyGeneration"`
			AccessPolicyConfig   string `json:"AccessPolicyConfig"`
			X509Token            string `json:"X509Token"`
			SAMLToken            string `json:"SAMLToken"`
			KerberosToken        string `json:"KerberosToken"`
			RELToken             string `json:"RELToken"`
			Extension            struct {
				TLS10     string `json:"TLS10"`
				Extension struct {
					Dot1X              string `json:"Dot1X"`
					SupportedEAPMethod string `json:"SupportedEAPMethod"`
					RemoteUserHandling string `json:"RemoteUserHandling"`
				} `json:"Extension"`
			} `json:"Extension"`
		} `json:"Security"`
	} `json:"Device"`
	Events struct {
		XAddr                                         string `json:"XAddr"`
		WSSubscriptionPolicySupport                   string `json:"WSSubscriptionPolicySupport"`
		WSPullPointSupport                            string `json:"WSPullPointSupport"`
		WSPausableSubscriptionManagerInterfaceSupport string `json:"WSPausableSubscriptionManagerInterfaceSupport"`
	} `json:"Events"`
	Imaging struct {
		XAddr string `json:"XAddr"`
	} `json:"Imaging"`
	Media struct {
		XAddr                 string `json:"XAddr"`
		StreamingCapabilities struct {
			RTPMulticast string `json:"RTPMulticast"`
			Rtptcp       string `json:"RTPTCP"`
			Rtprtsptcp   string `json:"RTPRTSPTCP"`
		} `json:"StreamingCapabilities"`
		Extension struct {
			ProfileCapabilities struct {
				MaximumNumberOfProfiles string `json:"MaximumNumberOfProfiles"`
			} `json:"ProfileCapabilities"`
		} `json:"Extension"`
	} `json:"Media"`
	Ptz struct {
		XAddr string `json:"XAddr"`
	} `json:"PTZ"`
	Extension struct {
		DeviceIO struct {
			XAddr        string `json:"XAddr"`
			VideoSources string `json:"VideoSources"`
			VideoOutputs string `json:"VideoOutputs"`
			AudioSources string `json:"AudioSources"`
			AudioOutputs string `json:"AudioOutputs"`
			RelayOutputs string `json:"RelayOutputs"`
		} `json:"DeviceIO"`
	} `json:"Extension"`
}
