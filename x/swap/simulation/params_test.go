package simulation

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParamChanges(t *testing.T) {
	s := rand.NewSource(1)
	r := rand.New(s)

	tests := []struct {
		composedKey string
		key         string
		simValue    string
		subspace    string
	}{
		{"swap/SwapDenom", "SwapDenom", "HeAerqyNEUzXPFGkqEGqiQWIXnkuHMYZLfGaEFPyynhwJyzAHyfjXUlrGhblTtxWduqtCDMLxiDHIMGFpXzp", "swap"},
		{"swap/SwapEnabled", "SwapEnabled", "true", "swap"},
		{"swap/ApproveBy", "ApproveBy", "sent1xeu4rw4zlakdguwys0c4lwgt4kehckpp0253jd", "swap"},
	}

	paramChanges := ParamChanges(r)

	for i, tt := range paramChanges {
		require.Equal(t, tests[i].composedKey, tt.ComposedKey())
		require.Equal(t, tests[i].key, tt.Key())
		require.Equal(t, tests[i].simValue, tt.SimValue()(r))
		require.Equal(t, tests[i].subspace, tt.Subspace())
	}
}
