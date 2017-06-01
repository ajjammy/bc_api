package Resp

type ResponseStatus int

const (
	SUCCESS ResponseStatus = iota
	FAIL
	ERROR
)

type Response struct {
	Status  string `json:"status"`
	Message string         `json:"message,omitempty"`
	//Link    `json:"links"`
	Data    interface{} `json:"data,omitempty"`
}
