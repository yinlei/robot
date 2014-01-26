package common

//	模块之间消息

type Message struct {
	id    int
	src   int
	des   int
	data  []byte
	extra interface{}
}
