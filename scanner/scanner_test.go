package scanner

import (
	"testing"

	"github.com/richsoap/soaplang/errors"
	"github.com/richsoap/soaplang/util"
)

func TestScanner(t *testing.T) {
	s := NewScanner("../material/scanner_test.txt")
	defer s.Close()
	if err := s.Open(); err != nil {
		t.Error("Init Error")
	}
	res := make([]byte, 0)
	tar := []byte{'A', 'B', 'C'}
	for {
		val, err := s.Read()
		if err == errors.ERR_EOF {
			break
		} else if err != nil {
			t.Error("Read Error")
		} else {
			res = append(res, val)
		}
	}
	if !util.ByteSliceEqual(res, tar) {
		t.Error("Val Error")
	}
}
