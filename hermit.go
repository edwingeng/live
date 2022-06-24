package live

import "github.com/edwingeng/live/internal"

type Hermit struct {
	d internal.Data
}

func (h Hermit) Marshal() (dAtA []byte, err error) {
	return h.d.Marshal()
}

func (h Hermit) MarshalTo(dAtA []byte) (int, error) {
	return h.d.MarshalTo(dAtA)
}

func (h Hermit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	return h.d.MarshalToSizedBuffer(dAtA)
}

func (h Hermit) Size() (n int) {
	return h.d.Size()
}

func FromHermitBinary(dAtA []byte) (Data, error) {
	var d internal.Data
	err := d.Unmarshal(dAtA)
	if err != nil {
		return Nil, err
	}
	return Data{&d}, nil
}
