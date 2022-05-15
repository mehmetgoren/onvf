package json_models

type DNSFromDHCP struct {
	Type        string
	IPv4Address string
}

type DNS struct {
	FromDHCP    string
	DNSFromDHCP []DNSFromDHCP
}
