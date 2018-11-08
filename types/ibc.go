package types

type IBCPacket struct {
	SrcChain  string      `json:"src_chain"`
	DestChain string      `json:"dest_chain"`
	Message   interface{} `json:"message"`
}
