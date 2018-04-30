// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package esi

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

func easyjson24365b2aDecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetInsurancePrices200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetInsurancePrices200OkList, 0, 2)
			} else {
				*out = GetInsurancePrices200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetInsurancePrices200Ok
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson24365b2aEncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetInsurancePrices200OkList) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v GetInsurancePrices200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson24365b2aEncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInsurancePrices200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson24365b2aEncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInsurancePrices200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson24365b2aDecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInsurancePrices200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson24365b2aDecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson24365b2aDecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetInsurancePrices200Ok) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "levels":
			if in.IsNull() {
				in.Skip()
				out.Levels = nil
			} else {
				in.Delim('[')
				if out.Levels == nil {
					if !in.IsDelim(']') {
						out.Levels = make([]GetInsurancePricesLevel, 0, 2)
					} else {
						out.Levels = []GetInsurancePricesLevel{}
					}
				} else {
					out.Levels = (out.Levels)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetInsurancePricesLevel
					(v4).UnmarshalEasyJSON(in)
					out.Levels = append(out.Levels, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "type_id":
			out.TypeId = int32(in.Int32())
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
func easyjson24365b2aEncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetInsurancePrices200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Levels) != 0 {
		const prefix string = ",\"levels\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Levels {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.TypeId != 0 {
		const prefix string = ",\"type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetInsurancePrices200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson24365b2aEncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetInsurancePrices200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson24365b2aEncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetInsurancePrices200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson24365b2aDecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetInsurancePrices200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson24365b2aDecodeGithubComAntihaxGoesiEsi1(l, v)
}