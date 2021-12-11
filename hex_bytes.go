package encodetype

import (
	"encoding"
	"encoding/hex"
	"encoding/xml"
)

type HexBytes []byte

var _ xml.MarshalerAttr = (*HexBytes)(nil)
var _ xml.UnmarshalerAttr = (*HexBytes)(nil)
var _ encoding.TextMarshaler = (*Base64Bytes)(nil)
var _ encoding.TextUnmarshaler = (*Base64Bytes)(nil)

func (v *HexBytes) MarshalXMLAttr(name xml.Name) (attr xml.Attr, err error) {
	attr = xml.Attr{Name: name}
	if *v != nil {
		attr.Value = hex.EncodeToString(*v)
	}
	return
}

func (v *HexBytes) UnmarshalXMLAttr(attr xml.Attr) (err error) {
	*v, err = hex.DecodeString(attr.Value)
	return
}

func (v *HexBytes) MarshalText() (text []byte, err error) {
	if *v != nil {
		text = []byte(hex.EncodeToString(*v))
	} else {
		text = []byte{}
	}
	return
}

func (v *HexBytes) UnmarshalText(text []byte) (err error) {
	*v, err = hex.DecodeString(string(text))
	return
}
