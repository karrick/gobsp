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

func testInt8(t *testing.T, value int, buf []byte) {
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

func TestBinaryInt8Codec(t *testing.T) {
	testInt8(t, 0, []byte{0x00})
	testInt8(t, 1, []byte{0x01})
	testInt8(t, -1, []byte{0xFF})
	testInt8(t, 2, []byte{0x02})
	testInt8(t, -2, []byte{0xFE})
	testInt8(t, 127, []byte{0x7f})
	testInt8(t, -127, []byte{0x81})
}

////////////////////////////////////////
// Uint8
////////////////////////////////////////

func testUint8(t *testing.T, value int, buf []byte) {
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

func TestBinaryUint8Codec(t *testing.T) {
	testUint8(t, 0, []byte{0x00})
	testUint8(t, 1, []byte{0x01})
	testUint8(t, 2, []byte{0x02})
	testUint8(t, 127, []byte{0x7f})
	testUint8(t, 128, []byte{0x80})
	testUint8(t, 129, []byte{0x81})
	testUint8(t, 254, []byte{0xfe})
	testUint8(t, 255, []byte{0xff})
}

////////////////////////////////////////
// Int16
////////////////////////////////////////

func testInt16(t *testing.T, value int, buf []byte) {
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

func TestBinaryINT16Codec(t *testing.T) {
	testInt16(t, 0, []byte{0x00, 0x00})
	testInt16(t, 1, []byte{0x00, 0x01})
	testInt16(t, -1, []byte{0xFF, 0xFF})
	testInt16(t, 2, []byte{0x00, 0x02})
	testInt16(t, -2, []byte{0xFF, 0xFE})
}

////////////////////////////////////////
// Uint16
////////////////////////////////////////

func testUint16(t *testing.T, value int, buf []byte) {
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

func TestBinaryUINT16Codec(t *testing.T) {
	testUint16(t, 0, []byte{0x00, 0x00})
	testUint16(t, 1, []byte{0x00, 0x01})
	testUint16(t, 2, []byte{0x00, 0x02})
	testUint16(t, 127, []byte{0x00, 0x7f})
	testUint16(t, 128, []byte{0x00, 0x80})
	testUint16(t, 129, []byte{0x00, 0x81})
	testUint16(t, 16383, []byte{0x3f, 0xff})
	testUint16(t, 16384, []byte{0x40, 0x00})
	testUint16(t, 16385, []byte{0x40, 0x01})
}

////////////////////////////////////////
// Int32
////////////////////////////////////////

func testInt32(t *testing.T, value int, buf []byte) {
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

func TestBinaryINT32Codec(t *testing.T) {
	testInt32(t, 0, []byte{0x00, 0x00, 0x00, 0x00})
	testInt32(t, 1, []byte{0x00, 0x00, 0x00, 0x01})
	testInt32(t, -1, []byte{0xFF, 0xFF, 0xFF, 0xFF})
	testInt32(t, 2, []byte{0x00, 0x00, 0x00, 0x02})
	testInt32(t, -2, []byte{0xFF, 0xFF, 0xFF, 0xFE})
}

////////////////////////////////////////
// Uint32
////////////////////////////////////////

func testUint32(t *testing.T, value int, buf []byte) {
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

func TestBinaryUINT32Codec(t *testing.T) {
	testUint32(t, 0, []byte{0x00, 0x00, 0x00, 0x00})
	testUint32(t, 1, []byte{0x00, 0x00, 0x00, 0x01})
	testUint32(t, 2, []byte{0x00, 0x00, 0x00, 0x02})
	testUint32(t, 127, []byte{0x00, 0x00, 0x00, 0x7f})
	testUint32(t, 128, []byte{0x00, 0x00, 0x00, 0x80})
	testUint32(t, 129, []byte{0x00, 0x00, 0x00, 0x81})
	testUint32(t, 16383, []byte{0x00, 0x00, 0x3f, 0xff})
	testUint32(t, 16384, []byte{0x00, 0x00, 0x40, 0x00})
	testUint32(t, 16385, []byte{0x00, 0x00, 0x40, 0x01})
}

////////////////////////////////////////
// Int64
////////////////////////////////////////

func testInt64(t *testing.T, value int, buf []byte) {
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

func TestBinaryINT64Codec(t *testing.T) {
	testInt64(t, 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	testInt64(t, 1, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	testInt64(t, -1, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	testInt64(t, 2, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
	testInt64(t, -2, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE})
}

////////////////////////////////////////
// Uint64
////////////////////////////////////////

func testUint64(t *testing.T, value int, buf []byte) {
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

func TestBinaryUINT64Codec(t *testing.T) {
	testUint64(t, 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	testUint64(t, 1, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	testUint64(t, 2, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
	testUint64(t, 127, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7f})
	testUint64(t, 128, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80})
	testUint64(t, 129, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81})
	testUint64(t, 16383, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xff})
	testUint64(t, 16384, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00})
	testUint64(t, 16385, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x01})
}

////////////////////////////////////////
// Float32
////////////////////////////////////////

func testFloat32(t *testing.T, value float64, buf []byte) {
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

func TestBinaryFLOAT32Codec(t *testing.T) {
	testFloat32(t, math.SmallestNonzeroFloat32, []byte{0x0, 0x0, 0x0, 0x1})
	testFloat32(t, math.MaxFloat32, []byte{0x7f, 0x7f, 0xff, 0xff})

	testFloat32(t, math.Sqrt2, []byte{0x3f, 0xb5, 0x4, 0xf3})
	testFloat32(t, math.SqrtE, []byte{0x3f, 0xd3, 0x9, 0x4c})
	testFloat32(t, math.SqrtPi, []byte{0x3f, 0xe2, 0xdf, 0xc5})
	testFloat32(t, math.SqrtPhi, []byte{0x3f, 0xa2, 0xd1, 0x8a})

	testFloat32(t, math.Ln2, []byte{0x3f, 0x31, 0x72, 0x18})
	testFloat32(t, math.Log2E, []byte{0x3f, 0xb8, 0xaa, 0x3b})
	testFloat32(t, math.Ln10, []byte{0x40, 0x13, 0x5d, 0x8e})
	testFloat32(t, math.Log10E, []byte{0x3e, 0xde, 0x5b, 0xd9})

	testFloat32(t, math.E, []byte{0x40, 0x2d, 0xf8, 0x54})
	testFloat32(t, math.Phi, []byte{0x3f, 0xcf, 0x1b, 0xbd})
	testFloat32(t, math.Pi, []byte{0x40, 0x49, 0xf, 0xdb})
}

////////////////////////////////////////
// Float64
////////////////////////////////////////

func testFloat64(t *testing.T, value float64, buf []byte) {
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

func TestBinaryFLOAT64Codec(t *testing.T) {
	testFloat64(t, math.SmallestNonzeroFloat64, []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1})
	testFloat64(t, math.MaxFloat64, []byte{0x7f, 0xef, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})

	testFloat64(t, math.Sqrt2, []byte{0x3f, 0xf6, 0xa0, 0x9e, 0x66, 0x7f, 0x3b, 0xcd})
	testFloat64(t, math.SqrtE, []byte{0x3f, 0xfa, 0x61, 0x29, 0x8e, 0x1e, 0x6, 0x9c})
	testFloat64(t, math.SqrtPi, []byte{0x3f, 0xfc, 0x5b, 0xf8, 0x91, 0xb4, 0xef, 0x6b})
	testFloat64(t, math.SqrtPhi, []byte{0x3f, 0xf4, 0x5a, 0x31, 0x46, 0xa8, 0x84, 0x56})

	testFloat64(t, math.Ln2, []byte{0x3f, 0xe6, 0x2e, 0x42, 0xfe, 0xfa, 0x39, 0xef})
	testFloat64(t, math.Log2E, []byte{0x3f, 0xf7, 0x15, 0x47, 0x65, 0x2b, 0x82, 0xfe})
	testFloat64(t, math.Ln10, []byte{0x40, 0x2, 0x6b, 0xb1, 0xbb, 0xb5, 0x55, 0x16})
	testFloat64(t, math.Log10E, []byte{0x3f, 0xdb, 0xcb, 0x7b, 0x15, 0x26, 0xe5, 0xe})

	testFloat64(t, math.E, []byte{0x40, 0x5, 0xbf, 0xa, 0x8b, 0x14, 0x57, 0x69})
	testFloat64(t, math.Phi, []byte{0x3f, 0xf9, 0xe3, 0x77, 0x9b, 0x97, 0xf4, 0xa8})
	testFloat64(t, math.Pi, []byte{0x40, 0x9, 0x21, 0xfb, 0x54, 0x44, 0x2d, 0x18})
}

////////////////////////////////////////
// VWI -- signed variable width integer
////////////////////////////////////////

func testVWICodec(t *testing.T, value int, buf []byte) {
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

func TestBinaryVWICodecOneByte(t *testing.T) {
	testVWICodec(t, 0, []byte{0x00})
	testVWICodec(t, -1, []byte{0x01})
	testVWICodec(t, 1, []byte{0x02})
	testVWICodec(t, -2, []byte{0x03})
	testVWICodec(t, 2, []byte{0x04})
	testVWICodec(t, -3, []byte{0x05})
	testVWICodec(t, 3, []byte{0x06})
}

func TestBinaryVWICodecMultipleBytes(t *testing.T) {
	testVWICodec(t, 64, []byte{0x80, 0x01})
	testVWICodec(t, -65, []byte{0x81, 0x01})

	testVWICodec(t, 2147483647, []byte("\xfe\xff\xff\xff\x0f"))
	testVWICodec(t, -2147483648, []byte("\xff\xff\xff\xff\x0f"))
	testVWICodec(t, 1082196484, []byte("\x88\x88\x88\x88\x08"))
	testVWICodec(t, 138521149956, []byte("\x88\x88\x88\x88\x88\x08"))
	testVWICodec(t, 17730707194372, []byte("\x88\x88\x88\x88\x88\x88\x08"))
	testVWICodec(t, 2269530520879620, []byte("\x88\x88\x88\x88\x88\x88\x88\x08"))
}

////////////////////////////////////////
// UVWI -- unsigned variable width integer
////////////////////////////////////////

func testUVWICodec(t *testing.T, value int, buf []byte) {
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

func TestBinaryUVWICodecOneByte(t *testing.T) {
	testUVWICodec(t, 0, []byte{0x00})
	testUVWICodec(t, 1, []byte{0x01})
	testUVWICodec(t, 2, []byte{0x02})
	testUVWICodec(t, 127, []byte{0x7f})
}

func TestBinaryUVWICodecMultipleBytes(t *testing.T) {
	testUVWICodec(t, 0x7F, []byte{0x7F})
	testUVWICodec(t, 0x80, []byte{0x80, 0x01})
	testUVWICodec(t, 0x81, []byte{0x81, 0x01})

	testUVWICodec(t, 0x00003FFF, []byte{0xFF, 0x7F})
	testUVWICodec(t, 0x00004000, []byte{0x80, 0x80, 0x01})
	testUVWICodec(t, 0x00004001, []byte{0x81, 0x80, 0x01})

	testUVWICodec(t, 0x001FFFFF, []byte{0xFF, 0xFF, 0x7F})
	testUVWICodec(t, 0x00200000, []byte{0x80, 0x80, 0x80, 0x01})
	testUVWICodec(t, 0x00200001, []byte{0x81, 0x80, 0x80, 0x01})

	testUVWICodec(t, 0x0FFFFFFF, []byte{0xFF, 0xFF, 0xFF, 0x7F})
	testUVWICodec(t, 0x10000000, []byte{0x80, 0x80, 0x80, 0x80, 0x01})
	testUVWICodec(t, 0x10000001, []byte{0x81, 0x80, 0x80, 0x80, 0x01})
}

func TestBinaryUVWICodecSigned(t *testing.T) {
	// testInt8(t, 0, []byte{0x00})
	// testInt8(t, -1, []byte{0x01})
	// testInt8(t, 1, []byte{0x02})
	// testInt8(t, -2, []byte{0x03})
	// testInt8(t, 2, []byte{0x04})
	// testInt8(t, -3, []byte{0x05})
	// testInt8(t, 3, []byte{0x06})
}

func BenchmarkUVWIBuffer(b *testing.B) {
	benchmarkUVWI(b, new(buffer.Buffer))
}

func BenchmarkUVWIBytes(b *testing.B) {
	benchmarkUVWI(b, new(bytes.Buffer))
}

func benchmarkUVWI(b *testing.B, scratch testBuffer) {
	const largeUVWIValue = 0x0FFFFFFF
	vin := UVWI(largeUVWIValue)
	var vout UVWI

	for i := 0; i < b.N; i++ {
		if err := vin.MarshalBinaryTo(scratch); err != nil {
			b.Fatal(err)
		}
		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			b.Fatal(err)
		}
		if actual, expected := vout, vin; actual != expected {
			b.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}
}

////////////////////////////////////////

func BenchmarkTwoBytesUINT16(b *testing.B) {
	vin := Uint16(0x00003FFF) // largest integer that can be stored in 2 bytes UVWI
	var vout Uint16
	scratch := new(buffer.Buffer) // ensure not using ByteReader ByteWriter optimization

	for i := 0; i < b.N; i++ {
		if err := vin.MarshalBinaryTo(scratch); err != nil {
			b.Fatal(err)
		}
		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			b.Fatal(err)
		}
		if actual, expected := vout, vin; actual != expected {
			b.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}
}

func BenchmarkTwoBytesUVWI(b *testing.B) {
	vin := UVWI(0x00003FFF) // largest integer that can be stored in 2 bytes UVWI
	var vout UVWI
	scratch := new(buffer.Buffer) // ensure not using ByteReader ByteWriter optimization

	for i := 0; i < b.N; i++ {
		if err := vin.MarshalBinaryTo(scratch); err != nil {
			b.Fatal(err)
		}
		if err := vout.UnmarshalBinaryFrom(scratch); err != nil {
			b.Fatal(err)
		}
		if actual, expected := vout, vin; actual != expected {
			b.Fatalf("Actual: %#v; Expected: %#v", actual, expected)
		}
	}
}

////////////////////////////////////////
// String
////////////////////////////////////////

func testStringCodec(t *testing.T, value string, buf []byte) {
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

func TestBinaryStringCodec(t *testing.T) {
	testStringCodec(t, "", []byte{0x0})
	testStringCodec(t, "short", []byte{0x05, 's', 'h', 'o', 'r', 't'})
	testStringCodec(t, "this is a slightly longer message",
		append([]byte{0x21}, []byte("this is a slightly longer message")...))
}

////////////////////////////////////////
// StringSlice
////////////////////////////////////////

func testStringSliceCodec(t *testing.T, value []String, buf []byte) {
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

func TestBinaryStringSliceCodec(t *testing.T) {
	testStringSliceCodec(t, StringSlice{}, []byte{0x0})
	testStringSliceCodec(t, StringSlice{String("one")},
		append([]byte{0x1, 0x3}, []byte("one")...))
	testStringSliceCodec(t, StringSlice{String("one"), String("two")},
		[]byte{
			0x2,
			0x3, 'o', 'n', 'e',
			0x3, 't', 'w', 'o',
		})
}
