package hasher

import (
	"crypto/md5"
	"fmt"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) ToMD5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
