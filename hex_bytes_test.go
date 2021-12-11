package encodetype

import (
	"encoding/xml"
	"testing"
)

func TestHexBytesXmlMarshalerAttr(t *testing.T) {
	v := struct {
		XMLName string `xml:"Test"`
		V       HexBytes
	}{V: HexBytes("Hello, world!")}
	out, err := xml.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "<Test><V>48656c6c6f2c20776f726c6421</V></Test>" {
		t.Fail()
	}
}

func TestHexBytesXmlUnmarshalerAttr(t *testing.T) {
	v := struct {
		XMLName string `xml:"Test"`
		V       HexBytes
	}{}
	data := []byte(`<Test><V>48656C6C6F2C20776F726C6421</V></Test>`)
	err := xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	if string(v.V) != "Hello, world!" {
		t.Fail()
	}
}

func TestHexBytesTextMarshaler(t *testing.T) {
	v := struct {
		XMLName string   `xml:"Test"`
		V       HexBytes `xml:",chardata"`
	}{V: HexBytes("Hello, world!")}
	out, err := xml.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "<Test>48656c6c6f2c20776f726c6421</Test>" {
		t.Fail()
	}
}

func TestHexBytesTextUnmarshaler(t *testing.T) {
	v := struct {
		XMLName string   `xml:"Test"`
		V       HexBytes `xml:",chardata"`
	}{}
	data := []byte(`<Test>48656C6C6F2C20776F726C6421</Test>`)
	err := xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	if string(v.V) != "Hello, world!" {
		t.Fail()
	}
}
