package protocols

type MessageProtocol interface {
	Pack(body interface{}) []byte
	UnPack(body []byte) []byte
}
