package response

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		body := Body{
			Code: 400,
			Data: nil,
			Msg:  "失败",
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := Body{
		Code: 200,
		Data: resp,
		Msg:  "成功",
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
