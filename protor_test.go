package live

import (
	"bytes"
	"reflect"
	"testing"
)

func TestData_NewProtor(t *testing.T) {
	if !reflect.DeepEqual(NewProtor(Nil), Protor{}) {
		t.Fatal(`!reflect.DeepEqual(NewProtor(Nil), Protor{})`)
	}
	if NewProtor(WrapInt(1000)).d.N != 1000 {
		t.Fatal(`NewProtor(WrapInt(1000)).d.N != 1000`)
	}
	if !bytes.Equal(NewProtor(WrapString("hello")).d.X, []byte("hello")) {
		t.Fatal(`!bytes.Equal(NewProtor(WrapString("hello")).d.X, []byte("hello"))`)
	}
}

func TestFromProtorBytes1(t *testing.T) {
	d1 := WrapInt64(100)
	buf, err := NewProtor(d1).Marshal()
	if err != nil {
		t.Fatal(err)
	}
	d2, err := FromProtorBytes(buf)
	if err != nil {
		t.Fatal(err)
	}
	if d2.Int() != 100 || d2.Bytes() != nil {
		t.Fatal(`d2.Int() != 100 || d2.Bytes() != nil`)
	}
}

func TestFromProtorBytes2(t *testing.T) {
	d1 := WrapString("hello")
	var buf [64]byte
	n, err := NewProtor(d1).MarshalTo(buf[:])
	if err != nil {
		t.Fatal(err)
	}
	d2, err := FromProtorBytes(buf[:n])
	if err != nil {
		t.Fatal(err)
	}
	if d2.Int() != 0 || !bytes.Equal(d2.Bytes(), []byte("hello")) {
		t.Fatal(`d2.Int() != 0 || !bytes.Equal(d2.Bytes(), []byte("hello"))`)
	}
}

func TestFromProtorBytes3(t *testing.T) {
	foo := struct {
		Str string
		Num int
	}{
		Str: "hello",
		Num: 100,
	}

	d1 := MustWrapObject(&foo)
	protor := NewProtor(d1)
	buf := make([]byte, protor.Size())
	n, err := protor.MarshalToSizedBuffer(buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(buf) {
		t.Fatal(`n != len(buf)`)
	}

	d2, err := FromProtorBytes(buf[:n])
	if err != nil {
		t.Fatal(err)
	}

	bar := foo
	bar.Str = ""
	bar.Num = 0
	d2.MustUnwrapObject(&bar)
	if bar.Str != foo.Str || bar.Num != foo.Num {
		t.Fatal(`bar.Str != foo.Str || bar.Num != foo.Num`)
	}
}

func TestFromProtorBytes4(t *testing.T) {
	d, err := FromProtorBytes([]byte("hello"))
	if err == nil {
		t.Fatal(`err == nil`)
	}
	if d != Nil {
		t.Fatal(`d != Nil`)
	}
}
