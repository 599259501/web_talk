package entity

import (
	"net/http"
	"websocket/config"
	"websocket/protocols"
)

type Response struct {
	w *http.ResponseWriter
}

func MakeResponse(w *http.ResponseWriter) *Response{
	return &Response{
		w: w,
	}
}

func (r *Response) Success(data map[string]interface{}){
	r.Response(config.SUCCESS_CODE, "ok", data)
}

func (r *Response) Fail(code int, msg string, data interface{}){
	r.Response(code, msg, data)
}


func (r *Response)Response(code int, msg string, data interface{}){
	jsonProtocol := protocols.NewJsonProtocol()

	(*r.w).Write(jsonProtocol.Pack(struct {
		code int
		msg string
		data interface{}
	}{
		code,msg, data,
	}))
}
