package gobsp

import (
	"bufio"
	"io"
	"io/ioutil"
)

// MessageType is a variable width integer that specifies which user-defined type a particular
// message payload should be interpreted as.
type MessageType VWI

// ErrScannerHasNoHandlers is an error that is returned when NewScanner is called without any
// message handlers.
type ErrScannerHasNoHandlers struct{}

func (e ErrScannerHasNoHandlers) Error() string {
	return "scanner has no message handlers"
}

// ErrUnknownMessageType is an error that is returned by Scanner.Handle() when the required
// user-defined message type has not been defined and there is no default message handler.
type ErrUnknownMessageType VWI

func (e ErrUnknownMessageType) Error() string {
	return "unknown message type: " + VWI(e).String()
}

// MessageHandler is any function that consumes the entirety of the specified io.Reader stream and
// returns any error that occurred while processing that message.
type MessageHandler func(io.Reader) error

// ScannerConfigurator is a function that modifies a newly created Scanner instance.
type ScannerConfigurator func(*Scanner) error

// DefaultHandler specifies a handler to invoke when the required message type does not have a
// defined handler.
func DefaultHandler(handler MessageHandler) ScannerConfigurator {
	return func(s *Scanner) error {
		s.defaultHandler = handler
		return nil
	}
}

// Handlers is used to specify the user-defined message types for a Scanner instance.
func Handlers(handlers map[uint32]MessageHandler) ScannerConfigurator {
	return func(s *Scanner) error {
		s.handlers = handlers
		return nil
	}
}

// NewScanner returns a new Scanner instance to process messages from the specified io.Reader
// stream, using the message handlers specified by the DefaultHandler and Handlers functions.
func NewScanner(ior io.Reader, configurators ...ScannerConfigurator) (*Scanner, error) {
	s := &Scanner{
		bufferedReader: bufio.NewReader(ior),
	}
	for _, c := range configurators {
		if err := c(s); err != nil {
			return nil, err
		}
	}
	if s.defaultHandler == nil && s.handlers == nil {
		return nil, ErrScannerHasNoHandlers{}
	}
	return s, nil
}

// Scanner defines an object used to scan binary messages from a stream of bytes from a particular
// io.Reader.
type Scanner struct {
	bufferedReader           *bufio.Reader
	err                      error
	messageType, messageSize VWI
	handlers                 map[uint32]MessageHandler
	defaultHandler           MessageHandler
}

// Err returns the error object associated with this scanner, or nil
// if no errors have occurred.
func (s *Scanner) Err() error {
	return s.err
}

// Reset sets the scannner's error state to nil.
func (s *Scanner) Reset() {
	s.err = nil
}

// Scan reads enough bytes from the stream to determine the message type and size. Normally returns
// true, but returns false on error or EOF.
func (s *Scanner) Scan() bool {
	if s.err != nil {
		return false
	}
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

// Handle invokes the message handler for the most recently received message type. If the required
// message handler is not defined, it invokes the default handler. If there is no default handler,
// it returns an error.
func (s *Scanner) Handle() error {
	if s.err != nil {
		return s.err
	}
	limitReader := io.LimitReader(s.bufferedReader, int64(s.messageSize))
	defer DiscardAll(limitReader)
	handler, ok := s.handlers[uint32(s.messageType)]
	if !ok {
		if s.defaultHandler == nil {
			s.err = ErrUnknownMessageType(s.messageType)
			return s.err
		}
		return s.defaultHandler(limitReader)
	}
	return handler(limitReader)
}

// DiscardAll discards the remaining bytes to be read from the specified io.Reader, returning any
// errors received while reading.
func DiscardAll(ior io.Reader) error {
	_, err := io.Copy(ioutil.Discard, ior)
	return err
}
