package types

type IBCPacket struct {
	SrcChainId  string    `json:"src_chain_id"`
	DestChainId string    `json:"dest_chain_id"`
	Message     Interface `json:"message"`
}

type Interface interface{}