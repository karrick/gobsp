package guanoloco

import (
	"bytes"
	"io"
	"runtime"
	"strings"
	"testing"
)

func ensure(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		_, file, line, ok := runtime.Caller(1)
		if !ok {
			t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
		} else {
			if index := strings.LastIndex(file, "/"); index != -1 {
				file = file[index+1:]
			}
			t.Errorf("Actual: %#v; Expected: %#v; %s:%d", actual, expected, file, line)
		}
	}
}

func TestBinaryScannerEOF(t *testing.T) {
	bb := new(bytes.Buffer)
	handlers := make(map[uint32]MessageHandler)
	scanner := NewScanner(bb, handlers)

	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Err(), nil)
}

func TestBinaryScannerHandleReturnsHandlerError(t *testing.T) {
	bb := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x01, 0x00,
	})

	handlers := map[uint32]MessageHandler{
		0x00: func(io.Reader) error {
			return io.EOF
		},
		0x01: func(io.Reader) error {
			return nil
		},
	}

	scanner := NewScanner(bb, handlers)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), io.EOF)
	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)
	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Err(), nil)
}

func TestBinaryScanner(t *testing.T) {
	t.Skip()

	bb := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x01, 0x00,
	})

	handlers := map[uint32]MessageHandler{
		0x00: func(ior io.Reader) error {
			return io.EOF
		},
		0x01: func(ior io.Reader) error {
			return nil
		},
	}

	scanner := NewScanner(bb, handlers)

	for scanner.Scan() {
		if err := scanner.Handle(); err != nil {
			t.Error(err)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
	}
}
