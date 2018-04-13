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

func easyjson13ace497DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesOkList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetCharactersCharacterIdClonesOkList, 0, 1)
			} else {
				*out = GetCharactersCharacterIdClonesOkList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetCharactersCharacterIdClonesOk
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetCharactersCharacterIdClonesOkList) {
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
func (v GetCharactersCharacterIdClonesOkList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesOkList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOkList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOkList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetCharactersCharacterIdClonesOk) {
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
		case "last_clone_jump_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastCloneJumpDate).UnmarshalJSON(data))
			}
		case "home_location":
			(out.HomeLocation).UnmarshalEasyJSON(in)
		case "last_station_change_date":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastStationChangeDate).UnmarshalJSON(data))
			}
		case "jump_clones":
			if in.IsNull() {
				in.Skip()
				out.JumpClones = nil
			} else {
				in.Delim('[')
				if out.JumpClones == nil {
					if !in.IsDelim(']') {
						out.JumpClones = make([]GetCharactersCharacterIdClonesJumpClone, 0, 1)
					} else {
						out.JumpClones = []GetCharactersCharacterIdClonesJumpClone{}
					}
				} else {
					out.JumpClones = (out.JumpClones)[:0]
				}
				for !in.IsDelim(']') {
					var v4 GetCharactersCharacterIdClonesJumpClone
					(v4).UnmarshalEasyJSON(in)
					out.JumpClones = append(out.JumpClones, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetCharactersCharacterIdClonesOk) {
	out.RawByte('{')
	first := true
	_ = first
	if true {
		const prefix string = ",\"last_clone_jump_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastCloneJumpDate).MarshalJSON())
	}
	if true {
		const prefix string = ",\"home_location\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.HomeLocation).MarshalEasyJSON(out)
	}
	if true {
		const prefix string = ",\"last_station_change_date\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastStationChangeDate).MarshalJSON())
	}
	if len(in.JumpClones) != 0 {
		const prefix string = ",\"jump_clones\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.JumpClones {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetCharactersCharacterIdClonesOk) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetCharactersCharacterIdClonesOk) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson13ace497EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOk) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetCharactersCharacterIdClonesOk) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson13ace497DecodeGithubComAntihaxGoesiEsi1(l, v)
}
