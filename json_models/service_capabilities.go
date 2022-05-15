package json_models

type ServiceCapabilities struct {
	Network struct {
		IPFilter            bool `json:"IPFilter,string"`
		ZeroConfiguration   bool `json:"ZeroConfiguration,string"`
		IPVersion6          bool `json:"IPVersion6,string"`
		DynDNS              bool `json:"DynDNS,string"`
		Dot11Configuration  bool `json:"Dot11Configuration,string"`
		Dot1XConfigurations int  `json:"Dot1XConfigurations,string"`
		HostnameFromDHCP    bool `json:"HostnameFromDHCP,string"`
		Ntp                 int  `json:"NTP,string"`
		DHCPv6              bool `json:"DHCPv6,string"`
	} `json:"Network"`
	Security struct {
		TLS10                bool `json:"TLS10,string"`
		TLS11                bool `json:"TLS11,string"`
		TLS12                bool `json:"TLS12,string"`
		OnboardKeyGeneration bool `json:"OnboardKeyGeneration,string"`
		AccessPolicyConfig   bool `json:"AccessPolicyConfig,string"`
		DefaultAccessPolicy  bool `json:"DefaultAccessPolicy,string"`
		Dot1X                bool `json:"Dot1X,string"`
		RemoteUserHandling   bool `json:"RemoteUserHandling,string"`
		X509Token            bool `json:"X509Token,string"`
		SAMLToken            bool `json:"SAMLToken,string"`
		KerberosToken        bool `json:"KerberosToken,string"`
		UsernameToken        bool `json:"UsernameToken,string"`
		HTTPDigest           bool `json:"HttpDigest,string"`
		RELToken             bool `json:"RELToken,string"`
		SupportedEAPMethods  int  `json:"SupportedEAPMethods,string"`
		MaxUsers             int  `json:"MaxUsers,string"`
		MaxUserNameLength    int  `json:"MaxUserNameLength,string"`
		MaxPasswordLength    int  `json:"MaxPasswordLength,string"`
	} `json:"Security"`
	System struct {
		DiscoveryResolve       bool `json:"DiscoveryResolve,string"`
		DiscoveryBye           bool `json:"DiscoveryBye,string"`
		RemoteDiscovery        bool `json:"RemoteDiscovery,string"`
		SystemBackup           bool `json:"SystemBackup,string"`
		SystemLogging          bool `json:"SystemLogging,string"`
		FirmwareUpgrade        bool `json:"FirmwareUpgrade,string"`
		HTTPFirmwareUpgrade    bool `json:"HttpFirmwareUpgrade,string"`
		HTTPSystemBackup       bool `json:"HttpSystemBackup,string"`
		HTTPSystemLogging      bool `json:"HttpSystemLogging,string"`
		HTTPSupportInformation bool `json:"HttpSupportInformation,string"`
		StorageConfiguration   bool `json:"StorageConfiguration,string"`
	} `json:"System"`
	Misc struct {
		AuxiliaryCommands string `json:"AuxiliaryCommands"`
	} `json:"Misc"`
}
