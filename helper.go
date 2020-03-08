package live

import (
	"encoding/binary"
	"math"
	"reflect"
	"strings"

	"github.com/edwingeng/live/internal"
)

var (
	UnsafeMode bool
)

type Helper struct {
	Whitelist []string
	Blacklist []string
}

func (h Helper) NewBool(v bool) Data {
	if v {
		return Data{d: internal.Data{N: 1}}
	} else {
		return Data{}
	}
}

func (h Helper) NewInt(v int) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewInt8(v int8) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewInt16(v int16) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewInt32(v int32) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewInt64(v int64) Data {
	return Data{d: internal.Data{
		N: v,
	}}
}

func (h Helper) NewUint(v uint) Data {
	switch v {
	case 0:
		return Data{}
	default:
		var buf [10]byte
		n := binary.PutUvarint(buf[:], uint64(v))
		return Data{d: internal.Data{
			X: buf[:n],
		}}
	}
}

func (h Helper) NewUint8(v uint8) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewUint16(v uint16) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewUint32(v uint32) Data {
	return Data{d: internal.Data{
		N: int64(v),
	}}
}

func (h Helper) NewUint64(v uint64) Data {
	switch v {
	case 0:
		return Data{}
	default:
		var buf [10]byte
		n := binary.PutUvarint(buf[:], v)
		return Data{d: internal.Data{
			X: buf[:n],
		}}
	}
}

func (h Helper) NewFloat32(v float32) Data {
	b := math.Float32bits(v)
	return Data{d: internal.Data{
		N: int64(b),
	}}
}

func (h Helper) NewFloat64(v float64) Data {
	b := math.Float64bits(v)
	var buf [10]byte
	n := binary.PutUvarint(buf[:], b)
	return Data{d: internal.Data{
		X: buf[:n],
	}}
}

func (h Helper) NewString(v string) Data {
	return Data{d: internal.Data{
		X: []byte(v),
	}}
}

func (h Helper) NewProtobufObj(v interface{}) Data {
	x, ok := v.(interface {
		Marshal() ([]byte, error)
		Unmarshal([]byte) error
	})
	if !ok {
		panic("v is not protobuf compatible")
	}
	if x != nil {
		bts, err := x.Marshal()
		if err != nil {
			panic(err)
		}
		return Data{d: internal.Data{
			X: bts,
		}}
	} else {
		return Data{}
	}
}

func (h Helper) NewJSONObj(v interface{}) Data {
	x, ok := v.(interface {
		Marshal() ([]byte, error)
		Unmarshal([]byte) error
	})
	if !ok {
		panic("v is not JSON compatible")
	}
	if x != nil {
		bts, err := x.Marshal()
		if err != nil {
			panic(err)
		}
		return Data{d: internal.Data{
			X: bts,
		}}
	} else {
		return Data{}
	}
}

func (h Helper) NewValue(v interface{}) Data {
	if UnsafeMode {
		return Data{v: v}
	}
	h.checkType(reflect.TypeOf(v))
	return Data{v: v}
}

func (h Helper) checkType(t reflect.Type) {
	switch t.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Array:
		h.checkType(t.Elem())
	case reflect.Chan:
		h.checkType(t.Elem())
	case reflect.Func:
		panic("live data does not support func")
	case reflect.Interface:
		panic("live data does not support interface")
	case reflect.Map:
		h.checkMapKeyType(t)
		h.checkType(t.Elem())
	case reflect.Ptr:
		h.checkType(t.Elem())
	case reflect.Slice:
		h.checkType(t.Elem())
	case reflect.String:
	case reflect.Struct:
		pkgPath := t.PkgPath()
		if len(h.Whitelist) > 0 {
			var found bool
			for _, x := range h.Whitelist {
				if x == pkgPath {
					found = true
					break
				}
			}
			if !found {
				panic(pkgPath + " is NOT in the whitelist of live.Helper")
			}
		}
		if len(h.Blacklist) > 0 {
			for _, x := range h.Blacklist {
				if strings.HasPrefix(x, pkgPath) {
					if n := len(pkgPath); len(x) == n || x[n] == '/' {
						panic(pkgPath + " is in the blacklist of live.Helper")
					}
				}
			}
		}
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			h.checkType(f.Type)
		}
	case reflect.UnsafePointer:
		panic("live data does not support unsafe pointer")
	default:
		panic("impossible")
	}
}

func (h Helper) checkMapKeyType(t reflect.Type) {
	switch t.Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.String:
	case reflect.Uintptr:
	default:
		panic("unsupported map key type: " + t.Name())
	}
}
