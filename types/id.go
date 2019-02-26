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

func (n ID) Bytes() []byte  { return []byte(n) }
func (n ID) String() string { return string(n) }
func (n ID) Len() int       { return len(n) }

func (n ID) Hash() string {
	hash := tmhash.Sum(n.Bytes())
	return hex.EncodeToString(hash)
}

func (n ID) HashTruncated() string {
	hash := tmhash.SumTruncated(n.Bytes())
	return hex.EncodeToString(hash)
}

func (n ID) Valid() bool {
	splits := strings.Split(n.String(), "/")
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

func (n IDs) Append(id ...ID) IDs { return append(n, id...) }
func (n IDs) Len() int            { return len(n) }
func (n IDs) Less(i, j int) bool  { return n[i].String() < n[j].String() }
func (n IDs) Swap(i, j int)       { n[i], n[j] = n[j], n[i] }

func (n IDs) Sort() IDs {
	sort.Sort(n)
	return n
}

func (n IDs) Search(id ID) int {
	index := sort.Search(len(n), func(i int) bool {
		return n[i].String() >= id.String()
	})

	if (index == n.Len()) ||
		(index < n.Len() && n[index].String() != id.String()) {
		return n.Len()
	}

	return index
}
