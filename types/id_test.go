package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewID(t *testing.T) {
	id1 := NewID("")
	require.Equal(t, ID(""), id1)
	require.Equal(t, []byte{}, id1.Bytes())
	require.Equal(t, "", id1.String())
	require.Equal(t, 0, id1.Len())
	require.Equal(t, false, id1.Valid())
}

func TestIDFromOwnerAndCount(t *testing.T) {
	id1 := IDFromOwnerAndCount(TestAddress1, 0)
	require.Equal(t, ID(TestAddress1.String()+"/0"), id1)
	require.Equal(t, []byte(TestAddress1.String()+"/0"), id1.Bytes())
	require.Equal(t, TestAddress1.String()+"/0", id1.String())
	require.Equal(t, 47, id1.Len())
	require.Equal(t, true, id1.Valid())
}

func TestNewIDs(t *testing.T) {
	ids := NewIDs()
	require.Equal(t, 0, ids.Len())
	require.Equal(t, IDs{}, ids.Sort())

	ids = ids.Append(ID("address/0"))
	require.Equal(t, 1, ids.Len())
	require.Equal(t, IDs{ID("address/0")}, ids.Sort())
	ids = ids.Append(ID("address/10"))
	require.Equal(t, 2, ids.Len())
	require.Equal(t, IDs{ID("address/0"), ID("address/10")}, ids.Sort())
	ids = ids.Append(ID("address/5"))
	require.Equal(t, 3, ids.Len())
	require.Equal(t, IDs{ID("address/0"), ID("address/10"), ID("address/5")}, ids.Sort())

	require.Equal(t, 0, ids.Search(ID("address/0")))
	require.Equal(t, 1, ids.Search(ID("address/10")))
	require.Equal(t, 2, ids.Search(ID("address/5")))
	require.Equal(t, 3, ids.Search(ID("address/-1")))
	require.Equal(t, 3, ids.Search(ID("address/1")))
	require.Equal(t, 3, ids.Search(ID("address/6")))
	require.Equal(t, 3, ids.Search(ID("address/11")))

	ids.Swap(1, 2)
	require.Equal(t, IDs{ID("address/0"), ID("address/5"), ID("address/10")}, ids)
}
