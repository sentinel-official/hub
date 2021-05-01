package types

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type PaginatedIterator struct {
	index int
	items []sdk.Iterator
}

func NewPaginatedIterator(items ...sdk.Iterator) *PaginatedIterator {
	return &PaginatedIterator{
		items: items,
	}
}

func (p *PaginatedIterator) Skip(skip int64) {
	if skip <= 0 {
		return
	}

	for index, iter := range p.items {
		for ; skip > 0 && iter.Valid(); iter.Next() {
			skip = skip - 1
		}

		if skip == 0 {
			p.index = index
			return
		}
	}
}

func (p *PaginatedIterator) Limit(limit int64, iterFunc func(iter sdk.Iterator)) {
	if limit <= 0 {
		limit = int64(math.MaxInt64)
	}

	for index, iter := range p.items {
		if index < p.index {
			continue
		}

		for ; limit > 0 && iter.Valid(); iter.Next() {
			iterFunc(iter)
			limit = limit - 1
		}
	}
}

func (p *PaginatedIterator) Close() {
	for _, iter := range p.items {
		iter.Close()
	}
}
