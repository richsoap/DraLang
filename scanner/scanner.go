package scanner

import(
	"github.com/richsoap/soaplang/errors"
	"os"
	"bufio"
)

type Scanner struct {
	filePath string
	file *os.File
	reader *bufio.Reader
}

func NewScanner(path string) *Scanner {
	return &Scanner{path, nil, nil}
}

func (s *Scanner) Open() error {
	if f, err := os.Open(s.filePath); err != nil {
		return err
	} else {
		s.file = f
	}
	s.reader = bufio.NewReader(s.file)
	return nil
}

func (s *Scanner) Read() (byte, error) {
	if val, err := s.reader.ReadByte(); err != nil {
		return 0, errors.ERR_EOF
	} else {
		return val, nil
	}
}

func (s *Scanner) Close() {
	if s.file != nil {
		s.file.Close()
	}
	s.file = nil
}