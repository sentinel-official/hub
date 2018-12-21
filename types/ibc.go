package types

import (
	"fmt"
)

func EgressKey(destChainID string, length uint64) string {
	return fmt.Sprintf("egress/%s/%d", destChainID, length)
}

func EgressLengthKey(destChainID string) string {
	return fmt.Sprintf("egress/%s", destChainID)
}

func IngressKey(srcChainID string, length uint64) string {
	return fmt.Sprintf("ingress/%s/%d", srcChainID, length)
}

func IngressLengthKey(srcChainID string) string {
	return fmt.Sprintf("ingress/%s", srcChainID)
}

type IBCPacket struct {
	SrcChainID  string    `json:"src_chain_id"`
	DestChainID string    `json:"dest_chain_id"`
	Message     Interface `json:"message"`
}

type Interface interface{}
