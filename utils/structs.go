package utils

type ClientHelloMsg struct {
	Random []byte
}

type ServerHelloMsg struct {
	Random      []byte
	Certificate []byte
}

type Block struct {
	Id   int
	Data []byte
	Hash string
}
