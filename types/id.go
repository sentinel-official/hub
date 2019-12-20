package types

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	NodeIDPrefix         = "node"
	SessionIDPrefix      = "sess"
	SubscriptionIDPrefix = "subs"
	ResolverIDPrefix     = "reso"
)

type ID interface {
	String() string
	Uint64() uint64
	Bytes() []byte
	Prefix() string
	IsEqual(ID) bool
	MarshalJSON() ([]byte, error)
}

var (
	_ ID = NodeID{}
	_ ID = SessionID{}
	_ ID = SubscriptionID{}
	_ ID = ResolverID{}
)

type NodeID []byte

func NewNodeID(i uint64) NodeID {
	return types.Uint64ToBigEndian(i)
}

func NewNodeIDFromString(s string) (NodeID, error) {
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid node id length")
	}
	
	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}
	
	return NewNodeID(i), nil
}

func (id NodeID) String() string {
	return fmt.Sprintf("%s%x", NodeIDPrefix, id.Uint64())
}

func (id NodeID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id NodeID) Bytes() []byte {
	return id
}

func (id NodeID) Prefix() string {
	return NodeIDPrefix
}

func (id NodeID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id NodeID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *NodeID) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	
	_id, err := NewNodeIDFromString(s)
	if err != nil {
		return err
	}
	
	*id = _id
	
	return nil
}

type SessionID []byte

func NewSessionID(i uint64) SessionID {
	return types.Uint64ToBigEndian(i)
}

func NewSessionIDFromString(s string) (SessionID, error) {
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid session id length")
	}
	
	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		panic(err)
	}
	
	return NewSessionID(i), nil
}

func (id SessionID) String() string {
	return fmt.Sprintf("%s%x", SessionIDPrefix, id.Uint64())
}

func (id SessionID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id SessionID) Bytes() []byte {
	return id
}

func (id SessionID) Prefix() string {
	return SessionIDPrefix
}

func (id SessionID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id SessionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *SessionID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	
	_id, err := NewSessionIDFromString(s)
	if err != nil {
		return err
	}
	
	*id = _id
	
	return nil
}

type SubscriptionID []byte

func NewSubscriptionID(i uint64) SubscriptionID {
	return types.Uint64ToBigEndian(i)
}

func NewSubscriptionIDFromString(s string) (SubscriptionID, error) {
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid subscription id length")
	}
	
	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		return nil, err
	}
	
	return NewSubscriptionID(i), nil
}

func (id SubscriptionID) String() string {
	return fmt.Sprintf("%s%x", SubscriptionIDPrefix, id.Uint64())
}

func (id SubscriptionID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id SubscriptionID) Bytes() []byte {
	return id
}

func (id SubscriptionID) Prefix() string {
	return SubscriptionIDPrefix
}

func (id SubscriptionID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id SubscriptionID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *SubscriptionID) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	
	_id, err := NewSubscriptionIDFromString(s)
	if err != nil {
		return err
	}
	
	*id = _id
	
	return nil
}

type ResolverID []byte

func NewResolverID(i uint64) ResolverID {
	return types.Uint64ToBigEndian(i)
}

func NewResolverIDFromString(s string) (ResolverID, error) {
	if len(s) < 5 {
		return nil, fmt.Errorf("invalid resolver id length")
	}
	
	i, err := strconv.ParseUint(s[4:], 16, 64)
	if err != nil {
		panic(err)
	}
	
	return NewResolverID(i), nil
}

func (id ResolverID) String() string {
	return fmt.Sprintf("%s%x", ResolverIDPrefix, id.Uint64())
}

func (id ResolverID) Uint64() uint64 {
	return binary.BigEndian.Uint64(id)
}

func (id ResolverID) Bytes() []byte {
	return id
}

func (id ResolverID) Prefix() string {
	return ResolverIDPrefix
}

func (id ResolverID) IsEqual(_id ID) bool {
	return id.String() == _id.String()
}

func (id ResolverID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func (id *ResolverID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	
	_id, err := NewResolverIDFromString(s)
	if err != nil {
		return err
	}
	
	*id = _id
	
	return nil
}

var _ sort.Interface = (*IDs)(nil)

type IDs []ID

func (ids IDs) Append(id ID) IDs {
	return append(ids, id)
}

func (ids IDs) Len() int {
	return len(ids)
}

func (ids IDs) Less(x, y int) bool {
	i := strings.Compare(ids[x].Prefix(), ids[y].Prefix())
	if i < 0 {
		return true
	} else if i == 0 {
		return ids[x].Uint64() < ids[y].Uint64()
	}
	
	return false
}

func (ids IDs) Swap(x, y int) {
	ids[x], ids[y] = ids[y], ids[x]
}

func (ids IDs) Sort() IDs {
	sort.Slice(ids, ids.Less)
	return ids
}

func (ids IDs) Delete(x int) IDs {
	ids[x] = ids[ids.Len()-1]
	return ids[:ids.Len()-1]
}

func (ids IDs) Search(id ID) int {
	i := id.Uint64()
	index := sort.Search(len(ids), func(x int) bool {
		return ids[x].Prefix() > id.Prefix() || ids[x].Uint64() >= i
	})
	
	if (index == ids.Len()) ||
		(index < ids.Len() && ids[index].String() != id.String()) {
		return ids.Len()
	}
	
	return index
}
