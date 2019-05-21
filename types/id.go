package types

import (
	"encoding/hex"
	"encoding/json"
	"sort"

	"github.com/tendermint/tendermint/crypto/tmhash"
)

type ID []byte

func IDFromString(str string) ID {
	bz, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}

	return ID(bz)
}

func (id ID) String() string               { return hex.EncodeToString(id) }
func (id ID) Bytes() []byte                { return id }
func (id ID) Len() int                     { return len(id) }
func (id ID) MarshalJSON() ([]byte, error) { return json.Marshal(id.String()) }

func (id *ID) UnmarshalJSON(data []byte) error {
	var str string

	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	*id = IDFromString(str)
	return nil
}

func (id ID) Hash() []byte          { return tmhash.Sum(id.Bytes()) }
func (id ID) HashTruncated() []byte { return tmhash.SumTruncated(id.Bytes()) }

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

// nolint: interfacer
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
