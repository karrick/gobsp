package guanoloco

import (
	"bufio"
	"io"
)

type MessageType VWI

type ErrUnknownMessageType VWI

func (e ErrUnknownMessageType) Error() string {
	return "unknown message type: " + VWI(e).String()
}

type MessageHandler func(io.Reader) error

func NewScanner(ior io.Reader, handlers map[uint32]MessageHandler) *Scanner {
	s := &Scanner{
		br:       bufio.NewReader(ior),
		handlers: handlers,
	}
	return s
}

type Scanner struct {
	br       *bufio.Reader
	err      error
	mt, ms   VWI
	handlers map[uint32]MessageHandler
}

func (s *Scanner) Scan() bool {
	if s.err = s.mt.UnmarshalBinaryFrom(s.br); s.err != nil {
		if s.err == io.EOF {
			s.err = nil
		}
		return false
	}
	if s.err = s.ms.UnmarshalBinaryFrom(s.br); s.err != nil {
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
	handler, ok := s.handlers[uint32(s.mt)]
	if !ok {
		s.err = ErrUnknownMessageType(s.mt)
		return s.err
	}
	return handler(io.LimitReader(s.br, int64(s.ms)))
}
