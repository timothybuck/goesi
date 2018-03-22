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

func easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetLoyaltyStoresCorporationIdOffersNotFoundList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetLoyaltyStoresCorporationIdOffersNotFoundList, 0, 4)
			} else {
				*out = GetLoyaltyStoresCorporationIdOffersNotFoundList{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetLoyaltyStoresCorporationIdOffersNotFound
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
func easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetLoyaltyStoresCorporationIdOffersNotFoundList) {
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
func (v GetLoyaltyStoresCorporationIdOffersNotFoundList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetLoyaltyStoresCorporationIdOffersNotFoundList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersNotFoundList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersNotFoundList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetLoyaltyStoresCorporationIdOffersNotFound) {
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
		case "error":
			out.Error_ = string(in.String())
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
func easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetLoyaltyStoresCorporationIdOffersNotFound) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Error_ != "" {
		const prefix string = ",\"error\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Error_))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetLoyaltyStoresCorporationIdOffersNotFound) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetLoyaltyStoresCorporationIdOffersNotFound) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32e8fee6EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersNotFound) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetLoyaltyStoresCorporationIdOffersNotFound) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32e8fee6DecodeGithubComAntihaxGoesiEsi1(l, v)
}