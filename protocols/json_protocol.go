package protocols

import "encoding/json"

type JsonProtocol struct {
}

func NewJsonProtocol() *JsonProtocol{
	return &JsonProtocol{};
}

func (pProtocol *JsonProtocol)Pack(body interface{}) []byte{
	jsonBody,_ := json.Marshal(body)
	return jsonBody
}

func (pProtocol *JsonProtocol)UnPack(body []byte) []byte{
	return []byte{}
}
