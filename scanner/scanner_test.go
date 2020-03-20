package scanner

import (
	"testing"

	"github.com/richsoap/soaplang/util"
)

func checkReadLine(t *testing.T, s *Scanner, tar []byte) {
	if res, err := s.ReadLine(); err != nil {
		t.Error("Read Error")
	} else if !util.ByteSliceEqual(res, tar) {
		t.Error("Val Error: tar: ", tar, " get: ", res)
	}
}

func checkLineNum(t *testing.T, s *Scanner, tar int) {
	if s.LineNum() != tar {
		t.Error("LineNum Error: tar ", tar, " get: ", s.LineNum())
	}
}

func TestScanner(t *testing.T) {
	s := NewScanner("../material/scanner_test.txt")
	if err := s.Open(); err != nil {
		t.Error("Init Error")
	}
	defer s.Close()
	checkReadLine(t, s, []byte{'A', 'B', 'C'})
	checkReadLine(t, s, []byte{'D', 'E', 'F'})
	checkReadLine(t, s, nil)
	checkLineNum(t, s, 2)
}
