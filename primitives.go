package gobsp

import (
	"io"
	"strconv"
	"unsafe"
)

type Binary interface {
	MarshalBinaryTo(io.Writer) error
	UnmarshalBinaryFrom(io.Reader) error
}

type Int8 int8

func (v Int8) MarshalBinaryTo(iow io.Writer) error {
	// Use ByteWriter optimization, if available
	if bw, ok := iow.(io.ByteWriter); ok {
		return bw.WriteByte(byte(v))
	}
	// Otherwise, just use tiny slice
	_, err := iow.Write([]byte{byte(v)})
	return err
}

func (v *Int8) UnmarshalBinaryFrom(ior io.Reader) error {
	// Use ByteReader optimization, if available
	if br, ok := ior.(io.ByteReader); ok {
		b, err := br.ReadByte()
		if err == nil {
			*v = Int8(b)
		}
		return err
	}
	// Otherwise, just use tiny slice
	buf := make([]byte, 1)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Int8(buf[0])
	}
	return err
}

func (v Int8) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Uint8 uint8

func (v Uint8) MarshalBinaryTo(iow io.Writer) error {
	// Use ByteWriter optimization, if available
	if bw, ok := iow.(io.ByteWriter); ok {
		return bw.WriteByte(byte(v))
	}
	// Otherwise, just use tiny slice
	_, err := iow.Write([]byte{byte(v)})
	return err
}

func (v *Uint8) UnmarshalBinaryFrom(ior io.Reader) error {
	// Use ByteReader optimization, if available
	if br, ok := ior.(io.ByteReader); ok {
		b, err := br.ReadByte()
		if err == nil {
			*v = Uint8(b)
		}
		return err
	}
	// Otherwise, just use tiny slice
	buf := make([]byte, 1)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Uint8(buf[0])
	}
	return err
}

func (v Uint8) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Int16 int16

func (v Int16) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Int16) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 2)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Int16(int16(buf[0])<<8 | int16(buf[1]))
	}
	return err
}

func (v Int16) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Uint16 uint16

func (v Uint16) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Uint16) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 2)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Uint16(uint16(buf[0])<<8 | uint16(buf[1]))
	}
	return err
}

func (v Uint16) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Int32 int32

func (v Int32) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Int32) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 4)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Int32(int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3]))
	}
	return err
}

func (v Int32) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Uint32 uint32

func (v Uint32) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Uint32) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 4)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Uint32(uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3]))
	}
	return err
}

func (v Uint32) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Int64 int64

func (v Int64) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Int64) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 8)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Int64(int64(buf[0])<<56 | int64(buf[1])<<48 | int64(buf[2])<<40 | int64(buf[3])<<32 |
			int64(buf[4])<<24 | int64(buf[5])<<16 | int64(buf[6])<<8 | int64(buf[7]))
	}
	return err
}

func (v Int64) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Uint64 uint64

func (v Uint64) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Uint64) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 8)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = Uint64(uint64(buf[0])<<56 | uint64(buf[1])<<48 | uint64(buf[2])<<40 | uint64(buf[3])<<32 |
			uint64(buf[4])<<24 | uint64(buf[5])<<16 | uint64(buf[6])<<8 | uint64(buf[7]))
	}
	return err
}

func (v Uint64) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

func encodeVWI(iow io.Writer, value uint64) error {
	// Use ByteWriter optimization if available (~3x optimization)
	if iobw, ok := iow.(io.ByteWriter); ok {
		if value == 0 {
			return iobw.WriteByte(byte(0))
		}
		for value > 0 {
			b := byte(value & 127)
			value >>= 7
			if value != 0 {
				b |= 128
			}
			if err := iobw.WriteByte(byte(b)); err != nil {
				return err
			}
		}
		return nil
	}

	// Otherwise, just use tiny buffer
	if value == 0 {
		_, err := iow.Write([]byte{0})
		return err
	}
	for value > 0 {
		b := byte(value & 127)
		value >>= 7
		if value != 0 {
			b |= 128
		}
		if _, err := iow.Write([]byte{b}); err != nil {
			return err
		}
	}
	return nil
}

