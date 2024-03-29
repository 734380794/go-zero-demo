// Code generated by goctl. DO NOT EDIT.
package types

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

type UserInfo struct {
	UserName string `json:"username"`
	Addr     string `json:"addr"`
	Id       uint   `json:"id"`
}

type UserInfoResponse struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
}
