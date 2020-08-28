package live

import (
	"github.com/edwingeng/live/internal"
)

type Persistent struct {
	d internal.Data
}

func (p Persistent) Marshal() (dAtA []byte, err error) {
	return p.d.Marshal()
}

func (p Persistent) MarshalTo(dAtA []byte) (int, error) {
	return p.d.MarshalTo(dAtA)
}

func (p Persistent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	return p.d.MarshalToSizedBuffer(dAtA)
}

func (p Persistent) Size() (n int) {
	return p.d.Size()
}

func (p Persistent) MarshalJSON() ([]byte, error) {
	return p.d.MarshalJSON()
}

func (p Persistent) PeekInternalBytes() ([]byte, bool) {
	if p.d.N == 0 {
		return p.d.X, true
	}
	return nil, false
}
