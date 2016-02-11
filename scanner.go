package guanoloco

import (
	"bufio"
	"io"
	"io/ioutil"
)

type MessageType VWI

type ErrUnknownMessageType VWI

func (e ErrUnknownMessageType) Error() string {
	return "unknown message type: " + VWI(e).String()
}

type MessageHandler func(io.Reader) error

func NewScanner(ior io.Reader, handlers map[uint32]MessageHandler) *Scanner {
	s := &Scanner{
		bufferedReader: bufio.NewReader(ior),
		handlers:       handlers,
	}
	return s
}

type Scanner struct {
	bufferedReader           *bufio.Reader
	err                      error
	messageType, messageSize VWI
	handlers                 map[uint32]MessageHandler
}

func (s *Scanner) Scan() bool {
	if s.err = s.messageType.UnmarshalBinaryFrom(s.bufferedReader); s.err != nil {
		if s.err == io.EOF {
			s.err = nil
		}
		return false
	}
	if s.err = s.messageSize.UnmarshalBinaryFrom(s.bufferedReader); s.err != nil {
		if s.err == io.EOF {
			s.err = nil
		}
		return false
	}
	return true
}

// Err returns the error object associated with this scanner, or nil
// if no errors have occurred.
func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) Handle() error {
	if s.err != nil {
		return s.err
	}

	limitReader := io.LimitReader(s.bufferedReader, int64(s.messageSize))
	handler, ok := s.handlers[uint32(s.messageType)]
	if !ok {
		io.Copy(ioutil.Discard, limitReader) // consume unknown message
		s.err = ErrUnknownMessageType(s.messageType)
		return s.err
	}
	return handler(limitReader)
}
