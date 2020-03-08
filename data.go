package live

import (
	"encoding/binary"
	"math"
)

func (d Data) ToBool() bool {
	return d.V == 1
}

func (d Data) ToInt() int {
	return int(d.V)
}

func (d Data) ToInt8() int8 {
	return int8(d.V)
}

func (d Data) ToInt16() int16 {
	return int16(d.V)
}

func (d Data) ToInt32() int32 {
	return int32(d.V)
}

func (d Data) ToInt64() int64 {
	return d.V
}

func (d Data) ToUint() uint {
	v, _ := binary.Uvarint(d.X)
	return uint(v)
}

func (d Data) ToUint8() uint8 {
	return uint8(d.V)
}

func (d Data) ToUint16() uint16 {
	return uint16(d.V)
}

func (d Data) ToUint32() uint32 {
	return uint32(d.V)
}

func (d Data) ToUint64() uint64 {
	v, _ := binary.Uvarint(d.X)
	return v
}

func (d Data) ToFloat32() float32 {
	return math.Float32frombits(uint32(d.V))
}

func (d Data) ToFloat64() float64 {
	v, _ := binary.Uvarint(d.X)
	return math.Float64frombits(v)
}

func (d Data) ToString() string {
	return string(d.X)
}

func (d Data) ToProtobuf(obj interface {
	Unmarshal([]byte) error
}) {
	if len(d.X) == 0 {
		return
	}
	err := obj.Unmarshal(d.X)
	if err != nil {
		panic(err)
	}
}

func (d Data) ToJSON(obj interface {
	UnmarshalJSON([]byte) error
}) {
	if len(d.X) == 0 {
		return
	}
	err := obj.UnmarshalJSON(d.X)
	if err != nil {
		panic(err)
	}
}
