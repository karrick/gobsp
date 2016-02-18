package gobsp

import (
	"bytes"
	"math"
	"testing"

	"github.com/karrick/buffer"
)

////////////////////////////////////////
// a few test interfaces
////////////////////////////////////////

type testBuffer interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type testBytes interface {
	Bytes() []byte
}

////////////////////////////////////////
// Int8
////////////////////////////////////////

func testBinaryInt8(t *testing.T, value int, buf []byte) {
	// ensure works for both io.Reader & io.Writer, and io.ByteReader & io.ByteWriter
	test := func(t *testing.T, value int, buf []byte, scratch testBuffer) {
		vin := Int8(value)
		var vout Int8

		if err := vin.MarshalBinaryTo(scratch); err != nil {
			t.Error(err)
		}

		if sb, ok := scratch.(testBytes); ok {
			if actual, expected := sb.Bytes(), buf; !bytes.Equal(actual, expected) {
				t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
			}
		}

		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			t.Error(err)
		}

		if actual, expected := vout, vin; actual != expected {
			t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}

	test(t, value, buf, new(buffer.Buffer))
	test(t, value, buf, new(bytes.Buffer))
}

func TestBinaryInt8(t *testing.T) {
	testBinaryInt8(t, 0, []byte{0x00})
	testBinaryInt8(t, 1, []byte{0x01})
	testBinaryInt8(t, -1, []byte{0xFF})
	testBinaryInt8(t, 2, []byte{0x02})
	testBinaryInt8(t, -2, []byte{0xFE})
	testBinaryInt8(t, 127, []byte{0x7f})
	testBinaryInt8(t, -127, []byte{0x81})
}

////////////////////////////////////////
// Uint8
////////////////////////////////////////

func testBinaryUint8(t *testing.T, value int, buf []byte) {
	// ensure works for both io.Reader & io.Writer, and io.ByteReader & io.ByteWriter
	test := func(t *testing.T, value int, buf []byte, scratch testBuffer) {
		vin := Uint8(value)
		var vout Uint8

		if err := vin.MarshalBinaryTo(scratch); err != nil {
			t.Error(err)
		}

		if sb, ok := scratch.(testBytes); ok {
			if actual, expected := sb.Bytes(), buf; !bytes.Equal(actual, expected) {
				t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
			}
		}

		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			t.Error(err)
		}

		if actual, expected := vout, vin; actual != expected {
			t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}

	test(t, value, buf, new(buffer.Buffer))
	test(t, value, buf, new(bytes.Buffer))
}

func TestBinaryUint8(t *testing.T) {
	testBinaryUint8(t, 0, []byte{0x00})
	testBinaryUint8(t, 1, []byte{0x01})
	testBinaryUint8(t, 2, []byte{0x02})
	testBinaryUint8(t, 127, []byte{0x7f})
	testBinaryUint8(t, 128, []byte{0x80})
	testBinaryUint8(t, 129, []byte{0x81})
	testBinaryUint8(t, 254, []byte{0xfe})
	testBinaryUint8(t, 255, []byte{0xff})
}

////////////////////////////////////////
// Int16
////////////////////////////////////////

