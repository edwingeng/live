package live

import (
	"bytes"
	"testing"
)

func TestFromHermitBytes1(t *testing.T) {
	d1 := WrapInt64(100)
	buf, err := d1.TurnIntoHermit().Marshal()
	if err != nil {
		t.Fatal(err)
	}
	d2, err := FromHermitBytes(buf)
	if err != nil {
		t.Fatal(err)
	}
	if d2.Int() != 100 || d2.Bytes() != nil {
		t.Fatal(`d2.Int() != 100 || d2.Bytes() != nil`)
	}
}

func TestFromHermitBytes2(t *testing.T) {
	d1 := WrapString("hello")
	var buf [64]byte
	n, err := d1.TurnIntoHermit().MarshalTo(buf[:])
	if err != nil {
		t.Fatal(err)
	}
	d2, err := FromHermitBytes(buf[:n])
	if err != nil {
		t.Fatal(err)
	}
	if d2.Int() != 0 || !bytes.Equal(d2.Bytes(), []byte("hello")) {
		t.Fatal(`d2.Int() != 0 || !bytes.Equal(d2.Bytes(), []byte("hello"))`)
	}
}

func TestFromHermitBytes3(t *testing.T) {
	foo := struct {
		Str string
		Num int
	}{
		Str: "hello",
		Num: 100,
	}

	d1 := MustWrapObject(&foo)
	hermit := d1.TurnIntoHermit()
	buf := make([]byte, hermit.Size())
	n, err := hermit.MarshalToSizedBuffer(buf)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(buf) {
		t.Fatal(`n != len(buf)`)
	}

	d2, err := FromHermitBytes(buf[:n])
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

func TestFromHermitBytes4(t *testing.T) {
	d, err := FromHermitBytes([]byte("hello"))
	if err == nil {
		t.Fatal(`err == nil`)
	}
	if d != Nil {
		t.Fatal(`d != Nil`)
	}
}
