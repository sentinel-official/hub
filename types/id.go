package types

import (
	"encoding/hex"
	"sort"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

type ID []byte

func NewID(b string) ID {
	return ID(b)
}

func (id ID) Bytes() []byte  { return id }
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

func (ids IDs) Delete(index int) IDs {
	ids[index] = ids[ids.Len()-1]
	return ids[:ids.Len()-1]
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
