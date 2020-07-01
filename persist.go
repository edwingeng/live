package live

import (
	"encoding/json"
	"reflect"

	"github.com/edwingeng/live/internal"
)

type Persistent Data

func (p Persistent) Marshal() (dAtA []byte, err error) {
	if p.v == nil {
		return nil, nil
	}
	switch v := p.v.(type) {
	case *internal.Data:
		return v.Marshal()
	default:
		panic("Marshal does not support type " + reflect.TypeOf(p.v).Name())
	}
}

func (p Persistent) MarshalTo(dAtA []byte) (int, error) {
	if p.v == nil {
		return 0, nil
	}
	switch v := p.v.(type) {
	case *internal.Data:
		return v.MarshalTo(dAtA)
	default:
		panic("MarshalTo does not support type " + reflect.TypeOf(p.v).Name())
	}
}

func (p Persistent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if p.v == nil {
		return 0, nil
	}
	switch v := p.v.(type) {
	case *internal.Data:
		return v.MarshalToSizedBuffer(dAtA)
	default:
		panic("MarshalToSizedBuffer does not support type " + reflect.TypeOf(p.v).Name())
	}
}

func (p Persistent) Size() (n int) {
	if p.v == nil {
		return 0
	}
	switch v := p.v.(type) {
	case *internal.Data:
		return v.Size()
	default:
		panic("Size does not support type " + reflect.TypeOf(p.v).Name())
	}
}

func (p Persistent) MarshalJSON() ([]byte, error) {
	if p.v == nil {
		return nil, nil
	}
	return json.Marshal(p.v)
}
