package Reader

import (
	"io"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)

	n, err := reader.Read(p)

	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

func OpenFile(name string) {
	file, err := os.Open(name)
	
	if err != nil {
		return err
	}

	defer file.Close()
}
