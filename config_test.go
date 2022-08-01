package live

import (
	"bytes"
	"encoding/json"
	"github.com/edwingeng/live/internal"
	"math"
	"testing"
	"unsafe"
)

func TestBlacklist(t *testing.T) {
	func() {
		defer rec()
		cfg := NewConfig([]string{"github.com/edwingeng/live/internal"})
		var v internal.Data
		cfg.WrapValueDirect(v)
		t.Fatal("WrapValueDirect() should panic")
	}()

	func() {
		defer rec()
		cfg := NewConfig([]string{"github.com/edwingeng/live"})
		var v internal.Data
		cfg.WrapValueDirect(v)
		t.Fatal("WrapValueDirect() should panic")
	}()

	func() {
		type foo struct {
			U internal.Data
		}
		type bar struct {
			V *foo
		}

		defer rec()
		cfg := NewConfig([]string{"github.com/edwingeng/live"})
		WrapValueDirect(&bar{}, cfg)
		t.Fatal("WrapValueDirect() should panic")
	}()

	cfg := NewConfig([]string{"github.com/edwin"})
	var v internal.Data
	cfg.WrapValueDirect(v)

	empty := NewConfig([]string{""})
	if empty.Blacklist != nil {
		t.Fatal(`empty.Blacklist != nil`)
	}
}

func TestBlacklist_cover(t *testing.T) {
	var bl blacklist = []string{
		"a/b/c",
		"a/x",
	}

	data := map[string]bool{
		"a/b/c":     true,
		"a/b":       false,
		"a/b/c/d":   true,
		"a/b/c/d/e": true,
		"a/x":       true,
		"a":         false,
		"a/x/y":     true,
		"a/x/y/z":   true,
		"h":         false,
		"h/i":       false,
	}

	for pkg, expected := range data {
		if bl.cover(pkg) != expected {
			t.Fatalf("bl.cover(pkg) != expected. pkg: %s, expected: %v", pkg, expected)
		}
	}
}