func decodeVWI(ior io.Reader) (uint64, error) {
	const mask = byte(127)
	const flag = byte(128)
	var value uint64

	// use ByteReader optimization if available (~3x optimization)
	if iobr, ok := ior.(io.ByteReader); ok {
		for shift := uint(0); ; shift += 7 {
			b, err := iobr.ReadByte()
			if err != nil {
				return 0, err
			}
			value |= uint64(b&mask) << shift
			if b&flag == 0 {
				break
			}
		}
		return value, nil
	}

	// Otherwise, just use tiny buffer
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(ior, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		value |= uint64(b&mask) << shift
		if b&flag == 0 {
			break
		}
	}
	return value, nil
}

type VWI int64

func (v VWI) MarshalBinaryTo(iow io.Writer) error {
	// move sign bit from most to least significant bit
	value := uint64((v << 1) ^ (v >> 63))
	return encodeVWI(iow, value)
}

func (v *VWI) UnmarshalBinaryFrom(ior io.Reader) error {
	value, err := decodeVWI(ior)
	if err == nil {
		// move the sign bit from least to most significant bit
		*v = VWI((int64(value>>1) ^ -int64(value&1)))
	}
	return err
}

func (v VWI) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type UVWI uint64

func (v UVWI) MarshalBinaryTo(iow io.Writer) error {
	return encodeVWI(iow, uint64(v))
}

func (v *UVWI) UnmarshalBinaryFrom(ior io.Reader) error {
	value, err := decodeVWI(ior)
	if err == nil {
		*v = UVWI(value)
	}
	return err
}

func (v UVWI) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

type Float32 float32

func (v Float32) MarshalBinaryTo(iow io.Writer) error {
	vv := *(*uint32)(unsafe.Pointer(&v))
	buf := []byte{
		byte(vv >> 24),
		byte(vv >> 16),
		byte(vv >> 8),
		byte(vv),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Float32) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(ior, buf); err != nil {
		return err
	}
	j := uint32(buf[0])<<24 |
		uint32(buf[1])<<16 |
		uint32(buf[2])<<8 |
		uint32(buf[3])
	*v = Float32(*(*float32)(unsafe.Pointer(&j)))
	return nil
}

func (v Float32) String() string {
	return strconv.FormatFloat(float64(v), 'g', -1, 32)
}

type Float64 float64

func (v Float64) MarshalBinaryTo(iow io.Writer) error {
	vv := *(*uint64)(unsafe.Pointer(&v))
	buf := []byte{
		byte(vv >> 56),
		byte(vv >> 48),
		byte(vv >> 40),
		byte(vv >> 32),
		byte(vv >> 24),
		byte(vv >> 16),
		byte(vv >> 8),
		byte(vv),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Float64) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(ior, buf); err != nil {
		return err
	}
	j := uint64(buf[0])<<56 |
		uint64(buf[1])<<48 |
		uint64(buf[2])<<40 |
		uint64(buf[3])<<32 |
		uint64(buf[4])<<24 |
		uint64(buf[5])<<16 |
		uint64(buf[6])<<8 |
		uint64(buf[7])
	*v = Float64(*(*float64)(unsafe.Pointer(&j)))
	return nil
}

func (v Float64) String() string {
	return strconv.FormatFloat(float64(v), 'g', -1, 64)
}

type String string

func (v String) MarshalBinaryTo(iow io.Writer) error {
	if err := UVWI(len(string(v))).MarshalBinaryTo(iow); err != nil {
		return err
	}
	_, err := io.WriteString(iow, string(v))
	return err
}

func (v *String) UnmarshalBinaryFrom(ior io.Reader) error {
	var size UVWI
	if err := size.UnmarshalBinaryFrom(ior); err != nil {
		return err
	}
	buf := make([]byte, size)
	_, err := io.ReadFull(ior, buf)
	if err == nil {
		*v = String(buf)
	}
	return err
}

func (v String) String() string {
	return string(v)
}

type StringSlice []String

func (v StringSlice) MarshalBinaryTo(iow io.Writer) error {
	size := UVWI(len(v))
	if err := size.MarshalBinaryTo(iow); err != nil {
		return err
	}
	for _, s := range v {
		if err := s.MarshalBinaryTo(iow); err != nil {
			return err
		}
	}
	return nil
}

func (v *StringSlice) UnmarshalBinaryFrom(ior io.Reader) error {
	var size UVWI
	if err := size.UnmarshalBinaryFrom(ior); err != nil {
		return err
	}
	ss := make([]String, size)
	for i := uint64(0); i < uint64(size); i++ {
		var s String
		if err := s.UnmarshalBinaryFrom(ior); err != nil {
			return err
		}
		ss[i] = s
	}
	*v = ss
	return nil
}