func testBinaryInt16(t *testing.T, value int, buf []byte) {
	vin := Int16(value)
	var vout Int16
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Int16(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryInt16(t *testing.T) {
	testBinaryInt16(t, 0, []byte("\x00\x00"))
	testBinaryInt16(t, 1, []byte("\x00\x01"))
	testBinaryInt16(t, -1, []byte("\xFF\xFF"))
	testBinaryInt16(t, 2, []byte("\x00\x02"))
	testBinaryInt16(t, -2, []byte("\xFF\xFE"))
}

////////////////////////////////////////
// Uint16
////////////////////////////////////////

func testBinaryUint16(t *testing.T, value int, buf []byte) {
	vin := Uint16(value)
	var vout Uint16
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Uint16(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryUint16(t *testing.T) {
	testBinaryUint16(t, 0, []byte("\x00\x00"))
	testBinaryUint16(t, 1, []byte("\x00\x01"))
	testBinaryUint16(t, 2, []byte("\x00\x02"))
	testBinaryUint16(t, 127, []byte("\x00\x7f"))
	testBinaryUint16(t, 128, []byte("\x00\x80"))
	testBinaryUint16(t, 129, []byte("\x00\x81"))
	testBinaryUint16(t, 16383, []byte("\x3f\xff"))
	testBinaryUint16(t, 16384, []byte("\x40\x00"))
	testBinaryUint16(t, 16385, []byte("\x40\x01"))
}

////////////////////////////////////////
// Int32
////////////////////////////////////////

func testBinaryInt32(t *testing.T, value int, buf []byte) {
	vin := Int32(value)
	var vout Int32
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Int32(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryInt32(t *testing.T) {
	testBinaryInt32(t, 0, []byte("\x00\x00\x00\x00"))
	testBinaryInt32(t, 1, []byte("\x00\x00\x00\x01"))
	testBinaryInt32(t, -1, []byte("\xFF\xFF\xFF\xFF"))
	testBinaryInt32(t, 2, []byte("\x00\x00\x00\x02"))
	testBinaryInt32(t, -2, []byte("\xFF\xFF\xFF\xFE"))
}

////////////////////////////////////////
// Uint32
////////////////////////////////////////

func testBinaryUint32(t *testing.T, value int, buf []byte) {
	vin := Uint32(value)
	var vout Uint32
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Uint32(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryUint32(t *testing.T) {
	testBinaryUint32(t, 0, []byte("\x00\x00\x00\x00"))
	testBinaryUint32(t, 1, []byte("\x00\x00\x00\x01"))
	testBinaryUint32(t, 2, []byte("\x00\x00\x00\x02"))
	testBinaryUint32(t, 127, []byte("\x00\x00\x00\x7f"))
	testBinaryUint32(t, 128, []byte("\x00\x00\x00\x80"))
	testBinaryUint32(t, 129, []byte("\x00\x00\x00\x81"))
	testBinaryUint32(t, 16383, []byte("\x00\x00\x3f\xff"))
	testBinaryUint32(t, 16384, []byte("\x00\x00\x40\x00"))
	testBinaryUint32(t, 16385, []byte("\x00\x00\x40\x01"))
}

////////////////////////////////////////
// Int64
////////////////////////////////////////

func testBinaryInt64(t *testing.T, value int, buf []byte) {
	vin := Int64(value)
	var vout Int64
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Int64(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryInt64(t *testing.T) {
	testBinaryInt64(t, 0, []byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
	testBinaryInt64(t, 1, []byte("\x00\x00\x00\x00\x00\x00\x00\x01"))
	testBinaryInt64(t, -1, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFF"))
	testBinaryInt64(t, 2, []byte("\x00\x00\x00\x00\x00\x00\x00\x02"))
	testBinaryInt64(t, -2, []byte("\xFF\xFF\xFF\xFF\xFF\xFF\xFF\xFE"))
}

////////////////////////////////////////
// Uint64
////////////////////////////////////////

func testBinaryUint64(t *testing.T, value int, buf []byte) {
	vin := Uint64(value)
	var vout Uint64
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Uint64(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryUint64(t *testing.T) {
	testBinaryUint64(t, 0, []byte("\x00\x00\x00\x00\x00\x00\x00\x00"))
	testBinaryUint64(t, 1, []byte("\x00\x00\x00\x00\x00\x00\x00\x01"))
	testBinaryUint64(t, 2, []byte("\x00\x00\x00\x00\x00\x00\x00\x02"))
	testBinaryUint64(t, 127, []byte("\x00\x00\x00\x00\x00\x00\x00\x7f"))
	testBinaryUint64(t, 128, []byte("\x00\x00\x00\x00\x00\x00\x00\x80"))
	testBinaryUint64(t, 129, []byte("\x00\x00\x00\x00\x00\x00\x00\x81"))
	testBinaryUint64(t, 16383, []byte("\x00\x00\x00\x00\x00\x00\x3f\xff"))
	testBinaryUint64(t, 16384, []byte("\x00\x00\x00\x00\x00\x00\x40\x00"))
	testBinaryUint64(t, 16385, []byte("\x00\x00\x00\x00\x00\x00\x40\x01"))
}

////////////////////////////////////////
// Float32
////////////////////////////////////////

func testBinaryFloat32(t *testing.T, value float64, buf []byte) {
	vin := Float32(value)
	var vout Float32
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Float32(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryFloat32(t *testing.T) {
	testBinaryFloat32(t, math.SmallestNonzeroFloat32, []byte("\x00\x00\x00\x01"))
	testBinaryFloat32(t, math.MaxFloat32, []byte("\x7f\x7f\xff\xff"))

	testBinaryFloat32(t, math.Sqrt2, []byte("\x3f\xb5\x04\xf3"))
	testBinaryFloat32(t, math.SqrtE, []byte("\x3f\xd3\x09\x4c"))
	testBinaryFloat32(t, math.SqrtPi, []byte("\x3f\xe2\xdf\xc5"))
	testBinaryFloat32(t, math.SqrtPhi, []byte("\x3f\xa2\xd1\x8a"))

	testBinaryFloat32(t, math.Ln2, []byte("\x3f\x31\x72\x18"))
	testBinaryFloat32(t, math.Log2E, []byte("\x3f\xb8\xaa\x3b"))
	testBinaryFloat32(t, math.Ln10, []byte("\x40\x13\x5d\x8e"))
	testBinaryFloat32(t, math.Log10E, []byte("\x3e\xde\x5b\xd9"))

	testBinaryFloat32(t, math.E, []byte("\x40\x2d\xf8\x54"))
	testBinaryFloat32(t, math.Phi, []byte("\x3f\xcf\x1b\xbd"))
	testBinaryFloat32(t, math.Pi, []byte("\x40\x49\x0f\xdb"))
}

////////////////////////////////////////
// Float64
////////////////////////////////////////

func testBinaryFloat64(t *testing.T, value float64, buf []byte) {
	vin := Float64(value)
	var vout Float64
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := vout, Float64(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryFloat64(t *testing.T) {
	testBinaryFloat64(t, math.SmallestNonzeroFloat64, []byte("\x00\x00\x00\x00\x00\x00\x00\x01"))
	testBinaryFloat64(t, math.MaxFloat64, []byte("\x7f\xef\xff\xff\xff\xff\xff\xff"))

	testBinaryFloat64(t, math.Sqrt2, []byte("\x3f\xf6\xa0\x9e\x66\x7f\x3b\xcd"))
	testBinaryFloat64(t, math.SqrtE, []byte("\x3f\xfa\x61\x29\x8e\x1e\x06\x9c"))
	testBinaryFloat64(t, math.SqrtPi, []byte("\x3f\xfc\x5b\xf8\x91\xb4\xef\x6b"))
	testBinaryFloat64(t, math.SqrtPhi, []byte("\x3f\xf4\x5a\x31\x46\xa8\x84\x56"))

	testBinaryFloat64(t, math.Ln2, []byte("\x3f\xe6\x2e\x42\xfe\xfa\x39\xef"))
	testBinaryFloat64(t, math.Log2E, []byte("\x3f\xf7\x15\x47\x65\x2b\x82\xfe"))
	testBinaryFloat64(t, math.Ln10, []byte("\x40\x02\x6b\xb1\xbb\xb5\x55\x16"))
	testBinaryFloat64(t, math.Log10E, []byte("\x3f\xdb\xcb\x7b\x15\x26\xe5\x0e"))

	testBinaryFloat64(t, math.E, []byte("\x40\x05\xbf\x0a\x8b\x14\x57\x69"))
	testBinaryFloat64(t, math.Phi, []byte("\x3f\xf9\xe3\x77\x9b\x97\xf4\xa8"))
	testBinaryFloat64(t, math.Pi, []byte("\x40\x09\x21\xfb\x54\x44\x2d\x18"))
}

////////////////////////////////////////
// VWI -- signed variable width integer
////////////////////////////////////////

func testBinaryVWI(t *testing.T, value int, buf []byte) {
	// ensure works for both io.Reader & io.Writer, and io.ByteReader & io.ByteWriter
	test := func(t *testing.T, value int, buf []byte, scratch testBuffer) {
		vin := VWI(value)
		var vout VWI

		if err := vin.MarshalBinaryTo(scratch); err != nil {
			t.Error(err)
		}

		if sb, ok := scratch.(testBytes); ok {
			if actual, expected := sb.Bytes(), buf; !bytes.Equal(actual, expected) {
				t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
			}
		}

		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			t.Error(err)
		}

		if actual, expected := vout, vin; actual != expected {
			t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}

	test(t, value, buf, new(buffer.Buffer))
	test(t, value, buf, new(bytes.Buffer))
}

func TestBinaryVWIOneByte(t *testing.T) {
	testBinaryVWI(t, 0, []byte("\x00"))
	testBinaryVWI(t, -1, []byte("\x01"))
	testBinaryVWI(t, 1, []byte("\x02"))
	testBinaryVWI(t, -2, []byte("\x03"))
	testBinaryVWI(t, 2, []byte("\x04"))
	testBinaryVWI(t, -3, []byte("\x05"))
	testBinaryVWI(t, 3, []byte("\x06"))
	testBinaryVWI(t, -63, []byte("\x7d"))
	testBinaryVWI(t, 63, []byte("\x7e"))
	testBinaryVWI(t, -64, []byte("\x7f"))
}

func TestBinaryVWIMultipleBytes(t *testing.T) {
	testBinaryVWI(t, 64, []byte("\x80\x01"))
	testBinaryVWI(t, -65, []byte("\x81\x01"))
	testBinaryVWI(t, 65, []byte("\x82\x01"))
	testBinaryVWI(t, 2147483647, []byte("\xfe\xff\xff\xff\x0f"))
	testBinaryVWI(t, -2147483648, []byte("\xff\xff\xff\xff\x0f"))
	testBinaryVWI(t, 1082196484, []byte("\x88\x88\x88\x88\x08"))
	testBinaryVWI(t, 138521149956, []byte("\x88\x88\x88\x88\x88\x08"))
	testBinaryVWI(t, 17730707194372, []byte("\x88\x88\x88\x88\x88\x88\x08"))
	testBinaryVWI(t, 2269530520879620, []byte("\x88\x88\x88\x88\x88\x88\x88\x08"))
}

////////////////////////////////////////
// UVWI -- unsigned variable width integer
////////////////////////////////////////

func testBinaryUVWI(t *testing.T, value int, buf []byte) {
	// ensure works for both io.Reader & io.Writer, and io.ByteReader & io.ByteWriter
	test := func(t *testing.T, value int, buf []byte, scratch testBuffer) {
		vin := UVWI(value)
		var vout UVWI

		if err := vin.MarshalBinaryTo(scratch); err != nil {
			t.Error(err)
		}

		if sb, ok := scratch.(testBytes); ok {
			if actual, expected := sb.Bytes(), buf; !bytes.Equal(actual, expected) {
				t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
			}
		}

		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			t.Error(err)
		}

		if actual, expected := vout, vin; actual != expected {
			t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}

	test(t, value, buf, new(buffer.Buffer))
	test(t, value, buf, new(bytes.Buffer))
}

func TestBinaryUVWIOneByte(t *testing.T) {
	testBinaryUVWI(t, 0, []byte("\x00"))
	testBinaryUVWI(t, 1, []byte("\x01"))
	testBinaryUVWI(t, 2, []byte("\x02"))
	testBinaryUVWI(t, 127, []byte("\x7f"))
}

func TestBinaryUVWIMultipleBytes(t *testing.T) {
	testBinaryUVWI(t, 0x7F, []byte("\x7F"))
	testBinaryUVWI(t, 0x80, []byte("\x80\x01"))
	testBinaryUVWI(t, 0x81, []byte("\x81\x01"))

	testBinaryUVWI(t, 0x00003FFF, []byte("\xFF\x7F"))
	testBinaryUVWI(t, 0x00004000, []byte("\x80\x80\x01"))
	testBinaryUVWI(t, 0x00004001, []byte("\x81\x80\x01"))

	testBinaryUVWI(t, 0x001FFFFF, []byte("\xFF\xFF\x7F"))
	testBinaryUVWI(t, 0x00200000, []byte("\x80\x80\x80\x01"))
	testBinaryUVWI(t, 0x00200001, []byte("\x81\x80\x80\x01"))

	testBinaryUVWI(t, 0x0FFFFFFF, []byte("\xFF\xFF\xFF\x7F"))
	testBinaryUVWI(t, 0x10000000, []byte("\x80\x80\x80\x80\x01"))
	testBinaryUVWI(t, 0x10000001, []byte("\x81\x80\x80\x80\x01"))
}

////////////////////////////////////////

func benchmarkCodec(b *testing.B, scratch testBuffer, vin, vout Binary) {
	for i := 0; i < b.N; i++ {
		if err := vin.MarshalBinaryTo(scratch); err != nil {
			b.Fatal(err)
		}
		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			b.Fatal(err)
		}
	}
}

////////////////////////////////////////
// relative difference between using io.ByteReader and io.ByteWriter vs. io.Reader and io.Writer.

func BenchmarkBinaryBinaryBuffer(b *testing.B) {
	scratch := new(buffer.Buffer)
	vin := UVWI(0xFFFFFFFFFFFFFFFF)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryBinaryBytes(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(0xFFFFFFFFFFFFFFFF)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

////////////////////////////////////////
// relative differences between integers of different lengths

func BenchmarkBinaryOneByteFixed(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := Uint8(1)
	var vout Uint8
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryOneByteVariable(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(1)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryTwoBytesFixed(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := Uint16(0x80)
	var vout Uint16
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryTwoBytesVariable(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(0x80)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryThreeBytesFixed(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := Uint32(0x00004000) // must use 4 bytes fixed to encode 3 bytes
	var vout Uint32
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryThreeBytesVariable(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(0x00004000)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryFourBytesFixed(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := Uint32(0x00200000)
	var vout Uint32
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryFourBytesVariable(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(0x00200000)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryEightBytesFixed(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := Uint64(0xFFFFFFFFFFFFFFFF)
	var vout Uint64
	benchmarkCodec(b, scratch, &vin, &vout)
}

func BenchmarkBinaryEightBytesVariable(b *testing.B) {
	scratch := new(bytes.Buffer)
	vin := UVWI(0xFFFFFFFFFFFFFFFF)
	var vout UVWI
	benchmarkCodec(b, scratch, &vin, &vout)
}

////////////////////////////////////////
// String
////////////////////////////////////////

func testBinaryString(t *testing.T, value string, buf []byte) {
	vin := String(value)
	var vout String
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := string(vout), value; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryString(t *testing.T) {
	testBinaryString(t, "", []byte("\x00"))
	testBinaryString(t, "short", []byte("\x05short"))
	testBinaryString(t, "this is a slightly longer message",
		[]byte("\x21this is a slightly longer message"))
}

////////////////////////////////////////
// StringSlice
////////////////////////////////////////

func testBinaryStringSlice(t *testing.T, value []String, buf []byte) {
	vin := StringSlice(value)
	var vout StringSlice
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := len(vout), len(vin); actual != expected {
		t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
	}
	for i := 0; i < len(vout); i++ {
		if actual, expected := vout[i], vin[i]; actual != expected {
			t.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}
}

func TestBinaryStringSlice(t *testing.T) {
	testBinaryStringSlice(t, StringSlice{}, []byte("\x00"))
	testBinaryStringSlice(t, StringSlice{String("one")}, []byte("\x01\x03one"))
	testBinaryStringSlice(t, StringSlice{String("one"), String("two")},
		[]byte("\x02\x03one\x03two"))
}
