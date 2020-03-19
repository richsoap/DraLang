package util

import("testing")

func TestByteSliceEqual(T *testing.T) {
	a := []byte{'a', 'b', 'c'}
	b := []byte{'a', 'a', 'c'}
	if ByteSliceEqual(a, b) {
		T.Error("Different Equal Error")
	}
	if !ByteSliceEqual(a, a) {
		T.Error("Same Equal Error")
	}
}