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
