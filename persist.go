package live

import (
	"github.com/edwingeng/live/internal"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
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

func (this *Persistent) Unmarshal(dAtA []byte) error {
	return this.d.Unmarshal(dAtA)
}

func (p Persistent) MarshalJSON() ([]byte, error) {
	return p.d.MarshalJSON()
}

func (p Persistent) MarshalEasyJSON(w *jwriter.Writer) {
	p.d.MarshalEasyJSON(w)
}

func (this *Persistent) UnmarshalJSON(dAtA []byte) error {
	return this.d.UnmarshalJSON(dAtA)
}

func (this *Persistent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	this.d.UnmarshalEasyJSON(l)
}

func (p Persistent) PeekInternalBytes() ([]byte, bool) {
	if p.d.N == 0 {
		return p.d.X, true
	}
	return nil, false
}

func (p Persistent) Data() Data {
	return Data{v: &p.d}
}
