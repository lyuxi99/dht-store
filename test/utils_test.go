package test

import (
	"DHT/internal/utils"
	"bytes"
	"crypto/sha1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBytesPower2(t *testing.T) {
	a := []byte{100, 5, 255, 255, 200, 50}
	assert.Equal(t, 0, bytes.Compare([]byte{100, 5, 255, 255, 200, 51}, utils.AddBytesPower2(a, 0)))
	assert.Equal(t, 0, bytes.Compare([]byte{100, 5, 255, 255, 200, 66}, utils.AddBytesPower2(a, 4)))
	assert.Equal(t, 0, bytes.Compare([]byte{100, 6, 0, 0, 8, 50}, utils.AddBytesPower2(a, 14)))
}
func TestIsInRange(t *testing.T) {
	assert.Equal(t, true, utils.IsInRange([]byte{10}, []byte{5}, []byte{20}))
	assert.Equal(t, true, utils.IsInRange([]byte{20}, []byte{5}, []byte{20}))
	assert.Equal(t, false, utils.IsInRange([]byte{5}, []byte{5}, []byte{20}))
	assert.Equal(t, false, utils.IsInRange([]byte{1}, []byte{5}, []byte{20}))
	assert.Equal(t, false, utils.IsInRange([]byte{25}, []byte{5}, []byte{20}))

	assert.Equal(t, true, utils.IsInRange([]byte{1}, []byte{200}, []byte{10}))
	assert.Equal(t, true, utils.IsInRange([]byte{10}, []byte{200}, []byte{10}))
	assert.Equal(t, false, utils.IsInRange([]byte{15}, []byte{200}, []byte{10}))
	assert.Equal(t, false, utils.IsInRange([]byte{200}, []byte{200}, []byte{10}))
	assert.Equal(t, true, utils.IsInRange([]byte{210}, []byte{200}, []byte{10}))

	assert.Equal(t, true, utils.IsInRange([]byte{100}, []byte{200}, []byte{200}))
	assert.Equal(t, true, utils.IsInRange([]byte{200}, []byte{200}, []byte{200}))
	assert.Equal(t, true, utils.IsInRange([]byte{210}, []byte{200}, []byte{200}))
	assert.Equal(t, true, utils.IsInRange([]byte{210}, []byte{200}, []byte{200}))
	assert.Equal(t, true, utils.IsInRange([]byte{210}, []byte{200}, []byte{200}))
}
func TestIsInRangeExclude(t *testing.T) {
	assert.Equal(t, true, utils.IsInRangeExclude([]byte{10}, []byte{5}, []byte{20}))
	assert.Equal(t, false, utils.IsInRangeExclude([]byte{10}, []byte{200}, []byte{10}))
}

func TestSHA1(t *testing.T) {
	data := []byte("abc")
	sum := sha1.Sum(data)
	assert.Equal(t, sum[:], utils.SHA1(data))
}
