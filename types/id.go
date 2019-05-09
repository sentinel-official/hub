package types

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

type ID string

func NewID(str string) ID {
	return ID(str)
}

func (id ID) Bytes() []byte  { return []byte(id) }
func (id ID) String() string { return string(id) }
func (id ID) Len() int       { return len(id) }

func (id ID) Hash() string {
	hash := tmhash.Sum(id.Bytes())
	return hex.EncodeToString(hash)
}

func (id ID) HashTruncated() string {
	hash := tmhash.SumTruncated(id.Bytes())
	return hex.EncodeToString(hash)
}

func (id ID) Valid() bool {
	splits := strings.Split(id.String(), "/")
	return len(splits) == 2
}

func IDFromOwnerAndCount(address csdkTypes.Address, count uint64) ID {
	id := fmt.Sprintf("%s/%d", address.String(), count)
	return NewID(id)
}

type IDs []ID

func NewIDs() IDs {
	return IDs{}
}

func (ids IDs) Append(id ...ID) IDs { return append(ids, id...) }
func (ids IDs) Len() int            { return len(ids) }
func (ids IDs) Less(i, j int) bool  { return ids[i].String() < ids[j].String() }
func (ids IDs) Swap(i, j int)       { ids[i], ids[j] = ids[j], ids[i] }

func (ids IDs) Sort() IDs {
	sort.Sort(ids)
	return ids
}

func (ids IDs) Search(id ID) int {
	index := sort.Search(len(ids), func(i int) bool {
		return ids[i].String() >= id.String()
	})

	if (index == ids.Len()) ||
		(index < ids.Len() && ids[index].String() != id.String()) {
		return ids.Len()
	}

	return index
}
