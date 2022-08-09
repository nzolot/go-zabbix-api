package zabbix

// ProxyInterface represents zabbix proxy interface type
// https://www.zabbix.com/documentation/6.0/en/manual/api/reference/proxy/object
type ProxyInterface struct {
	DNS   string `json:"dns,omitempty"`
	IP    string `json:"ip,omitempty"`
	Port  string `json:"port,omitempty"`
	UseIP int    `json:"useip,string,omitempty"`
}

// ProxyInterfaces is an array of ProxyInterface
type ProxyInterfaces []ProxyInterface
