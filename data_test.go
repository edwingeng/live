package live

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/edwingeng/live/internal"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"reflect"
	"testing"
)

type (
	marshalFunc   func(d Data) ([]byte, error)
	unmarshalFunc func(dAtA []byte) (Data, error)
)

func marshalIndirect(t *testing.T, marshal marshalFunc, unmarshal unmarshalFunc) {
	a1 := []Data{Nil, {100}, {"str"}}
	for _, d := range a1 {
		bts, err := marshal(d)
		if err != nil {
			t.Fatal(err)
		}
		if string(bts) != "{}" {
			t.Fatal(`string(bts) != "{}"`)
		}
	}

	a2 := []Data{
		{&internal.Data{X: []byte("str")}},
		{&internal.Data{N: 100}},
	}
	r2 := []string{
		`{"X":"` + base64.StdEncoding.EncodeToString([]byte("str")) + `"}`,
		`{"N":100}`,
	}
	for i, d := range a2 {
		bts, err := marshal(d)
		if err != nil {
			t.Fatal(err)
		}
		if string(bts) != r2[i] {
			t.Fatal(`string(bts) != r2[i]`)
		}
	}

	mustEqual := func(d1, d2 Data, t *testing.T) {
		t.Helper()
		if reflect.TypeOf(d1) != reflect.TypeOf(d2) {
			t.Fatal(`reflect.TypeOf(d1) != reflect.TypeOf(d2)`)
		}
		if reflect.TypeOf(d1) != liveDataType {
			if d1 != d2 {
				t.Fatal(`d1 != d2`)
			}
		}
		v1 := d1.v.(*internal.Data)
		v2 := d2.v.(*internal.Data)
		if (v1.N != 0 || v2.N != 0) && v1.N != v2.N {
			t.Fatal(`v1.N != v2.N`)
		}
		if !bytes.Equal(v1.X, v2.X) {
			t.Fatal(`!bytes.Equal(v1.X, v2.X)`)
		}
	}

	for i, str := range r2 {
		d, err := unmarshal([]byte(str))
		if err != nil {
			t.Fatal(err)
		}
		mustEqual(d, a2[i], t)
	}
}

func TestData_MarshalJSON(t *testing.T) {
	marshal := func(d Data) ([]byte, error) {
		return d.MarshalJSON()
	}
	unmarshal := func(dAtA []byte) (Data, error) {
		var d Data
		err := d.UnmarshalJSON(dAtA)
		return d, err
	}
	marshalIndirect(t, marshal, unmarshal)

	type foo struct {
		C int
		D Data
	}

	var f1, f2 foo
	f1.C = 100
	f1.D = WrapInt64(200)
	data, err := json.Marshal(&f1)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &f2); err != nil {
		t.Fatal(err)
	}
	if f2.C != 100 || f2.D.Int64() != 200 && f2.D.Bytes() != nil {
		t.Fatal(`f2.C != 100 || f2.D.Int64() != 200 && f2.D.Bytes() != nil`)
	}
}

func TestData_MarshalEasyJSON(t *testing.T) {
	marshal := func(d Data) ([]byte, error) {
		w := jwriter.Writer{}
		d.MarshalEasyJSON(&w)
		return w.Buffer.BuildBytes(), w.Error
	}
	unmarshal := func(dAtA []byte) (Data, error) {
		var d Data
		r := jlexer.Lexer{Data: dAtA}
		d.UnmarshalEasyJSON(&r)
		return d, r.Error()
	}
	marshalIndirect(t, marshal, unmarshal)
}

func TestData_TurnIntoHermit(t *testing.T) {
	if !reflect.DeepEqual(Nil.TurnIntoHermit(), Hermit{}) {
		t.Fatal(`!reflect.DeepEqual(Nil.TurnIntoHermit(), TurnIntoHermit{})`)
	}
	if WrapInt(1000).TurnIntoHermit().d.N != 1000 {
		t.Fatal(`WrapInt(1000).TurnIntoHermit().d.N != 1000`)
	}
	if !bytes.Equal(WrapString("hello").TurnIntoHermit().d.X, []byte("hello")) {
		t.Fatal(`!bytes.Equal(WrapString("hello").TurnIntoHermit().d.X, []byte("hello"))`)
	}
}

func TestData_Marshalable(t *testing.T) {
	if WrapInt(100).Marshalable() != true {
		t.Fatal(`WrapInt(100).Marshalable() != true`)
	}
	cfg := NewConfig(nil)
	if cfg.WrapValueDirect(100).Marshalable() != false {
		t.Fatal(`cfg.WrapValueDirect(100).Marshalable() != false`)
	}
}
