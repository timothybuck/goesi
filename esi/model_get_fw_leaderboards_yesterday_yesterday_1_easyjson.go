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

func easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi(in *jlexer.Lexer, out *GetFwLeaderboardsYesterdayYesterday1List) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetFwLeaderboardsYesterdayYesterday1List, 0, 8)
			} else {
				*out = GetFwLeaderboardsYesterdayYesterday1List{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetFwLeaderboardsYesterdayYesterday1
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
func easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi(out *jwriter.Writer, in GetFwLeaderboardsYesterdayYesterday1List) {
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
func (v GetFwLeaderboardsYesterdayYesterday1List) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsYesterdayYesterday1List) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsYesterdayYesterday1List) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsYesterdayYesterday1List) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi(l, v)
}
func easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi1(in *jlexer.Lexer, out *GetFwLeaderboardsYesterdayYesterday1) {
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
		case "amount":
			out.Amount = int32(in.Int32())
		case "faction_id":
			out.FactionId = int32(in.Int32())
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
func easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi1(out *jwriter.Writer, in GetFwLeaderboardsYesterdayYesterday1) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Amount != 0 {
		const prefix string = ",\"amount\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Amount))
	}
	if in.FactionId != 0 {
		const prefix string = ",\"faction_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.FactionId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetFwLeaderboardsYesterdayYesterday1) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetFwLeaderboardsYesterdayYesterday1) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE23e4f59EncodeGithubComAntihaxGoesiEsi1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetFwLeaderboardsYesterdayYesterday1) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetFwLeaderboardsYesterdayYesterday1) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE23e4f59DecodeGithubComAntihaxGoesiEsi1(l, v)
}