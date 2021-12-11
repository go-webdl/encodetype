package encodetype

import (
	"encoding"
	"encoding/base64"
	"encoding/xml"
)

type Base64Bytes []byte

var _ xml.MarshalerAttr = (*Base64Bytes)(nil)
var _ xml.UnmarshalerAttr = (*Base64Bytes)(nil)
var _ encoding.TextMarshaler = (*Base64Bytes)(nil)
var _ encoding.TextUnmarshaler = (*Base64Bytes)(nil)

func (v *Base64Bytes) MarshalXMLAttr(name xml.Name) (attr xml.Attr, err error) {
	attr = xml.Attr{Name: name}
	if *v != nil {
		attr.Value = base64.StdEncoding.EncodeToString(*v)
	}
	return
}

func (v *Base64Bytes) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	*v, err = base64.StdEncoding.DecodeString(attr.Value)
	return
}

func (v *Base64Bytes) MarshalText() (text []byte, err error) {
	if *v != nil {
		text = []byte(base64.StdEncoding.EncodeToString(*v))
	} else {
		text = []byte{}
	}
	return
}

func (v *Base64Bytes) UnmarshalText(text []byte) (err error) {
	*v, err = base64.StdEncoding.DecodeString(string(text))
	return
}
