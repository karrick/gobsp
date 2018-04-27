package gobsp

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

func TestBinaryScannerNoHandlers(t *testing.T) {
	bb := new(bytes.Buffer)
	_, err := NewScanner(bb)
	if _, ok := err.(ErrScannerHasNoHandlers); err == nil || !ok {
		t.Errorf("Actual: %#v; Expected: %#v", err, ErrScannerHasNoHandlers{})
	}
}

func TestBinaryScannerEOF(t *testing.T) {
	bb := new(bytes.Buffer)
	handlers := make(map[uint32]MessageHandler)

	scanner, err := NewScanner(bb, Handlers(handlers))
	if err != nil {
		t.Fatal(err)
	}

	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Err(), nil)
}

func TestBinaryScannerHandleReturnsHandlerError(t *testing.T) {
	bb := bytes.NewReader([]byte{
		0x00, 0x02 /* payload: */, 0xDE, 0xAD,
		0x01, 0x02 /* payload: */, 0xBE, 0xEF,
	})

	handlers := map[uint32]MessageHandler{
		0: func(io.Reader) error {
			return io.ErrNoProgress // some dummy token error
		},
		1: func(io.Reader) error {
			return nil
		},
	}

	scanner, err := NewScanner(bb, Handlers(handlers))
	if err != nil {
		t.Fatal(err)
	}

	// First scan
	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), io.ErrNoProgress)
	ensure(t, scanner.Err(), error(nil))

	// Second scan
	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), error(nil))

	// Third scan
	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Handle(), error(nil))
}

func TestBinaryScannerHandleUnknownMessageTypeWithoutDefaultHandler(t *testing.T) {
	bb := bytes.NewBuffer([]byte{
		0x00, 0x00, // normal message
		0x01, 0x04, 0xDE, 0xAD, 0xBE, 0xEF, // unknown message and payload
		0x00, 0x00, // normal message to ensure it keeps going
	})

	handlers := map[uint32]MessageHandler{
		0: func(ior io.Reader) error {
			return nil
		},
	}

	scanner, err := NewScanner(bb, Handlers(handlers))
	if err != nil {
		t.Fatal(err)
	}

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), ErrUnknownMessageType(1))

	ensure(t, scanner.Scan(), false) // ??? b/c it remembers error ???
	ensure(t, scanner.Err(), ErrUnknownMessageType(0x01))
}

func TestBinaryScannerHandleUnknownMessageTypeWithDefaultHandler(t *testing.T) {
	bb := bytes.NewBuffer([]byte{
		0x00, 0x00, // normal message
		0x01, 0x04, 0xDE, 0xAD, 0xBE, 0xEF, // unknown message and payload
		0x00, 0x00, // normal message to ensure it keeps going
	})

	handlers := map[uint32]MessageHandler{
		0: func(ior io.Reader) error {
			return nil
		},
	}

	scanner, err := NewScanner(bb,
		DefaultHandler(DiscardAll),
		Handlers(handlers))
	if err != nil {
		t.Fatal(err)
	}

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Err(), nil)
}

func TestBinaryScannerHandleForgetsToConsumeAll(t *testing.T) {
	bb := bytes.NewBuffer([]byte{
		0x00, 0x00, // normal message to ensure it keeps going
		0x01, 0x04, 0xDE, 0xAD, 0xBE, 0xEF, // message and payload
		0x00, 0x00, // normal message to ensure it keeps going
	})

	handlers := map[uint32]MessageHandler{
		0: func(ior io.Reader) error {
			return nil
		},
		1: func(ior io.Reader) error {
			// read a few of the bytes, but not all of them
			buf := make([]byte, 2)
			io.ReadFull(ior, buf)
			return nil
		},
	}

	scanner, err := NewScanner(bb,
		Handlers(handlers))
	if err != nil {
		t.Fatal(err)
	}

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), true)
	ensure(t, scanner.Handle(), nil)

	ensure(t, scanner.Scan(), false)
	ensure(t, scanner.Err(), nil)
}
