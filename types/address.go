package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"gopkg.in/yaml.v2"
)

const (
	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
	Bech32MainPrefix = "sent"

	// PrefixValidator is the prefix for validator keys.
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys.
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys.
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys.
	PrefixOperator = "oper"

	// PrefixProvider is the prefix for provider keys.
	PrefixProvider = "prov"
	// PrefixNode is the prefix for node keys.
	PrefixNode = "node"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address.
	Bech32PrefixAccAddr = Bech32MainPrefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key.
	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address.
	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key.
	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a validator's consensus address.
	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a validator's consensus public key.
	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic

	// Bech32PrefixProvAddr defines the Bech32 prefix of a providers's address.
	Bech32PrefixProvAddr = Bech32MainPrefix + PrefixProvider
	// Bech32PrefixProvPub defines the Bech32 prefix of a providers's public key.
	Bech32PrefixProvPub = Bech32MainPrefix + PrefixProvider + PrefixPublic
	// Bech32PrefixNodeAddr defines the Bech32 prefix of a node's address.
	Bech32PrefixNodeAddr = Bech32MainPrefix + PrefixNode
	// Bech32PrefixNodePub defines the Bech32 prefix of a node's public key.
	Bech32PrefixNodePub = Bech32MainPrefix + PrefixNode + PrefixPublic
)

var (
	_ sdk.Address    = ProvAddress{}
	_ yaml.Marshaler = ProvAddress{}

	_ sdk.Address    = NodeAddress{}
	_ yaml.Marshaler = NodeAddress{}
)

type ProvAddress []byte

func (p ProvAddress) Equals(address sdk.Address) bool {
	if p.Empty() && address == nil {
		return true
	}

	return bytes.Equal(p.Bytes(), address.Bytes())
}

func (p ProvAddress) Empty() bool {
	return bytes.Equal(p.Bytes(), ProvAddress{}.Bytes())
}

func (p ProvAddress) Bytes() []byte {
	return p
}

func (p ProvAddress) String() string {
	if p.Empty() {
		return ""
	}

	s, err := bech32.ConvertAndEncode(GetConfig().GetBech32ProviderAddrPrefix(), p.Bytes())
	if err != nil {
		panic(err)
	}

	return s
}

func (p ProvAddress) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		_, _ = f.Write([]byte(p.String()))
	case 'p':
		_, _ = f.Write([]byte(fmt.Sprintf("%p", p)))
	default:
		_, _ = f.Write([]byte(fmt.Sprintf("%X", p.Bytes())))
	}
}

func (p ProvAddress) Marshal() ([]byte, error) {
	return p, nil
}

func (p ProvAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p ProvAddress) MarshalYAML() (interface{}, error) {
	return p.String(), nil
}

func (p *ProvAddress) Unmarshal(data []byte) error {
	*p = data
	return nil
}

func (p *ProvAddress) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	address, err := ProvAddressFromBech32(s)
	if err != nil {
		return err
	}

	*p = address
	return nil
}

func (p *ProvAddress) UnmarshalYAML(data []byte) error {
	var s string
	if err := yaml.Unmarshal(data, &s); err != nil {
		return err
	}

	address, err := ProvAddressFromBech32(s)
	if err != nil {
		return err
	}

	*p = address
	return nil
}

func ProvAddressFromBech32(s string) (ProvAddress, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ProvAddress{}, nil
	}

	bz, err := sdk.GetFromBech32(s, GetConfig().GetBech32ProviderAddrPrefix())
	if err != nil {
		return nil, err
	}

	if err = sdk.VerifyAddressFormat(bz); err != nil {
		return nil, err
	}

	return bz, nil
}

type NodeAddress []byte

func (n NodeAddress) Equals(address sdk.Address) bool {
	if n.Empty() && address == nil {
		return true
	}

	return bytes.Equal(n.Bytes(), address.Bytes())
}

func (n NodeAddress) Empty() bool {
	return bytes.Equal(n.Bytes(), NodeAddress{}.Bytes())
}

func (n NodeAddress) Bytes() []byte {
	return n
}

func (n NodeAddress) String() string {
	if n.Empty() {
		return ""
	}

	s, err := bech32.ConvertAndEncode(GetConfig().GetBech32NodeAddrPrefix(), n.Bytes())
	if err != nil {
		panic(err)
	}

	return s
}

func (n NodeAddress) Format(f fmt.State, c rune) {
	switch c {
	case 's':
		_, _ = f.Write([]byte(n.String()))
	case 'p':
		_, _ = f.Write([]byte(fmt.Sprintf("%p", n)))
	default:
		_, _ = f.Write([]byte(fmt.Sprintf("%X", n.Bytes())))
	}
}

func (n NodeAddress) Marshal() ([]byte, error) {
	return n, nil
}

func (n NodeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n NodeAddress) MarshalYAML() (interface{}, error) {
	return n.String(), nil
}

func (n *NodeAddress) Unmarshal(data []byte) error {
	*n = data
	return nil
}

func (n *NodeAddress) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	address, err := NodeAddressFromBech32(s)
	if err != nil {
		return err
	}

	*n = address
	return nil
}

func (n *NodeAddress) UnmarshalYAML(data []byte) error {
	var s string
	if err := yaml.Unmarshal(data, &s); err != nil {
		return err
	}

	address, err := NodeAddressFromBech32(s)
	if err != nil {
		return err
	}

	*n = address
	return nil
}

func NodeAddressFromBech32(s string) (NodeAddress, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return NodeAddress{}, nil
	}

	bz, err := sdk.GetFromBech32(s, GetConfig().GetBech32NodeAddrPrefix())
	if err != nil {
		return nil, err
	}

	if err = sdk.VerifyAddressFormat(bz); err != nil {
		return nil, err
	}

	return bz, nil
}
