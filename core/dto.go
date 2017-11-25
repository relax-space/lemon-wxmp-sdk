package core

type RespErrorDto struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
