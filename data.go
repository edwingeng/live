// Package live provides some types and functions designed for plugin hot reload,
// which solves the problem that the same type in different versions of a plugin
// are actually not the same at runtime.
package live

import (
	"encoding/binary"
	"encoding/json"
	"github.com/edwingeng/live/internal"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	"math"
)

var (
	Nil Data
)

// Data is a handy general-purpose data wrapper.
type Data struct {
	v interface{}
}

func (d Data) Bool() bool {
	return d.v.(*internal.Data).N == 1
}

func (d Data) Int() int {
	return int(d.v.(*internal.Data).N)
}

func (d Data) Int8() int8 {
	return int8(d.v.(*internal.Data).N)
}

func (d Data) Int16() int16 {
	return int16(d.v.(*internal.Data).N)
}

func (d Data) Int32() int32 {
	return int32(d.v.(*internal.Data).N)
}

func (d Data) Int64() int64 {
	return d.v.(*internal.Data).N
}

func (d Data) Uint() uint {
	return uint(d.v.(*internal.Data).N)
}

func (d Data) Uint8() uint8 {
	return uint8(d.v.(*internal.Data).N)
}

func (d Data) Uint16() uint16 {
	return uint16(d.v.(*internal.Data).N)
}

func (d Data) Uint32() uint32 {
	return uint32(d.v.(*internal.Data).N)
}

func (d Data) Uint64() uint64 {
	return uint64(d.v.(*internal.Data).N)
}

func (d Data) Float32() float32 {
	return math.Float32frombits(uint32(d.v.(*internal.Data).N))
}

func (d Data) Float64() float64 {
	return math.Float64frombits(uint64(d.v.(*internal.Data).N))
}

func (d Data) String() string {
	return string(d.v.(*internal.Data).X)
}

func (d Data) Bytes() []byte {
	return d.v.(*internal.Data).X
}

func (d Data) Complex64() complex64 {
	x := d.v.(*internal.Data).X
	r := binary.LittleEndian.Uint32(x[:4])
	i := binary.LittleEndian.Uint32(x[4:])
	return complex(math.Float32frombits(r), math.Float32frombits(i))
}

func (d Data) Complex128() complex128 {
	x := d.v.(*internal.Data).X
	r := binary.LittleEndian.Uint64(x[:8])
	i := binary.LittleEndian.Uint64(x[8:])
	return complex(math.Float64frombits(r), math.Float64frombits(i))
}

func (d Data) UnwrapObject(out interface{}) error {
	if d.v == nil {
		return nil
	}
	if len(d.v.(*internal.Data).X) == 0 {
		return nil
	}
	x, ok := out.(interface {
		UnmarshalJSON([]byte) error
	})
	if ok {
		return x.UnmarshalJSON(d.v.(*internal.Data).X)
	} else {
		return json.Unmarshal(d.v.(*internal.Data).X, out)
	}
}

func (d Data) MustUnwrapObject(out interface{}) {
	err := d.UnwrapObject(out)
	if err != nil {
		panic(err)
	}
}

type ProtobufUnmarshaler interface {
	Unmarshal([]byte) error
}

func (d Data) UnwrapProtobufObject(out ProtobufUnmarshaler) error {
	if d.v == nil {
		return nil
	}
	if len(d.v.(*internal.Data).X) == 0 {
		return nil
	}
	return out.Unmarshal(d.v.(*internal.Data).X)
}

func (d Data) MustUnwrapProtobufObject(out ProtobufUnmarshaler) {
	err := d.UnwrapProtobufObject(out)
	if err != nil {
		panic(err)
	}
}

func (d Data) Value() interface{} {
	return d.v
}

func (d Data) Marshallable() bool {
	_, ok := d.v.(*internal.Data)
	return ok
}

func (d Data) MarshalJSON() ([]byte, error) {
	x, ok := d.v.(*internal.Data)
	if ok {
		return x.MarshalJSON()
	}
	return internal.Data{}.MarshalJSON()
}

func (d Data) MarshalEasyJSON(w *jwriter.Writer) {
	x, ok := d.v.(*internal.Data)
	if ok {
		x.MarshalEasyJSON(w)
		return
	}
	internal.Data{}.MarshalEasyJSON(w)
}

func (d *Data) UnmarshalJSON(dAtA []byte) error {
	var v internal.Data
	err := v.UnmarshalJSON(dAtA)
	if err != nil {
		return err
	}
	d.v = &v
	return nil
}

func (d *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	var v internal.Data
	v.UnmarshalEasyJSON(l)
	d.v = &v
}

func (d Data) TurnIntoHermit() Hermit {
	x, ok := d.v.(*internal.Data)
	if ok {
		return Hermit{*x}
	}
	return Hermit{}
}
