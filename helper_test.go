package live

import (
	"bytes"
	"github.com/edwingeng/live/internal"
	"math"
	"testing"
)

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64

	MaxInt    = 1<<(intSize-1) - 1
	MinInt    = -1 << (intSize - 1)
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint   = 1<<intSize - 1
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)

func rec() {
	_ = recover()
}

func TestWrapBool(t *testing.T) {
	a := []bool{false, true}
	for _, v := range a {
		d := WrapBool(v)
		if d.Bool() != v {
			t.Fatal("d.Bool() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Bool()
		t.Fatal("d.Bool() should panic")
	}()
}

func TestWrapInt(t *testing.T) {
	a := []int{MinInt, -1, 0, 1, 100, MaxInt}
	for _, v := range a {
		d := WrapInt(v)
		if d.Int() != v {
			t.Fatal("d.Int() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Int()
		t.Fatal("d.Int() should panic")
	}()
}

func TestWrapInt8(t *testing.T) {
	a := []int8{MinInt8, -1, 0, 1, 100, MaxInt8}
	for _, v := range a {
		d := WrapInt8(v)
		if d.Int8() != v {
			t.Fatal("d.Int8() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Int8()
		t.Fatal("d.Int8() should panic")
	}()
}

func TestWrapInt16(t *testing.T) {
	a := []int16{MinInt16, -1, 0, 1, 100, MaxInt16}
	for _, v := range a {
		d := WrapInt16(v)
		if d.Int16() != v {
			t.Fatal("d.Int16() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Int16()
		t.Fatal("d.Int16() should panic")
	}()
}

func TestWrapInt32(t *testing.T) {
	a := []int32{MinInt32, -1, 0, 1, 100, MaxInt32}
	for _, v := range a {
		d := WrapInt32(v)
		if d.Int32() != v {
			t.Fatal("d.Int32() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Int32()
		t.Fatal("d.Int32() should panic")
	}()
}

func TestWrapInt64(t *testing.T) {
	a := []int64{MinInt64, -1, 0, 1, 100, MaxInt64}
	for _, v := range a {
		d := WrapInt64(v)
		if d.Int64() != v {
			t.Fatal("d.Int64() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Int64()
		t.Fatal("d.Int64() should panic")
	}()
}

func TestWrapUint(t *testing.T) {
	a := []uint{0, 1, 100, MaxUint}
	for _, v := range a {
		d := WrapUint(v)
		if d.Uint() != v {
			t.Fatal("d.Uint() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Uint()
		t.Fatal("d.Uint() should panic")
	}()
}

func TestWrapUint8(t *testing.T) {
	a := []uint8{0, 1, 100, MaxUint8}
	for _, v := range a {
		d := WrapUint8(v)
		if d.Uint8() != v {
			t.Fatal("d.Uint8() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Uint8()
		t.Fatal("d.Uint8() should panic")
	}()
}

func TestWrapUint16(t *testing.T) {
	a := []uint16{0, 1, 100, MaxUint16}
	for _, v := range a {
		d := WrapUint16(v)
		if d.Uint16() != v {
			t.Fatal("d.Uint16() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Uint16()
		t.Fatal("d.Uint16() should panic")
	}()
}

func TestWrapUint32(t *testing.T) {
	a := []uint32{0, 1, 100, MaxUint32}
	for _, v := range a {
		d := WrapUint32(v)
		if d.Uint32() != v {
			t.Fatal("d.Uint32() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Uint32()
		t.Fatal("d.Uint32() should panic")
	}()
}

func TestWrapUint64(t *testing.T) {
	a := []uint64{0, 1, 100, MaxUint64}
	for _, v := range a {
		d := WrapUint64(v)
		if d.Uint64() != v {
			t.Fatal("d.Uint64() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Uint64()
		t.Fatal("d.Uint64() should panic")
	}()
}

func TestWrapFloat32(t *testing.T) {
	a := []float32{math.SmallestNonzeroFloat32, -9.9, -1, 0, 1, 100, 1000.28, math.MaxFloat32}
	for _, v := range a {
		d := WrapFloat32(v)
		if d.Float32() != v {
			t.Fatal("d.Float32() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Float32()
		t.Fatal("d.Float32() should panic")
	}()
}

func TestWrapFloat64(t *testing.T) {
	a := []float64{math.SmallestNonzeroFloat64, -9.9, -1, 0, 1, 100, 1000.28, math.MaxFloat64}
	for _, v := range a {
		d := WrapFloat64(v)
		if d.Float64() != v {
			t.Fatal("d.Float64() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer rec()
		var d Data
		d.Float64()
		t.Fatal("d.Float64() should panic")
	}()
}

func TestWrapString(t *testing.T) {
	a := []string{"", "hello", "it is a good day to die"}
	for _, v := range a {
		d := WrapString(v)
		if d.String() != v {
			t.Fatal("d.String() != v")
		}
		if d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0")
		}
	}

	func() {
		defer rec()
		var d Data
		_ = d.String()
		t.Fatal("d.String() should panic")
	}()
}

func TestWrapBytes(t *testing.T) {
	a := []string{"", "hello", "it is a good day to die"}
	for _, v := range a {
		v := []byte(v)
		d := WrapBytes(v)
		if !bytes.Equal(d.Bytes(), v) {
			t.Fatal("!bytes.Equal(d.Bytes(), v)")
		}
		if d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0")
		}
	}

	if WrapBytes(nil).Bytes() != nil {
		t.Fatal("WrapBytes(nil).Bytes() != nil")
	}

	func() {
		defer rec()
		_ = Nil.Bytes()
		t.Fatal("Nil.Bytes() should panic")
	}()
}

func TestWrapComplex64(t *testing.T) {
	a := []float32{0, 0, 1, -1, math.Pi, math.E, math.MaxFloat32, math.SmallestNonzeroFloat32}
	for i := 0; i < len(a); i += 2 {
		d := WrapComplex64(complex(a[i], a[i+1]))
		v := d.Complex64()
		if real(v) != a[i] || imag(v) != a[i+1] {
			t.Fatal(`real(v) != a[i] || imag(v) != a[i+1]`)
		}
		if d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0")
		}
	}

	func() {
		defer rec()
		var d Data
		_ = d.Complex64()
		t.Fatal("d.Complex64() should panic")
	}()
}

func TestWrapComplex128(t *testing.T) {
	a := []float64{0, 0, 1, -1, math.Pi, math.E, math.MaxFloat64, math.SmallestNonzeroFloat64}
	for i := 0; i < len(a); i += 2 {
		d := WrapComplex128(complex(a[i], a[i+1]))
		v := d.Complex128()
		if real(v) != a[i] || imag(v) != a[i+1] {
			t.Fatal(`real(v) != a[i] || imag(v) != a[i+1]`)
		}
		if d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil || d.v.(*internal.Data).N != 0")
		}
	}

	func() {
		defer rec()
		var d Data
		_ = d.Complex128()
		t.Fatal("d.Complex128() should panic")
	}()
}

func TestMustWrapObject(t *testing.T) {
	var obj1, obj2 struct {
		A int64
		B string
	}
	obj1.A = 100
	obj2.B = "hello"

	MustWrapObject(&obj1).MustUnwrapObject(&obj2)
	if obj2 != obj1 {
		t.Fatal("obj2 != obj1")
	}

	var obj3, obj4 internal.Data
	obj3.X = []byte("hello")
	obj4.N = 100
	MustWrapObject(obj3).MustUnwrapObject(&obj4)
	if !bytes.Equal(obj4.X, obj3.X) {
		t.Fatal(`!bytes.Equal(obj4.X, obj3.X)`)
	}
	if obj4.N != 100 {
		t.Fatal(`obj4.N != 100`)
	}

	obj5, obj6 := obj1, obj1
	Nil.MustUnwrapObject(&obj6)
	if obj6 != obj5 {
		t.Fatal(`obj6 != obj5`)
	}
	liveZero.MustUnwrapObject(&obj6)
	if obj6 != obj5 {
		t.Fatal(`obj6 != obj5`)
	}

	if MustWrapObject(nil) != liveZero {
		t.Fatal(`MustWrapObject(nil) != liveZero`)
	}
	if data, err := wrapObjectImpl(nil); err != nil || data != liveZero {
		t.Fatal(`data, err := wrapObjectImpl(nil); err != nil || data != liveZero`)
	}
}

func TestMustWrapProtobufObject(t *testing.T) {
	if MustWrapProtobufObject(nil) != liveZero {
		t.Fatal(`MustWrapProtobufObject(nil) != liveZero`)
	}

	if MustWrapProtobufObject(&internal.Data{}).Bytes() != nil {
		t.Fatal(`MustWrapProtobufObject(&internal.Data{}).Bytes() != nil`)
	}

	var data internal.Data
	Data{}.MustUnwrapProtobufObject(&data)
	if data.X != nil || data.N != 0 {
		t.Fatal(`data.X != nil || data.N != 0`)
	}
	liveZero.MustUnwrapProtobufObject(&data)
	if data.X != nil || data.N != 0 {
		t.Fatal(`data.X != nil || data.N != 0`)
	}

	var d internal.Data
	d.X = []byte("hello")
	d.N = 100
	MustWrapProtobufObject(&d).MustUnwrapProtobufObject(&data)
	if string(data.X) != "hello" {
		t.Fatal(`string(data.X) != "hello"`)
	}
	if data.N != 100 {
		t.Fatal(`data.N != 100`)
	}
}
