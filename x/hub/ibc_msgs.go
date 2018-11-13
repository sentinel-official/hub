package hub

type IBCMsgCoinLocker struct {
	SrcChainId  string        `json:"src_chain_id"`
	DestChainId string        `json:"dest_chain_id"`
	Message     MsgCoinLocker `json:"message"`
}

type IBCMsgLockCoins struct {
	SrcChainId  string       `json:"src_chain_id"`
	DestChainId string       `json:"dest_chain_id"`
	Message     MsgLockCoins `json:"message"`
}

type IBCMsgReleaseCoins struct {
	SrcChainId  string          `json:"src_chain_id"`
	DestChainId string          `json:"dest_chain_id"`
	Message     MsgReleaseCoins `json:"message"`
}

type IBCMsgReleaseCoinsToMany struct {
	SrcChainId  string                `json:"src_chain_id"`
	DestChainId string                `json:"dest_chain_id"`
	Message     MsgReleaseCoinsToMany `json:"message"`
}