func TestConfig_WrapPrimitives(t *testing.T) {
	cfg := NewConfig(nil)
	if cfg.WrapValueDirect(1).Value() != 1 {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(int8(1)).Value() != int8(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(int16(1)).Value() != int16(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(int32(1)).Value() != int32(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(int64(1)).Value() != int64(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(uint(1)).Value() != uint(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(uint8(1)).Value() != uint8(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(uint16(1)).Value() != uint16(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(uint32(1)).Value() != uint32(1) {
		t.Fatal("impossible")
	}
	if cfg.WrapValueDirect(uint64(1)).Value() != uint64(1) {
		t.Fatal("impossible")
	}

	if WrapValueDirect(true, cfg).Value() != true {
		t.Fatal("impossible")
	}
	if WrapValueDirect(float32(0.01), cfg).Value() != float32(0.01) {
		t.Fatal("impossible")
	}
	if WrapValueDirect(0.01, cfg).Value() != 0.01 {
		t.Fatal("impossible")
	}
	if WrapValueDirect("hello", cfg).Value() != "hello" {
		t.Fatal("impossible")
	}
	if !bytes.Equal(WrapValueDirect([]byte("hello"), cfg).Value().([]byte), []byte("hello")) {
		t.Fatal("impossible")
	}
}

func TestConfig_WrapUintptr(t *testing.T) {
	func() {
		defer rec()
		cfg := NewConfig(nil)
		cfg.WrapValueDirect(uintptr(0))
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapComplex64(t *testing.T) {
	cfg := NewConfig(nil)
	a := []complex64{
		complex(0, 0.12345),
		complex(1, 1.54321),
	}
	for _, v := range a {
		d := cfg.WrapValueDirect(v)
		switch u := d.Value().(type) {
		case complex64:
			if real(u) != real(v) {
				t.Fatal(`real(u) != real(v)`)
			}
			if imag(u) != imag(v) {
				t.Fatal(`imag(u) != imag(v)`)
			}
		default:
			t.Fatal("unexpected data type")
		}
	}
}

func TestConfig_WrapComplex128(t *testing.T) {
	cfg := NewConfig(nil)
	a := []complex128{
		complex(0, 0.12345),
		complex(1, 1.54321),
	}
	for _, v := range a {
		d := cfg.WrapValueDirect(v)
		switch u := d.Value().(type) {
		case complex128:
			if real(u) != real(v) {
				t.Fatal(`real(u) != real(v)`)
			}
			if imag(u) != imag(v) {
				t.Fatal(`imag(u) != imag(v)`)
			}
		default:
			t.Fatal("unexpected data type")
		}
	}
}

func TestConfig_WrapArray(t *testing.T) {
	cfg := NewConfig(nil)
	var n1, n2, n3 uint64 = 0, 1, math.MaxUint64
	a := [3]*uint64{&n1, &n2, &n3}
	d := cfg.WrapValueDirect(a)
	switch u := d.Value().(type) {
	case [3]*uint64:
		if u != a {
			t.Fatal(`u != a`)
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer rec()
		vals := [3]func(){
			func() {},
			func() {},
			func() {},
		}
		WrapValueDirect(vals, cfg)
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapChan(t *testing.T) {
	cfg := NewConfig(nil)
	a := []uint64{0, 1, math.MaxUint64}
	type wrapper struct {
		n uint64
	}
	ch := make(chan wrapper, len(a))
	for _, v := range a {
		ch <- wrapper{n: v}
	}

	d := cfg.WrapValueDirect(ch)
	switch u := d.Value().(type) {
	case chan wrapper:
		for _, v := range a {
			x := <-u
			if x.n != v {
				t.Fatal("x.n != v")
			}
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer rec()
		cfg.WrapValueDirect(make(chan []uintptr))
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapFunc(t *testing.T) {
	defer rec()
	cfg := NewConfig(nil)
	cfg.WrapValueDirect(func() {})
	t.Fatal("WrapValueDirect() should panic")
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

func TestConfig_WrapInterface(t *testing.T) {
	cfg := NewConfig(nil)
	var v interface{} = 100
	cfg.WrapValueDirect(v)

	if cfg.WrapValueDirect(nil).Value() != nil {
		t.Fatal("WrapValueDirect(nil).Value() != nil")
	}

	if cfg.WrapValueDirect([]byte(nil)).Value() == nil {
		t.Fatal("WrapValueDirect([]byte(nil)).Value() == nil")
	}

	var w1 myPrinterWrapper1
	w1.p = myPrinter{}
	func() {
		defer rec()
		cfg.WrapValueDirect(&w1)
		t.Fatal("WrapValueDirect() should panic")
	}()

	var w2 myPrinterWrapper2
	w2.p = myPrinter{}
	if cfg.WrapValueDirect(&w2).Value() != &w2 {
		t.Fatal("cfg.WrapValueDirect(&w2).Value() != &w2")
	}
}

func TestConfig_WrapMap(t *testing.T) {
	cfg := NewConfig(nil)
	m := map[int]string{
		1: "10",
		2: "20",
		3: "30",
	}
	d := cfg.WrapValueDirect(m)
	switch u := d.Value().(type) {
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
		defer rec()
		type wrapper struct {
			n uintptr
		}
		xMap := make(map[wrapper][]interface{})
		cfg.WrapValueDirect(xMap)
		t.Fatal("WrapValueDirect() should panic")
	}()

	func() {
		defer rec()
		xMap := make(map[int][]interface{})
		cfg.WrapValueDirect(xMap)
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapPointer(t *testing.T) {
	cfg := NewConfig(nil)
	n := 100
	p := &n
	d := WrapValueDirect(&p, cfg)
	switch u := d.Value().(type) {
	case **int:
		if u != &p {
			t.Fatal("u != &p")
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer rec()
		var v interface{}
		cfg.WrapValueDirect(&v)
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapSlice(t *testing.T) {
	cfg := NewConfig(nil)
	var n1, n2, n3 uint64 = 0, 1, math.MaxUint64
	a := []*uint64{&n1, &n2, &n3}
	d := cfg.WrapValueDirect(a)
	switch u := d.Value().(type) {
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
		defer rec()
		vals := []func(){
			func() {},
			func() {},
			func() {},
		}
		WrapValueDirect(vals, cfg)
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_WrapUnsafePointer(t *testing.T) {
	func() {
		defer rec()
		cfg := NewConfig(nil)
		ptr := unsafe.Pointer(&cfg)
		cfg.WrapValueDirect(ptr)
		t.Fatal("WrapValueDirect() should panic")
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

func (o *omega) setA(a int) {
	o.A = a
}

func TestConfig_WrapStruct(t *testing.T) {
	cfg := NewConfig(nil)
	v := omega{
		A: 100,
		B: "hello",
		C: struct {
			D int8
			E bool
		}{D: 1, E: true},
	}
	v.setA(200)

	d := cfg.WrapValueDirect(v)
	switch u := d.Value().(type) {
	case omega:
		if u != v {
			t.Fatal("u != v")
		}
	default:
		t.Fatal("unexpected data type")
	}

	func() {
		defer rec()
		q := struct {
			N int
			X interface{}
		}{}
		cfg.WrapValueDirect(q)
		t.Fatal("WrapValueDirect() should panic")
	}()
}

func TestConfig_EmbeddedLiveData(t *testing.T) {
	type foo struct {
		X Data
	}

	var f1 foo
	f1.X = WrapInt(100)
	cfg := NewConfig(nil)
	cfg.WrapValueDirect(&f1)

	data, err := json.Marshal(&f1)
	if err != nil {
		t.Fatal(err)
	}
	var f2 foo
	if err := json.Unmarshal(data, &f2); err != nil {
		t.Fatal(err)
	}
	if f2.X.Int() != 100 {
		t.Fatal(`f2.X.Int() != 100`)
	}
}

func TestConfig_FieldWithLiveTag(t *testing.T) {
	func() {
		defer rec()
		cfg := NewConfig(nil)
		cfg.WrapValueDirect(&struct {
			X func() `live:"false"`
		}{})
		t.Fatal(`WrapValueDirect() should panic`)
	}()

	func() {
		defer rec()
		cfg := NewConfig(nil)
		cfg.WrapValueDirect(&struct {
			X func() `live:"0"`
		}{})
		t.Fatal(`WrapValueDirect() should panic`)
	}()

	cfg := NewConfig(nil)
	cfg.WrapValueDirect(&struct {
		X func() `live:"true"`
	}{})
	cfg.WrapValueDirect(&struct {
		X func() `live:"1"`
	}{})
}

func TestConfig_SkipTypeCheck(t *testing.T) {
	cfg := NewConfig(nil)
	cfg.SkipTypeCheck = true
	cfg.WrapValueDirect(&struct {
		X func()
	}{})
}

type node1 struct {
	X *node1
}

type node2 struct {
	X *node3
}

type node3 struct {
	X *node2
}

func TestConfig_CyclicTypeReference(t *testing.T) {
	cfg := NewConfig(nil)
	var n1 node1
	cfg.WrapValueDirect(&n1)
	var n2 node2
	cfg.WrapValueDirect(&n2)
}
