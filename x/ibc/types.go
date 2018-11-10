package ibc

import "github.com/ironman0x7b2/sentinel-sdk/x/hub"

type IBCMsgCoinLocker struct {
	SrcChainId  string            `json:"src_chain_id"`
	DestChainId string            `json:"dest_chain_id"`
	Message     hub.MsgCoinLocker `json:"message"`
}

type IBCMsgLockCoins struct {
	SrcChainId  string           `json:"src_chain_id"`
	DestChainId string           `json:"dest_chain_id"`
	Message     hub.MsgLockCoins `json:"message"`
}

type IBCMsgReleaseCoins struct {
	SrcChainId  string              `json:"src_chain_id"`
	DestChainId string              `json:"dest_chain_id"`
	Message     hub.MsgReleaseCoins `json:"message"`
}

type IBCMsgReleaseCoinsToMany struct {
	SrcChainId  string                    `json:"src_chain_id"`
	DestChainId string                    `json:"dest_chain_id"`
	Message     hub.MsgReleaseCoinsToMany `json:"message"`
}
