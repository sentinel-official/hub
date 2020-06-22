package types

import (
	"fmt"
)

type Bandwidth struct {
	Upload   uint64 `json:"upload"`
	Download uint64 `json:"download"`
}

func NewBandwidth(upload, download uint64) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}

func (n Bandwidth) IsAnyZero() bool {
	return n.Upload == 0 || n.Download == 0
}

func (n Bandwidth) IsAllZero() bool {
	return n.Upload == 0 && n.Download == 0
}

func (n Bandwidth) String() string {
	return fmt.Sprintf("%d↑, %d↓", n.Upload, n.Download)
}
