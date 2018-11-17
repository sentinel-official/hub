package types

type IBCPacket struct {
	SrcChainID  string    `json:"src_chain_id"`
	DestChainID string    `json:"dest_chain_id"`
	Message     Interface `json:"message"`
}

type Interface interface{}
