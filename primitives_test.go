package guanoloco

import (
	"bytes"
	"math"
	"testing"
)

func testUINT32Codec(t *testing.T, value uint64, buf []byte) {
	testUINT32Encode(t, value, buf)
	testUINT32Decode(t, value, buf)
}

func testUINT32Encode(t *testing.T, value uint64, buf []byte) {
	v := Uint32(value)
	bb := new(bytes.Buffer)

	if err := v.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func testUINT32Decode(t *testing.T, value uint64, buf []byte) {
	bb := bytes.NewBuffer(buf)
	var v Uint32

	if err := v.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := v, Uint32(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryUINT32Codec(t *testing.T) {
	testUINT32Codec(t, 0, []byte{0x00, 0x00, 0x00, 0x00})
	testUINT32Codec(t, 1, []byte{0x00, 0x00, 0x00, 0x01})
	testUINT32Codec(t, 2, []byte{0x00, 0x00, 0x00, 0x02})
	testUINT32Codec(t, 127, []byte{0x00, 0x00, 0x00, 0x7f})
	testUINT32Codec(t, 128, []byte{0x00, 0x00, 0x00, 0x80})
	testUINT32Codec(t, 129, []byte{0x00, 0x00, 0x00, 0x81})
	testUINT32Codec(t, 16383, []byte{0x00, 0x00, 0x3f, 0xff})
	testUINT32Codec(t, 16384, []byte{0x00, 0x00, 0x40, 0x00})
	testUINT32Codec(t, 16385, []byte{0x00, 0x00, 0x40, 0x01})
}

////////////////////////////////////////

func testUINT64Codec(t *testing.T, value uint64, buf []byte) {
	testUINT64Encode(t, value, buf)
	testUINT64Decode(t, value, buf)
}

func testUINT64Encode(t *testing.T, value uint64, buf []byte) {
	v := Uint64(value)
	bb := new(bytes.Buffer)

	if err := v.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func testUINT64Decode(t *testing.T, value uint64, buf []byte) {
	bb := bytes.NewBuffer(buf)
	var v Uint64

	if err := v.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := v, Uint64(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryUINT64Codec(t *testing.T) {
	testUINT64Codec(t, 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	testUINT64Codec(t, 1, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01})
	testUINT64Codec(t, 2, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02})
	testUINT64Codec(t, 127, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7f})
	testUINT64Codec(t, 128, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80})
	testUINT64Codec(t, 129, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x81})
	testUINT64Codec(t, 16383, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xff})
	testUINT64Codec(t, 16384, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00})
	testUINT64Codec(t, 16385, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x01})
}

////////////////////////////////////////

func testFLOAT64Codec(t *testing.T, value float64) {
	vin := Float64(value)
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Fatal(err)
	}

	var vout Float64

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := float64(vout), value; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryFLOAT64Codec(t *testing.T) {
	testFLOAT64Codec(t, math.SmallestNonzeroFloat64)
	testFLOAT64Codec(t, math.MaxFloat64)

	testFLOAT64Codec(t, math.Sqrt2)
	testFLOAT64Codec(t, math.SqrtE)
	testFLOAT64Codec(t, math.SqrtPi)
	testFLOAT64Codec(t, math.SqrtPhi)

	testFLOAT64Codec(t, math.Ln2)
	testFLOAT64Codec(t, math.Log2E)
	testFLOAT64Codec(t, math.Ln10)
	testFLOAT64Codec(t, math.Log10E)

	testFLOAT64Codec(t, math.E)
	testFLOAT64Codec(t, math.Phi)
	testFLOAT64Codec(t, math.Pi)
}

////////////////////////////////////////

func testVWICodec(t *testing.T, value uint64, buf []byte) {
	testVWIEncode(t, value, buf)
	testVWIDecode(t, value, buf)
}

func testVWIEncode(t *testing.T, value uint64, buf []byte) {
	v := VWI(value)
	bb := new(bytes.Buffer)

	if err := v.MarshalBinaryTo(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := bb.Bytes(), buf; !bytes.Equal(actual, expected) {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func testVWIDecode(t *testing.T, value uint64, buf []byte) {
	bb := bytes.NewBuffer(buf)
	var v VWI

	if err := v.UnmarshalBinaryFrom(bb); err != nil {
		t.Error(err)
	}

	if actual, expected := v, VWI(value); actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryVWICodecOneByte(t *testing.T) {
	testVWICodec(t, 0, []byte{0x00})
	testVWICodec(t, 1, []byte{0x01})
	testVWICodec(t, 2, []byte{0x02})
	testVWICodec(t, 127, []byte{0x7f})
}

func TestBinaryVWICodecTwoBytes(t *testing.T) {
	testVWICodec(t, 128, []byte{0x80, 0x01})
	testVWICodec(t, 129, []byte{0x81, 0x01})
	testVWICodec(t, 16383, []byte{0xff, 0x7f})
	testVWICodec(t, 16384, []byte{0x80, 0x80, 0x01})
	testVWICodec(t, 16385, []byte{0x81, 0x80, 0x01})
}

////////////////////////////////////////

func testStringCodec(t *testing.T, value string) {
	vin := String(value)
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Fatal(err)
	}

	var vout String

	if err := vout.UnmarshalBinaryFrom(bb); err != nil {
		t.Fatal(err)
	}

	if actual, expected := string(vout), value; actual != expected {
		t.Errorf("Actual: %#v; Expected: %#v", actual, expected)
	}
}

func TestBinaryStringCodec(t *testing.T) {
	testStringCodec(t, "")
	testStringCodec(t, "short")
	testStringCodec(t, "this is a slightly longer message")
}

////////////////////////////////////////

func testStringSliceCodec(t *testing.T, value []String) {
	vin := StringSlice(value)
	bb := new(bytes.Buffer)

	if err := vin.MarshalBinaryTo(bb); err != nil {
		t.Fatal(err)
	}

	var vout StringSlice

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
	testStringSliceCodec(t, StringSlice{})
	testStringSliceCodec(t, StringSlice{String("one")})
	testStringSliceCodec(t, StringSlice{String("one"), String("two")})
}
