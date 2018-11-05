package jkdbs

type SCDevStatus struct {
	Mac        string
	IP         string
	DevType    string
	Version    string
	Time       int64
	RemoteIp   string
	ActionType string
	Online     bool
}

type SCDevStatusHistory struct {
	Index      int64
	Version    string
	Time       int64
	DevTime    int64
	RemoteIp   string
	DevType    string
	Mac        string
	IP         string
	ActionType string
	Duration   int
}
