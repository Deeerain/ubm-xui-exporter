package api

type ApiResponse[T any] struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     T      `json:"obj"`
}

type IPInfo struct {
	Ip string `json:"ip"`
}

type ClientIpInfo struct {
	Id  int      `json:"id"`
	Ips []IPInfo `json:"ips"`
}

type MemoryStatus struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type ServerStatus struct {
	Cpu   int          `json:"cpu"`
	Mem   MemoryStatus `json:"mem"`
	Swap  MemoryStatus `json:"swap"`
	Disk  MemoryStatus `json:"disk"`
	NetIO struct {
		Up   int `json:"up"`
		Down int `json:"down"`
	} `json:"netIO"`
	TcpCount int `json:"tcpCount"`
}
