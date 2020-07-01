package live

import (
	"encoding/json"
	"reflect"
)

type Persistent Data

func (p Persistent) Marshal() (dAtA []byte, err error) {
	if p.v == nil {
		return nil, nil
	}
	data, ok := p.v.(interface {
		Marshal() (dAtA []byte, err error)
	})
	if !ok {
		panic("Marshal does not support type " + reflect.TypeOf(p.v).Name())
	}
	return data.Marshal()
}

func (p Persistent) MarshalTo(dAtA []byte) (int, error) {
	if p.v == nil {
		return 0, nil
	}
	data, ok := p.v.(interface {
		MarshalTo(dAtA []byte) (int, error)
	})
	if !ok {
		panic("MarshalTo does not support type " + reflect.TypeOf(p.v).Name())
	}
	return data.MarshalTo(dAtA)
}

func (p Persistent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	if p.v == nil {
		return 0, nil
	}
	data, ok := p.v.(interface {
		MarshalToSizedBuffer(dAtA []byte) (int, error)
	})
	if !ok {
		panic("MarshalToSizedBuffer does not support type " + reflect.TypeOf(p.v).Name())
	}
	return data.MarshalToSizedBuffer(dAtA)
}

func (p Persistent) Size() (n int) {
	if p.v == nil {
		return 0
	}
	data, ok := p.v.(interface {
		Size() (n int)
	})
	if !ok {
		panic("Size does not support type " + reflect.TypeOf(p.v).Name())
	}
	return data.Size()
}

func (p Persistent) MarshalJSON() ([]byte, error) {
	if p.v == nil {
		return nil, nil
	}
	return json.Marshal(p.v)
}
