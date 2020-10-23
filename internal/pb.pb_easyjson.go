// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package internal

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonEbc20654DecodeGithubComEdwingengLiveInternal(in *jlexer.Lexer, out *Data) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "X":
			if in.IsNull() {
				in.Skip()
				out.X = nil
			} else {
				out.X = in.Bytes()
			}
		case "N":
			out.N = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonEbc20654EncodeGithubComEdwingengLiveInternal(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.X) != 0 {
		const prefix string = ",\"X\":"
		first = false
		out.RawString(prefix[1:])
		out.Base64Bytes(in.X)
	}
	if in.N != 0 {
		const prefix string = ",\"N\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.N))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Data) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonEbc20654EncodeGithubComEdwingengLiveInternal(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Data) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonEbc20654EncodeGithubComEdwingengLiveInternal(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Data) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonEbc20654DecodeGithubComEdwingengLiveInternal(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonEbc20654DecodeGithubComEdwingengLiveInternal(l, v)
}
