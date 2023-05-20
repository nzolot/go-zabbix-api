package zabbix

type (
	AvailableType int
	StatusType    int
	TimestampType uint64
)

type PreProc struct {
	Type               string `json:"type"`
	Params             string `json:"params,omitempty"`
	ErrorHandler       string `json:"error_handler"`
	ErrorHandlerParams string `json:"error_handler_params"`
}

type PreProcs []PreProc

const (
	Available   AvailableType = 1
	Unavailable AvailableType = 2

	Monitored   StatusType = 0
	Unmonitored StatusType = 1

	ActiveProxy  StatusType = 5
	PassiveProxy StatusType = 6
)

const (
	ZbxApiErrorParameters int = -32602
	ZbxApiErrorInternal   int = -32500
)

type (
	// SeverityType of a trigger
	// Zabbix severity see : https://www.zabbix.com/documentation/3.2/manual/api/reference/trigger/object
	SeverityType int
)

const (
	// Different severity see : https://www.zabbix.com/documentation/3.2/manual/config/triggers/severity

	// Ok value is used for Services
	Ok SeverityType = -1
	// NotClassified is Not classified severity
	NotClassified SeverityType = 0
	// Information is Information severity
	Information SeverityType = 1
	// Warning is Warning severity
	Warning SeverityType = 2
	// Average is Average severity
	Average SeverityType = 3
	// High is high severity
	High SeverityType = 4
	// Critical is critical severity
	Critical SeverityType = 5
)
