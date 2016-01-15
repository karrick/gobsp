package guanoloco

import (
	"io"
	"strconv"
	"unsafe"
)

////////////////////////////////////////

type Binary interface {
	MarshalBinaryTo(io.Writer) error
	UnmarshalBinaryFrom(io.Reader) error
}

////////////////////////////////////////

type Uint32 uint32

func (v *Uint32) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(*v >> 24),
		byte(*v >> 16),
		byte(*v >> 8),
		byte(*v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Uint32) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(ior, buf); err != nil {
		return err
	}
	*v = Uint32(uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3]))
	return nil
}

func (v Uint32) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

////////////////////////////////////////

type Uint64 uint64

func (v *Uint64) MarshalBinaryTo(iow io.Writer) error {
	buf := []byte{
		byte(*v >> 56),
		byte(*v >> 48),
		byte(*v >> 40),
		byte(*v >> 32),
		byte(*v >> 24),
		byte(*v >> 16),
		byte(*v >> 8),
		byte(*v),
	}
	_, err := iow.Write(buf)
	return err
}

func (v *Uint64) UnmarshalBinaryFrom(ior io.Reader) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(ior, buf); err != nil {
		return err
	}
	*v = Uint64(uint64(buf[0])<<56 | uint64(buf[1])<<48 | uint64(buf[2])<<40 | uint64(buf[3])<<32 | uint64(buf[4])<<24 | uint64(buf[5])<<16 | uint64(buf[6])<<8 | uint64(buf[7]))
	return nil
}

func (v Uint64) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

////////////////////////////////////////

type VWI uint64

func (v *VWI) MarshalBinaryTo(iow io.Writer) error {
	var err error
	value := uint(*v)
	if value == 0 {
		_, err = iow.Write([]byte{0})
	}
	for value > 0 {
		b := byte(value & 127)
		value >>= 7
		if value != 0 {
			b |= 128
		}
		if _, err = iow.Write([]byte{b}); err != nil {
			break
		}
	}
	return err
}

func (v *VWI) UnmarshalBinaryFrom(ior io.Reader) error {
	const mask = byte(127)
	const flag = byte(128)
	var value uint64

	buf := make([]byte, 1)

	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(ior, buf); err != nil {
			return err
		}
		b := buf[0]
		value |= uint64(b&mask) << shift
		if b&flag == 0 {
			break
		}
	}

	*v = VWI(value)
	return nil
}

func (v VWI) String() string {
	return strconv.FormatUint(uint64(v), 10)
}

////////////////////////////////////////

type Float64 float64

func (v *Float64) MarshalBinaryTo(iow io.Writer) error {
	vv := *(*uint64)(unsafe.Pointer(v))
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

////////////////////////////////////////

type String string

func (v *String) MarshalBinaryTo(iow io.Writer) error {
	size := VWI(len(string(*v)))
	if err := size.MarshalBinaryTo(iow); err != nil {
		return err
	}
	_, err := io.WriteString(iow, string(*v))
	return err
}

func (v *String) UnmarshalBinaryFrom(ior io.Reader) error {
	var size VWI
	if err := size.UnmarshalBinaryFrom(ior); err != nil {
		return err
	}
	buf := make([]byte, size)
	if _, err := io.ReadFull(ior, buf); err != nil {
		return err
	}
	*v = String(buf)
	return nil
}

func (v String) String() string {
	return string(v)
}

////////////////////////////////////////

type StringSlice []String

func (v *StringSlice) MarshalBinaryTo(iow io.Writer) error {
	var err error
	size := VWI(len(*v))
	if err = size.MarshalBinaryTo(iow); err != nil {
		return err
	}
	for _, s := range *v {
		if err := s.MarshalBinaryTo(iow); err != nil {
			return err
		}
	}
	return nil
}

func (v *StringSlice) UnmarshalBinaryFrom(ior io.Reader) error {
	var size VWI
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