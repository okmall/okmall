package model

// Result indicate response result
type Result struct {
	ErrorNo int         `json:"errno"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}
