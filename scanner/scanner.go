package scanner

import (
	"bufio"
	"io"
	"os"
)

type Scanner struct {
	filePath string
	file     *os.File
	reader   *bufio.Reader
	lineNum  int
	eof      bool
}

func NewScanner(path string) *Scanner {
	return &Scanner{path, nil, nil, 0, true}
}

func (s *Scanner) Open() error {
	if f, err := os.Open(s.filePath); err != nil {
		return err
	} else {
		s.file = f
	}
	s.reader = bufio.NewReader(s.file)
	s.lineNum = 0
	s.eof = false
	return nil
}

// Make sure return return for user's calling is a complete line
func (s *Scanner) ReadLine() ([]byte, error) {
	if s.eof {
		return nil, nil
	}
	if val, isPrefix, err := s.reader.ReadLine(); err != nil {
		if err == io.EOF {
			return nil, nil
		} else {
			return nil, err
		}
	} else if isPrefix {
		tail, tailerr := s.ReadLine()
		if tailerr != nil {
			return nil, err
		} else if tail == nil {
			return val, nil
		}
		val = append(val, tail...)
		return val, nil
	} else {
		s.lineNum++
		return val, nil
	}
}

func (s *Scanner) Close() {
	if s.file != nil {
		s.file.Close()
	}
	s.file = nil
}

func (s *Scanner) LineNum() int {
	return s.lineNum
}
