// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

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

func easyjsonCd93bc43DecodeForumSrcConfig(in *jlexer.Lexer, out *Config) {
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
		case "forum.service.port":
			out.Port = int(in.Int())
		case "forum.service.db.host":
			out.DbHost = string(in.String())
		case "forum.service.db.port":
			out.DbPort = uint16(in.Uint16())
		case "forum.service.db.database":
			out.DbDatabase = string(in.String())
		case "forum.service.db.user":
			out.DbUser = string(in.String())
		case "forum.service.db.password":
			out.DbPassword = string(in.String())
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
func easyjsonCd93bc43EncodeForumSrcConfig(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forum.service.port\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Port))
	}
	{
		const prefix string = ",\"forum.service.db.host\":"
		out.RawString(prefix)
		out.String(string(in.DbHost))
	}
	{
		const prefix string = ",\"forum.service.db.port\":"
		out.RawString(prefix)
		out.Uint16(uint16(in.DbPort))
	}
	{
		const prefix string = ",\"forum.service.db.database\":"
		out.RawString(prefix)
		out.String(string(in.DbDatabase))
	}
	{
		const prefix string = ",\"forum.service.db.user\":"
		out.RawString(prefix)
		out.String(string(in.DbUser))
	}
	{
		const prefix string = ",\"forum.service.db.password\":"
		out.RawString(prefix)
		out.String(string(in.DbPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonCd93bc43EncodeForumSrcConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonCd93bc43EncodeForumSrcConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonCd93bc43DecodeForumSrcConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonCd93bc43DecodeForumSrcConfig(l, v)
}
