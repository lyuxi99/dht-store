package utils

import (
	"bytes"
	"crypto/sha1"
	"log"
	"os"
)

// Exists checks whether the given path exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CheckAndMakeDir creates directory if the path does not exist.
func CheckAndMakeDir(path string) {
	if !Exists(path) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			log.Fatalf("create folder \"%v\" error", path)
		}
	}
}

// SHA1 computes the sha1 sum of the data.
func SHA1(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}

// CheckIdentity checks whether the identity is valid
func CheckIdentity(id []byte, address string) bool {
	return bytes.Equal(id, SHA1([]byte(address)))
}

// IsInRange returns whether c in the range (l,r].
func IsInRange(c, l, r []byte) bool {
	if bytes.Compare(l, r) < 0 {
		return bytes.Compare(l, c) < 0 && bytes.Compare(c, r) <= 0
	} else {
		return bytes.Compare(l, c) < 0 || bytes.Compare(c, r) <= 0
	}
}

// IsInRangeExclude returns whether c in the range (l,r).
func IsInRangeExclude(c, l, r []byte) bool {
	return !bytes.Equal(c, r) && IsInRange(c, l, r)
}

// AddBytesPower2 adds 2**`exp` to `a`.
func AddBytesPower2(a []byte, exp int) []byte {
	c := make([]byte, len(a))
	copy(c, a)
	M := len(a)

	pos := M - 1 - exp/8
	exp = exp % 8
	carry := 1 << exp
	for i := pos; i >= 0; i-- {
		cur := int(c[i]) + carry
		carry = cur / 256
		c[i] = uint8(cur % 256)
	}
	return c
}
