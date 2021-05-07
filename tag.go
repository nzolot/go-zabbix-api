package zabbix

type Tag struct {
	//TagID   string `json:"hosttagid,omitempty"`
	//HostID  string `json:"hostid"`
	TagName string `json:"tag"`
	Value   string `json:"value"`
}

type Tags []Tag
