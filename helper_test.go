package live

import (
	"bytes"
	"math"
	"testing"

	"github.com/edwingeng/live/internal"
)

func TestBool(t *testing.T) {
	h := NewHelper(nil)
	a := []bool{true, false}
	for _, v := range a {
		d := h.WrapBool(v)
		if d.ToBool() != v {
			t.Fatal("d.ToBool() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToBool()
		t.Fatal("d.ToBool() should panic")
	}()
}

func TestInt(t *testing.T) {
	h := NewHelper(nil)
	a := []int{-1, 0, 1, 10, 1000}
	for _, v := range a {
		d := h.WrapInt(v)
		if d.ToInt() != v {
			t.Fatal("d.ToInt() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToInt()
		t.Fatal("d.ToInt() should panic")
	}()
}

func TestInt8(t *testing.T) {
	h := NewHelper(nil)
	a := []int8{math.MinInt8, -1, 0, 1, 10, math.MaxInt8}
	for _, v := range a {
		d := h.WrapInt8(v)
		if d.ToInt8() != v {
			t.Fatal("d.ToInt8() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToInt8()
		t.Fatal("d.ToInt8() should panic")
	}()
}

func TestInt16(t *testing.T) {
	h := NewHelper(nil)
	a := []int16{math.MinInt16, -1, 0, 1, 10, math.MaxInt16}
	for _, v := range a {
		d := h.WrapInt16(v)
		if d.ToInt16() != v {
			t.Fatal("d.ToInt16() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToInt16()
		t.Fatal("d.ToInt16() should panic")
	}()
}

func TestInt32(t *testing.T) {
	h := NewHelper(nil)
	a := []int32{math.MinInt32, -1, 0, 1, 10, math.MaxInt32}
	for _, v := range a {
		d := h.WrapInt32(v)
		if d.ToInt32() != v {
			t.Fatal("d.ToInt32() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToInt32()
		t.Fatal("d.ToInt32() should panic")
	}()
}

func TestInt64(t *testing.T) {
	h := NewHelper(nil)
	a := []int64{math.MinInt64, -1, 0, 1, 10, math.MaxInt64}
	for _, v := range a {
		d := h.WrapInt64(v)
		if d.ToInt64() != v {
			t.Fatal("d.ToInt64() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToInt64()
		t.Fatal("d.ToInt64() should panic")
	}()
}

func TestUint(t *testing.T) {
	h := NewHelper(nil)
	a := []uint{0, 1, 10, 1000}
	for _, v := range a {
		d := h.WrapUint(v)
		if d.ToUint() != v {
			t.Fatal("d.ToUint() != v")
		}
		if d.v.(*internal.Data).X == nil && d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToUint()
		t.Fatal("d.ToUint() should panic")
	}()
}

func TestUint8(t *testing.T) {
	h := NewHelper(nil)
	a := []uint8{0, 1, 10, math.MaxUint8}
	for _, v := range a {
		d := h.WrapUint8(v)
		if d.ToUint8() != v {
			t.Fatal("d.ToUint8() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToUint8()
		t.Fatal("d.ToUint8() should panic")
	}()
}

func TestUint16(t *testing.T) {
	h := NewHelper(nil)
	a := []uint16{0, 1, 10, math.MaxUint16}
	for _, v := range a {
		d := h.WrapUint16(v)
		if d.ToUint16() != v {
			t.Fatal("d.ToUint16() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToUint16()
		t.Fatal("d.ToUint16() should panic")
	}()
}

func TestUint32(t *testing.T) {
	h := NewHelper(nil)
	a := []uint32{0, 1, 10, math.MaxUint32}
	for _, v := range a {
		d := h.WrapUint32(v)
		if d.ToUint32() != v {
			t.Fatal("d.ToUint32() != v")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToUint32()
		t.Fatal("d.ToUint32() should panic")
	}()
}

func TestUint64(t *testing.T) {
	h := NewHelper(nil)
	a := []uint64{0, 1, 10, math.MaxUint64}
	for _, v := range a {
		d := h.WrapUint64(v)
		if d.ToUint64() != v {
			t.Fatal("d.ToUint64() != v")
		}
		if d.v.(*internal.Data).X == nil && d.v.(*internal.Data).N != 0 {
			t.Fatal("d.v.(*internal.Data).X == nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToUint64()
		t.Fatal("d.ToUint64() should panic")
	}()
}

func TestUintptr(t *testing.T) {
	h := NewHelper(nil)
	a := []uintptr{0, 1, 10, 1000}
	for _, v := range a {
		d := h.WrapValue(v)
		switch u := d.V().(type) {
		case uintptr:
			if u != v {
				t.Fatal("u != v")
			}
		default:
			t.Fatal("unexpected data type")
		}
	}
}

func TestFloat32(t *testing.T) {
	h := NewHelper(nil)
	a := []float32{-9.9, -1, 0, 1, 10, 1000.28}
	for _, v := range a {
		d := h.WrapFloat32(v)
		if math.Abs(float64(d.ToFloat32()-v)) > 0.000001 {
			t.Fatal("math.Abs(float64(d.ToFloat32()-v)) > 0.000001")
		}
		if d.v.(*internal.Data).X != nil {
			t.Fatal("d.v.(*internal.Data).X != nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToFloat32()
		t.Fatal("d.ToFloat32() should panic")
	}()
}

func TestFloat64(t *testing.T) {
	h := NewHelper(nil)
	a := []float64{-9.9, -1, 0, 1, 10, 1000.28}
	for _, v := range a {
		d := h.WrapFloat64(v)
		if math.Abs(d.ToFloat64()-v) > 0.000001 {
			t.Fatal("math.Abs(d.ToFloat64()-v) > 0.000001")
		}
		if d.v.(*internal.Data).X == nil {
			t.Fatal("d.v.(*internal.Data).X == nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToFloat64()
		t.Fatal("d.ToFloat64() should panic")
	}()
}

func TestString(t *testing.T) {
	h := NewHelper(nil)
	a := []string{"", "hello", "it is a good day to die"}
	for _, v := range a {
		d := h.WrapString(v)
		if d.ToString() != v {
			t.Fatal("d.ToString() != v")
		}
		if d.v.(*internal.Data).X == nil {
			t.Fatal("d.v.(*internal.Data).X == nil")
		}
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var d Data
		d.ToString()
		t.Fatal("d.ToString() should panic")
	}()
}

func TestBytes(t *testing.T) {
	h := NewHelper(nil)
	a := [][]byte{[]byte(""), []byte("hello"), []byte("it is a good day to die")}
	for _, v := range a {
		d := h.WrapBytes(v)
		if !bytes.Equal(d.ToBytes(), v) {
			t.Fatal("!bytes.Equal(d.ToBytes(), v)")
		}
		if d.v.(*internal.Data).X == nil {
			t.Fatal("d.v.(*internal.Data).X == nil")
		}
	}

	if h.WrapBytes(nil).ToBytes() != nil {
		t.Fatal("h.WrapBytes(nil).ToBytes() != nil")
	}
	if Nil.ToBytes() != nil {
		t.Fatal("Nil.ToBytes() != nil")
	}
}

func TestComplex64(t *testing.T) {
	h := NewHelper(nil)
	a := []complex64{
		complex(0, 0.5),
		complex(1, 1.5),
	}
	for _, v := range a {
		d := h.WrapValue(v)
		switch u := d.V().(type) {
		case complex64:
			if math.Abs(float64(real(u)-real(v))) > 0.000001 {
				t.Fatal("math.Abs(float64(real(u)-real(v))) > 0.000001")
			}
			if math.Abs(float64(imag(u)-imag(v))) > 0.000001 {
				t.Fatal("math.Abs(float64(imag(u)-imag(v))) > 0.000001")
			}
		default:
			t.Fatal("unexpected data type")
		}
	}
}

func TestComplex128(t *testing.T) {
	h := NewHelper(nil)
	a := []complex128{
		complex(0, 0.5),
		complex(1, 1.5),
	}
	for _, v := range a {
		d := h.WrapValue(v)
		switch u := d.V().(type) {
		case complex128:
			if math.Abs(real(u)-real(v)) > 0.000001 {
				t.Fatal("math.Abs(real(u)-real(v)) > 0.000001")
			}
			if math.Abs(imag(u)-imag(v)) > 0.000001 {
				t.Fatal("math.Abs(imag(u)-imag(v)) > 0.000001")
			}
		default:
			t.Fatal("unexpected data type")
		}
	}
}

func TestArray(t *testing.T) {
	h := NewHelper(nil)
	var n1, n2, n3 uint64 = 0, 1, math.MaxUint64
	a := [3]*uint64{&n1, &n2, &n3}
	d := h.WrapValue(a)
	switch u := d.V().(type) {
	case [3]*uint64:
		for i, v := range u {
			if v != a[i] {
				t.Fatal("v != a[i]")
			}
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer func() {
			_ = recover()
		}()
		a := []func(){
			func() {},
			func() {},
			func() {},
		}
		h.WrapValue(a)
		t.Fatal("h.WrapValue() should panic")
	}()
}

func TestChan(t *testing.T) {
	h := NewHelper(nil)
	a := []uint64{0, 1, math.MaxUint64}
	type w struct {
		n uint64
	}
	ch := make(chan w, len(a))
	for _, v := range a {
		ch <- w{n: v}
	}

	d := h.WrapValue(ch)
	switch u := d.V().(type) {
	case chan w:
		for _, v := range a {
			x := <-u
			if x.n != v {
				t.Fatal("x.n != v")
			}
		}
	default:
		t.Fatal("unexpected data type")
	}
}

func TestFunc(t *testing.T) {
	h := NewHelper(nil)
	func() {
		defer func() {
			_ = recover()
		}()
		h.WrapValue(func() {})
		t.Fatal("h.WrapValue() should panic")
	}()
}

func TestEmbedded(t *testing.T) {
	h := NewHelper(nil)
	d := h.WrapValue(h.WrapInt(1))
	if d.V().(Data).ToInt() != 1 {
		t.Fatal("d.V().(Data).ToInt() != 1")
	}

	type Super struct {
		d Data
	}
	var super Super
	h.WrapValue(super)
}

type printer interface {
	Print()
}

type myPrinter struct{}

func (_ myPrinter) Print() {}

type myPrinterWrapper1 struct {
	p printer
}

type myPrinterWrapper2 struct {
	p printer `live:"true"`
}

func TestInterface(t *testing.T) {
	h := NewHelper(nil)
	var v interface{} = 100
	h.WrapValue(v)

	if h.WrapValue(nil).V() != nil {
		t.Fatal("h.WrapValue(nil).V() != nil")
	}

	if h.WrapValue([]byte(nil)).V() == nil {
		t.Fatal("h.WrapValue([]byte(nil)).V() == nil")
	}

	var w1 myPrinterWrapper1
	w1.p = myPrinter{}
	func() {
		defer func() {
			_ = recover()
		}()
		h.WrapValue(&w1)
		t.Fatal("WrapValue should panic")
	}()

	var w2 myPrinterWrapper2
	w2.p = myPrinter{}
	if h.WrapValue(&w2).V() != &w2 {
		t.Fatal("h.WrapValue(&w2).V() != &w2")
	}
}

func TestMap(t *testing.T) {
	h := NewHelper(nil)
	m := map[int]string{
		1: "10",
		2: "20",
		3: "30",
	}
	d := h.WrapValue(m)
	switch u := d.V().(type) {
	case map[int]string:
		if len(u) != len(m) {
			t.Fatal("len(u) != len(m)")
		}
		for k := range u {
			if _, ok := m[k]; !ok {
				t.Fatal("_, ok := m[k]; !ok")
			}
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer func() {
			_ = recover()
		}()
		m := make(map[int][]interface{})
		h.WrapValue(m)
		t.Fatal("h.WrapValue() should panic")
	}()
}

func TestPointer(t *testing.T) {
	h := NewHelper(nil)
	n := 100
	p := &n
	d := h.WrapValue(&p)
	switch u := d.V().(type) {
	case **int:
		if u != &p {
			t.Fatal("u != &p")
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer func() {
			_ = recover()
		}()
		var v interface{}
		h.WrapValue(&v)
		t.Fatal("h.WrapValue() should panic")
	}()
}

func TestSlice(t *testing.T) {
	h := NewHelper(nil)
	var n1, n2, n3 uint64 = 0, 1, math.MaxUint64
	a := []*uint64{&n1, &n2, &n3}
	d := h.WrapValue(a)
	switch u := d.V().(type) {
	case []*uint64:
		for i, v := range u {
			if v != a[i] {
				t.Fatal("v != a[i]")
			}
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer func() {
			_ = recover()
		}()
		a := []func(){
			func() {},
			func() {},
			func() {},
		}
		h.WrapValue(a)
		t.Fatal("h.WrapValue() should panic")
	}()
}

type omega struct {
	A int
	B string
	C struct {
		D int8
		E bool
	}
}

func (o *omega) SetA(a int) {
	o.A = a
}

func TestStruct(t *testing.T) {
	h := NewHelper(nil)
	v := omega{
		A: 100,
		B: "hello",
		C: struct {
			D int8
			E bool
		}{D: 1, E: true},
	}
	v.SetA(200)

	d := h.WrapValue(v)
	switch u := d.V().(type) {
	case omega:
		if u != v {
			t.Fatal("u != v")
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer func() {
			_ = recover()
		}()
		v := struct {
			N int
			X interface{}
		}{}
		h.WrapValue(v)
		t.Fatal("h.WrapValue() should panic")
	}()
}

func TestBlacklist(t *testing.T) {
	func() {
		defer func() {
			_ = recover()
		}()
		h := NewHelper([]string{"github.com/edwingeng/live/internal"})
		var v internal.Data
		h.WrapValue(v)
		t.Fatal("h.WrapValue() should panic")
	}()

	func() {
		defer func() {
			_ = recover()
		}()
		h := NewHelper([]string{"github.com/edwingeng/live"})
		var v internal.Data
		h.WrapValue(v)
		t.Fatal("h.WrapValue() should panic")
	}()

	func() {
		h := NewHelper([]string{"github.com/edwin"})
		var v internal.Data
		h.WrapValue(v)
	}()
}

func TestJSON(t *testing.T) {
	var obj1, obj2 struct {
		A int64
		B string
	}
	obj1.A = 100
	obj2.B = "hello"

	h := NewHelper(nil)
	h.WrapJSONObj(&obj1).ToJSONObj(&obj2)
	if obj2 != obj1 {
		t.Fatal("obj2 != obj1")
	}
}

func TestHelper_FromInternalBytes(t *testing.T) {
	h := NewHelper(nil)
	d := h.WrapUint64(0xFFFFFFFFFFFFFFFF)
	internalBytes, ok := d.Persistent().PeekInternalBytes()
	if !ok {
		t.Fatal("PeekInternalBytes does not work")
	}

	newData := h.FromInternalBytes(internalBytes)
	if newData.ToUint64() != 0xFFFFFFFFFFFFFFFF {
		t.Fatal("newData.ToUint64() != 0xFFFFFFFFFFFFFFFF")
	}
}
