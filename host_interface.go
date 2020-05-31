package zabbix

type (
	// InterfaceType different interface type
	InterfaceType int
)

const (
	// Differente type of zabbix interface
	// see "type" in https://www.zabbix.com/documentation/3.2/manual/api/reference/hostinterface/object

	// Agent type
	Agent InterfaceType = 1
	// SNMP type
	SNMP InterfaceType = 2
	// IPMI type
	IPMI InterfaceType = 3
	// JMX type
	JMX InterfaceType = 4
)

// HostInterface represents zabbix host interface type
// https://www.zabbix.com/documentation/3.2/manual/api/reference/hostinterface/object
type HostInterface struct {
	DNS   string        `json:"dns"`
	IP    string        `json:"ip"`
	Main  int           `json:"main,string"`
	Port  string        `json:"port"`
	Type  InterfaceType `json:"type,string"`
	UseIP int           `json:"useip,string"`
	InterfaceId int     `json:"interfaceid,string"`
}

// HostInterfaces is an array of HostInterface
type HostInterfaces []HostInterface
