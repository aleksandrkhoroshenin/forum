// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjsonD2b7633eDecodeForumSrcDictsModels(in *jlexer.Lexer, out *Vote) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "voice":
			out.Voice = int32(in.Int32())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels(out *jwriter.Writer, in Vote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"nickname\":"
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	{
		const prefix string = ",\"voice\":"
		out.RawString(prefix)
		out.Int32(int32(in.Voice))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Vote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Vote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Vote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Vote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels1(in *jlexer.Lexer, out *Users) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Users, 0, 8)
			} else {
				*out = Users{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *User
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(User)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeForumSrcDictsModels1(out *jwriter.Writer, in Users) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Users) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Users) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Users) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Users) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels1(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels2(in *jlexer.Lexer, out *UserUpdate) {
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
		case "about":
			out.About = string(in.String())
		case "email":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Email).UnmarshalJSON(data))
			}
		case "fullname":
			out.Fullname = string(in.String())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels2(out *jwriter.Writer, in UserUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.About != "" {
		const prefix string = ",\"about\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.About))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Email).MarshalJSON())
	}
	if in.Fullname != "" {
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels2(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels3(in *jlexer.Lexer, out *User) {
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
		case "about":
			out.About = string(in.String())
		case "email":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Email).UnmarshalJSON(data))
			}
		case "fullname":
			out.Fullname = string(in.String())
		case "nickname":
			out.Nickname = string(in.String())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels3(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	if in.About != "" {
		const prefix string = ",\"about\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.About))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Email).MarshalJSON())
	}
	{
		const prefix string = ",\"fullname\":"
		out.RawString(prefix)
		out.String(string(in.Fullname))
	}
	if in.Nickname != "" {
		const prefix string = ",\"nickname\":"
		out.RawString(prefix)
		out.String(string(in.Nickname))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels3(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels4(in *jlexer.Lexer, out *ThreadUpdate) {
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
		case "message":
			out.Message = string(in.String())
		case "title":
			out.Title = string(in.String())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels4(out *jwriter.Writer, in ThreadUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ThreadUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ThreadUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ThreadUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ThreadUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels4(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels5(in *jlexer.Lexer, out *Thread) {
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
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "id":
			out.ID = int32(in.Int32())
		case "message":
			out.Message = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "title":
			out.Title = string(in.String())
		case "votes":
			out.Votes = int32(in.Int32())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels5(out *jwriter.Writer, in Thread) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix[1:])
		out.String(string(in.Author))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int32(int32(in.ID))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	if in.Votes != 0 {
		const prefix string = ",\"votes\":"
		out.RawString(prefix)
		out.Int32(int32(in.Votes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Thread) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Thread) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Thread) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Thread) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels5(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels6(in *jlexer.Lexer, out *Status) {
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
		case "forum":
			out.Forum = int32(in.Int32())
		case "post":
			out.Post = int64(in.Int64())
		case "thread":
			out.Thread = int32(in.Int32())
		case "user":
			out.User = int32(in.Int32())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels6(out *jwriter.Writer, in Status) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"forum\":"
		out.RawString(prefix[1:])
		out.Int32(int32(in.Forum))
	}
	{
		const prefix string = ",\"post\":"
		out.RawString(prefix)
		out.Int64(int64(in.Post))
	}
	{
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int32(int32(in.Thread))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.Int32(int32(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Status) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Status) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Status) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Status) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels6(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels7(in *jlexer.Lexer, out *Posts) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Posts, 0, 8)
			} else {
				*out = Posts{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v4 *Post
			if in.IsNull() {
				in.Skip()
				v4 = nil
			} else {
				if v4 == nil {
					v4 = new(Post)
				}
				(*v4).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v4)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeForumSrcDictsModels7(out *jwriter.Writer, in Posts) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Posts) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Posts) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Posts) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Posts) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels7(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels8(in *jlexer.Lexer, out *PostUpdate) {
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
		case "message":
			out.Message = string(in.String())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels8(out *jwriter.Writer, in PostUpdate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Message != "" {
		const prefix string = ",\"message\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostUpdate) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostUpdate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostUpdate) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostUpdate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels8(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels9(in *jlexer.Lexer, out *PostFull) {
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
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(User)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(Forum)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(Post)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(Thread)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
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
func easyjsonD2b7633eEncodeForumSrcDictsModels9(out *jwriter.Writer, in PostFull) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Author != nil {
		const prefix string = ",\"author\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Author).MarshalEasyJSON(out)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	if in.Post != nil {
		const prefix string = ",\"post\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Post).MarshalEasyJSON(out)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Thread).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostFull) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostFull) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostFull) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostFull) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels9(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels10(in *jlexer.Lexer, out *Post) {
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
		case "author":
			out.Author = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "forum":
			out.Forum = string(in.String())
		case "id":
			out.ID = int64(in.Int64())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "message":
			out.Message = string(in.String())
		case "parent":
			out.Parent = int64(in.Int64())
		case "thread":
			out.Thread = int32(in.Int32())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels10(out *jwriter.Writer, in Post) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"author\":"
		out.RawString(prefix[1:])
		out.String(string(in.Author))
	}
	if true {
		const prefix string = ",\"created\":"
		out.RawString(prefix)
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		out.RawString(prefix)
		out.String(string(in.Forum))
	}
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Int64(int64(in.ID))
	}
	if in.IsEdited {
		const prefix string = ",\"isEdited\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsEdited))
	}
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix)
		out.String(string(in.Message))
	}
	if in.Parent != 0 {
		const prefix string = ",\"parent\":"
		out.RawString(prefix)
		out.Int64(int64(in.Parent))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		out.RawString(prefix)
		out.Int32(int32(in.Thread))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Post) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Post) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Post) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Post) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels10(l, v)
}
func easyjsonD2b7633eDecodeForumSrcDictsModels11(in *jlexer.Lexer, out *Forum) {
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
		case "posts":
			out.Posts = int64(in.Int64())
		case "slug":
			out.Slug = string(in.String())
		case "threads":
			out.Threads = int32(in.Int32())
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
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
func easyjsonD2b7633eEncodeForumSrcDictsModels11(out *jwriter.Writer, in Forum) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Posts != 0 {
		const prefix string = ",\"posts\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Posts))
	}
	{
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	if in.Threads != 0 {
		const prefix string = ",\"threads\":"
		out.RawString(prefix)
		out.Int32(int32(in.Threads))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Forum) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeForumSrcDictsModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Forum) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeForumSrcDictsModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Forum) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeForumSrcDictsModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Forum) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeForumSrcDictsModels11(l, v)
}
