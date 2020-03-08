package live

import (
	"encoding/binary"
	"math"
	"reflect"

	"github.com/edwingeng/live/internal"
)

type Data struct {
	d internal.Data
	v interface{}
}

func (d Data) ToBool() bool {
	return d.d.N == 1
}

func (d Data) ToInt() int {
	return int(d.d.N)
}

func (d Data) ToInt8() int8 {
	return int8(d.d.N)
}

func (d Data) ToInt16() int16 {
	return int16(d.d.N)
}

func (d Data) ToInt32() int32 {
	return int32(d.d.N)
}

func (d Data) ToInt64() int64 {
	return d.d.N
}

func (d Data) ToUint() uint {
	v, _ := binary.Uvarint(d.d.X)
	return uint(v)
}

func (d Data) ToUint8() uint8 {
	return uint8(d.d.N)
}

func (d Data) ToUint16() uint16 {
	return uint16(d.d.N)
}

func (d Data) ToUint32() uint32 {
	return uint32(d.d.N)
}

func (d Data) ToUint64() uint64 {
	v, _ := binary.Uvarint(d.d.X)
	return v
}

func (d Data) ToFloat32() float32 {
	return math.Float32frombits(uint32(d.d.N))
}

func (d Data) ToFloat64() float64 {
	v, _ := binary.Uvarint(d.d.X)
	return math.Float64frombits(v)
}

func (d Data) ToString() string {
	return string(d.d.X)
}

func (d Data) V() interface{} {
	return d.v
}

func (d Data) ToProtobufObj(obj interface {
	Unmarshal([]byte) error
}) {
	if len(d.d.X) == 0 {
		return
	}
	err := obj.Unmarshal(d.d.X)
	if err != nil {
		panic(err)
	}
}

func (d Data) ToJSONObj(obj interface {
	UnmarshalJSON([]byte) error
}) {
	if len(d.d.X) == 0 {
		return
	}
	err := obj.UnmarshalJSON(d.d.X)
	if err != nil {
		panic(err)
	}
}

func (d Data) Marshal() (dAtA []byte, err error) {
	if d.v != nil {
		panic("Marshal does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.Marshal()
}

func (d Data) MarshalTo(dAtA []byte) (int, error) {
	if d.v != nil {
		panic("MarshalTo does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.MarshalTo(dAtA)
}

func (d Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if d.v != nil {
		panic("MarshalToSizedBuffer does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.MarshalToSizedBuffer(dAtA)
}

func (d Data) Size() (n int) {
	if d.v != nil {
		panic("Size does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.Size()
}

func (d *Data) Unmarshal(dAtA []byte) error {
	if d.v != nil {
		panic("Unmarshal does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.Unmarshal(dAtA)
}

func (d Data) MarshalJSON() ([]byte, error) {
	if d.v != nil {
		panic("MarshalJSON does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.MarshalJSON()
}

func (d *Data) UnmarshalJSON(data []byte) error {
	if d.v != nil {
		panic("UnmarshalJSON does not support type " + reflect.TypeOf(d.v).Name())
	}
	return d.d.UnmarshalJSON(data)
}
