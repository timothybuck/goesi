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

func easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdFittings200OkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdFittings200OkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdFittings200OkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdFittings200Ok
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
func easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdFittings200OkList) {
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
func (v GetCharactersCharacterIdFittings200OkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFittings200OkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFittings200OkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFittings200OkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdFittings200Ok) {
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
		case "description":
			out.Description = string(in.String())
		case "fitting_id":
			out.FittingId = int32(in.Int32())
		case "items":
			if in.IsNull() {
				in.Skip()
				out.Items = nil
			} else {
				in.Delim('[')
				if out.Items == nil {
					if !in.IsDelim(']') {
						out.Items = make([]GetCharactersCharacterIdFittingsItem, 0, 5)
					} else {
						out.Items = []GetCharactersCharacterIdFittingsItem{}
					}
				} else {
					out.Items = (out.Items)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdFittingsItem
					easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi2(in, &v4)
					out.Items = append(out.Items, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "name":
			out.Name = string(in.String())
		case "ship_type_id":
			out.ShipTypeId = int32(in.Int32())
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
func easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdFittings200Ok) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Description != "" {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if in.FittingId != 0 {
		const prefix string = ",\"fitting_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FittingId))
	}
	if len(in.Items) != 0 {
		const prefix string = ",\"items\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Items {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi2(out, v6)
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.ShipTypeId != 0 {
		const prefix string = ",\"ship_type_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.ShipTypeId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdFittings200Ok) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdFittings200Ok) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdFittings200Ok) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdFittings200Ok) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi1(l, v)
}
func easyjson729ea2d8DecodeGithubComAntihaxGoesiEsi2(in *jlexer.Lexer, out *GetCharactersCharacterIdFittingsItem) {
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
		case "flag":
			out.Flag = int32(in.Int32())
		case "quantity":
			out.Quantity = int32(in.Int32())
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
func easyjson729ea2d8EncodeGithubComAntihaxGoesiEsi2(out *jwriter.Writer, in GetCharactersCharacterIdFittingsItem) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Flag != 0 {
		const prefix string = ",\"flag\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Flag))
	}
	if in.Quantity != 0 {
		const prefix string = ",\"quantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Quantity))
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
