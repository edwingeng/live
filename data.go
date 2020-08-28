package live

import (
	"encoding/binary"
	"encoding/json"
	"math"

	"github.com/edwingeng/live/internal"
)

var (
	Nil Data
)

type Data struct {
	v interface{}
}

func (d Data) ToBool() bool {
	return d.v.(*internal.Data).N == 1
}

func (d Data) ToInt() int {
	return int(d.v.(*internal.Data).N)
}

func (d Data) ToInt8() int8 {
	return int8(d.v.(*internal.Data).N)
}

func (d Data) ToInt16() int16 {
	return int16(d.v.(*internal.Data).N)
}

func (d Data) ToInt32() int32 {
	return int32(d.v.(*internal.Data).N)
}

func (d Data) ToInt64() int64 {
	return d.v.(*internal.Data).N
}

func (d Data) ToUint() uint {
	v, _ := binary.Uvarint(d.v.(*internal.Data).X)
	return uint(v)
}

func (d Data) ToUint8() uint8 {
	return uint8(d.v.(*internal.Data).N)
}

func (d Data) ToUint16() uint16 {
	return uint16(d.v.(*internal.Data).N)
}

func (d Data) ToUint32() uint32 {
	return uint32(d.v.(*internal.Data).N)
}

func (d Data) ToUint64() uint64 {
	v, _ := binary.Uvarint(d.v.(*internal.Data).X)
	return v
}

func (d Data) ToFloat32() float32 {
	return math.Float32frombits(uint32(d.v.(*internal.Data).N))
}

func (d Data) ToFloat64() float64 {
	v, _ := binary.Uvarint(d.v.(*internal.Data).X)
	return math.Float64frombits(v)
}

func (d Data) ToString() string {
	return string(d.v.(*internal.Data).X)
}

func (d Data) ToBytes() []byte {
	if d.v != nil {
		return d.v.(*internal.Data).X
	} else {
		return nil
	}
}

func (d Data) V() interface{} {
	return d.v
}

func (d Data) ToProtobufObj(obj interface {
	Unmarshal([]byte) error
}) {
	if len(d.v.(*internal.Data).X) == 0 {
		return
	}
	err := obj.Unmarshal(d.v.(*internal.Data).X)
	if err != nil {
		panic(err)
	}
}

func (d Data) ToJSONObj(obj interface{}) {
	if d.v == nil {
		return
	}
	if len(d.v.(*internal.Data).X) == 0 {
		return
	}
	x, ok := obj.(interface {
		UnmarshalJSON([]byte) error
	})
	if ok {
		err := x.UnmarshalJSON(d.v.(*internal.Data).X)
		if err != nil {
			panic(err)
		}
	} else {
		err := json.Unmarshal(d.v.(*internal.Data).X, obj)
		if err != nil {
			panic(err)
		}
	}
}

func (d Data) Persistent() (Persistent, bool) {
	x, ok := d.v.(*internal.Data)
	if ok {
		return Persistent{
			d: *x,
		}, true
	}

	return Persistent{}, false
}
