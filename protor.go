package live

import "github.com/edwingeng/live/internal"

type Protor struct {
	d internal.Data
}

func NewProtor(d Data) Protor {
	if x, ok := d.v.(*internal.Data); ok {
		return Protor{*x}
	}
	return Protor{}
}

func (p Protor) Marshal() (dAtA []byte, err error) {
	return p.d.Marshal()
}

func (p Protor) MarshalTo(dAtA []byte) (int, error) {
	return p.d.MarshalTo(dAtA)
}

func (p Protor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	return p.d.MarshalToSizedBuffer(dAtA)
}

func (p Protor) Size() (n int) {
	return p.d.Size()
}

func FromProtorBytes(dAtA []byte) (Data, error) {
	var d internal.Data
	err := d.Unmarshal(dAtA)
	if err != nil {
		return Nil, err
	}
	return Data{&d}, nil
}
