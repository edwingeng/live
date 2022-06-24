package live

import (
	"encoding/binary"
	"encoding/json"
	"github.com/edwingeng/live/internal"
	"math"
)

var (
	liveZero = Data{&internal.Data{}}
	liveNum1 = Data{&internal.Data{N: 1}}
)

func WrapBool(v bool) Data {
	if !v {
		return liveZero
	} else {
		return liveNum1
	}
}

func wrapInteger(v int64) Data {
	switch v {
	case 0:
		return liveZero
	case 1:
		return liveNum1
	default:
		return Data{&internal.Data{N: v}}
	}
}

func WrapInt(v int) Data {
	return wrapInteger(int64(v))
}

func WrapInt8(v int8) Data {
	return wrapInteger(int64(v))
}

func WrapInt16(v int16) Data {
	return wrapInteger(int64(v))
}

func WrapInt32(v int32) Data {
	return wrapInteger(int64(v))
}

func WrapInt64(v int64) Data {
	return wrapInteger(v)
}

func WrapUint(v uint) Data {
	return wrapInteger(int64(v))
}

func WrapUint8(v uint8) Data {
	return wrapInteger(int64(v))
}

func WrapUint16(v uint16) Data {
	return wrapInteger(int64(v))
}

func WrapUint32(v uint32) Data {
	return wrapInteger(int64(v))
}

func WrapUint64(v uint64) Data {
	return wrapInteger(int64(v))
}

func WrapFloat32(v float32) Data {
	return wrapInteger(int64(math.Float32bits(v)))
}

func WrapFloat64(v float64) Data {
	return wrapInteger(int64(math.Float64bits(v)))
}

func WrapString(v string) Data {
	return Data{&internal.Data{X: []byte(v)}}
}

func WrapBytes(v []byte) Data {
	return Data{&internal.Data{X: v}}
}

func WrapComplex64(v complex64) Data {
	buf := make([]byte, 8, 8)
	binary.LittleEndian.PutUint32(buf[:4], math.Float32bits(real(v)))
	binary.LittleEndian.PutUint32(buf[4:], math.Float32bits(imag(v)))
	return Data{&internal.Data{X: buf}}
}

func WrapComplex128(v complex128) Data {
	buf := make([]byte, 16, 16)
	binary.LittleEndian.PutUint64(buf[:8], math.Float64bits(real(v)))
	binary.LittleEndian.PutUint64(buf[8:], math.Float64bits(imag(v)))
	return Data{&internal.Data{X: buf}}
}

func wrapObjectImpl(x json.Marshaler) Data {
	if x == nil {
		return liveZero
	}

	bts, err := x.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return Data{&internal.Data{X: bts}}
}

func WrapObject(obj any) Data {
	if obj == nil {
		return liveZero
	}

	if x, ok := obj.(json.Marshaler); ok {
		return wrapObjectImpl(x)
	}

	bts, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return Data{&internal.Data{X: bts}}
}

type ProtobufMarshaler interface {
	Marshal() ([]byte, error)
}

func WrapProtobufObject(obj ProtobufMarshaler) Data {
	if obj == nil {
		return liveZero
	}

	bts, err := obj.Marshal()
	if err != nil {
		panic(err)
	}
	if len(bts) == 0 {
		return liveZero
	}

	return Data{&internal.Data{X: bts}}
}

func WrapValueDirect(v any, cfg Config) Data {
	return cfg.WrapValueDirect(v)
}
