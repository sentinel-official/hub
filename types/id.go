package types

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ProviderIDPrefix     = "prov"
	NodeIDPrefix         = "node"
	SessionIDPrefix      = "sess"
	SubscriptionIDPrefix = "subs"
)

type ID interface {
	String() string
	Uint64() uint64
	Bytes() []byte
	Prefix() string
	IsEqual(v ID) bool
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(bytes []byte) error
}

var (
	_ ID = &ProviderID{}
	_ ID = &NodeID{}
	_ ID = &SessionID{}
	_ ID = &SubscriptionID{}
)

type ProviderID []byte

func NewProviderID(i uint64) ProviderID {
	return sdk.Uint64ToBigEndian(i)
}

func (p ProviderID) String() string {
	return fmt.Sprintf("%s%x", ProviderID{}, p.Uint64())
}

func (p ProviderID) Uint64() uint64 {
	return binary.BigEndian.Uint64(p)
}

func (p ProviderID) Bytes() []byte {
	return p
}

func (p ProviderID) Prefix() string {
	return ProviderIDPrefix
}

func (p ProviderID) IsEqual(v ID) bool {
	return p.String() == v.String()
}

func (p ProviderID) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.String())
}

func (p *ProviderID) UnmarshalJSON(bytes []byte) error {
	var v string
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}

	id, err := NewProviderIDFromString(v)
	if err != nil {
		return err
	}

	*p = id
	return nil
}

func NewProviderIDFromString(s string) (ProviderID, error) {
	if len(s) < 5 || s[:4] != ProviderIDPrefix {
		return nil, fmt.Errorf("invalid provider id")
	}

	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}

	return NewProviderID(i), nil
}

type NodeID []byte

func NewNodeID(i uint64) NodeID {
	return sdk.Uint64ToBigEndian(i)
}

func (n NodeID) String() string {
	return fmt.Sprintf("%s%x", NodeIDPrefix, n.Uint64())
}

func (n NodeID) Uint64() uint64 {
	return binary.BigEndian.Uint64(n)
}

func (n NodeID) Bytes() []byte {
	return n
}

func (n NodeID) Prefix() string {
	return NodeIDPrefix
}

func (n NodeID) IsEqual(v ID) bool {
	return n.String() == v.String()
}

func (n NodeID) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *NodeID) UnmarshalJSON(bytes []byte) error {
	var x string
	if err := json.Unmarshal(bytes, &x); err != nil {
		return err
	}

	id, err := NewNodeIDFromString(x)
	if err != nil {
		return err
	}

	*n = id
	return nil
}

func NewNodeIDFromString(s string) (NodeID, error) {
	if len(s) < 5 || s[:4] != NodeIDPrefix {
		return nil, fmt.Errorf("invalid node id")
	}

	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}

	return NewNodeID(i), nil
}

type SubscriptionID []byte

func NewSubscriptionID(i uint64) SubscriptionID {
	return sdk.Uint64ToBigEndian(i)
}

func (s SubscriptionID) String() string {
	return fmt.Sprintf("%s%x", SubscriptionIDPrefix, s.Uint64())
}

func (s SubscriptionID) Uint64() uint64 {
	return binary.BigEndian.Uint64(s)
}

func (s SubscriptionID) Bytes() []byte {
	return s
}

func (s SubscriptionID) Prefix() string {
	return SubscriptionIDPrefix
}

func (s SubscriptionID) IsEqual(v ID) bool {
	return s.String() == v.String()
}

func (s SubscriptionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SubscriptionID) UnmarshalJSON(bytes []byte) error {
	var x string
	if err := json.Unmarshal(bytes, &x); err != nil {
		return err
	}

	id, err := NewSubscriptionIDFromString(x)
	if err != nil {
		return err
	}

	*s = id
	return nil
}

func NewSubscriptionIDFromString(s string) (SubscriptionID, error) {
	if len(s) < 5 || s[:4] != SubscriptionIDPrefix {
		return nil, fmt.Errorf("invalid subscription id")
	}

	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}

	return NewSubscriptionID(i), nil
}

type SessionID []byte

func NewSessionID(i uint64) SessionID {
	return sdk.Uint64ToBigEndian(i)
}

func (s SessionID) String() string {
	return fmt.Sprintf("%s%x", SessionIDPrefix, s.Uint64())
}

func (s SessionID) Uint64() uint64 {
	return binary.BigEndian.Uint64(s)
}

func (s SessionID) Bytes() []byte {
	return s
}

func (s SessionID) Prefix() string {
	return SessionIDPrefix
}

func (s SessionID) IsEqual(v ID) bool {
	return s.String() == v.String()
}

func (s SessionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *SessionID) UnmarshalJSON(bytes []byte) error {
	var x string
	if err := json.Unmarshal(bytes, &x); err != nil {
		return err
	}

	id, err := NewSessionIDFromString(x)
	if err != nil {
		return err
	}

	*s = id
	return nil
}

func NewSessionIDFromString(s string) (SessionID, error) {
	if len(s) < 5 || s[:4] != SessionIDPrefix {
		return nil, fmt.Errorf("invalid session id")
	}

	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}

	return NewSessionID(i), nil
}

var _ sort.Interface = (*IDs)(nil)

type IDs []ID

func (i IDs) Append(id ID) IDs {
	return append(i, id)
}

func (i IDs) Len() int {
	return len(i)
}

func (i IDs) Less(x, y int) bool {
	v := strings.Compare(i[x].Prefix(), i[y].Prefix())
	if v < 0 {
		return true
	} else if v == 0 {
		return i[x].Uint64() < i[y].Uint64()
	}

	return false
}

func (i IDs) Swap(x, y int) {
	i[x], i[y] = i[y], i[x]
}

func (i IDs) Sort() IDs {
	sort.Slice(i, i.Less)
	return i
}

func (i IDs) Delete(x int) IDs {
	i[x] = i[i.Len()-1]
	return i[:i.Len()-1]
}

func (i IDs) Search(id ID) int {
	v := id.Uint64()
	index := sort.Search(len(i), func(x int) bool {
		return i[x].Prefix() > id.Prefix() || i[x].Uint64() >= v
	})

	if index == i.Len() || !i[index].IsEqual(id) {
		return i.Len()
	}

	return index
}
