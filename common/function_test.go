package common

import (
	// "bytes"
	"testing"
)

func TestInt64Tobytes(t *testing.T) {
	num := uint64(-123)
	data := Uint64Tobytes(num)
	new_num := BytesToUint64(data)
	if new_num != num {
		t.Errorf("new_num:%d neq num:%d", new_num, num)
	}
}
