package types

import (
	"fmt"
	"sort"
	"strconv"
)

type ID uint64

func NewIDFromUInt64(i uint64) ID { return ID(i) }

func NewIDFromString(s string) ID {
	i, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		panic(err)
	}

	return NewIDFromUInt64(i)
}

func (id ID) Uint64() uint64      { return uint64(id) }
func (id ID) String() string      { return fmt.Sprintf("%X", id.Uint64()) }
func (id ID) IsEqual(_id ID) bool { return id.Uint64() == _id.Uint64() }

type IDs []ID

func (ids IDs) Append(id ...ID) IDs { return append(ids, id...) }
func (ids IDs) Len() int            { return len(ids) }

func (ids IDs) Sort() IDs {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	return ids
}

func (ids IDs) Delete(index int) IDs {
	ids[index] = ids[ids.Len()-1]
	return ids[:ids.Len()-1]
}

func (ids IDs) Search(id ID) int {
	index := sort.Search(len(ids), func(i int) bool {
		return ids[i] >= id
	})

	if (index == ids.Len()) ||
		(index < ids.Len() && ids[index] != id) {
		return ids.Len()
	}

	return index
}
