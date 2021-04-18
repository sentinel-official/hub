package v0_5

import (
	"encoding/hex"
	"encoding/json"
)

const (
	EthereumHashLength = 32
)

type (
	EthereumHash [EthereumHashLength]byte
)

func BytesToHash(b []byte) EthereumHash {
	var a EthereumHash
	a.SetBytes(b)
	return a
}

func (e *EthereumHash) SetBytes(b []byte) {
	if len(b) > len(e) {
		b = b[len(b)-EthereumHashLength:]
	}

	copy(e[EthereumHashLength-len(b):], b)
}

func (e EthereumHash) Bytes() []byte { return e[:] }

func (e EthereumHash) Hex() string { return hex.EncodeToString(e[:]) }

func (e EthereumHash) String() string { return e.Hex() }

func (e EthereumHash) MarshalJSON() ([]byte, error) { return json.Marshal(e.String()) }

func (e *EthereumHash) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	hash, err := hex.DecodeString(s)
	if err != nil {
		return err
	}

	*e = BytesToHash(hash)
	return nil
}
