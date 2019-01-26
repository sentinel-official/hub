package types

type Bandwidth struct {
	Upload   uint64 `json:"upload"`
	Download uint64 `json:"download"`
}
