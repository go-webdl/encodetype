package encodetype

import (
	"encoding/xml"
	"testing"
)

func TestBase64BytesXmlMarshalerAttr(t *testing.T) {
	v := struct {
		XMLName string `xml:"Test"`
		V       Base64Bytes
	}{V: Base64Bytes("Hello, world!")}
	out, err := xml.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "<Test><V>SGVsbG8sIHdvcmxkIQ==</V></Test>" {
		t.Fail()
	}
}

func TestBase64BytesXmlUnmarshalerAttr(t *testing.T) {
	v := struct {
		XMLName string `xml:"Test"`
		V       Base64Bytes
	}{}
	data := []byte(`<Test><V>SGVsbG8sIHdvcmxkIQ==</V></Test>`)
	err := xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	if string(v.V) != "Hello, world!" {
		t.Fail()
	}
}

func TestBase64BytesTextMarshaler(t *testing.T) {
	v := struct {
		XMLName string      `xml:"Test"`
		V       Base64Bytes `xml:",chardata"`
	}{V: Base64Bytes("Hello, world!")}
	out, err := xml.Marshal(&v)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "<Test>SGVsbG8sIHdvcmxkIQ==</Test>" {
		t.Fail()
	}
}

func TestBase64BytesTextUnmarshaler(t *testing.T) {
	v := struct {
		XMLName string      `xml:"Test"`
		V       Base64Bytes `xml:",chardata"`
	}{}
	data := []byte(`<Test>SGVsbG8sIHdvcmxkIQ==</Test>`)
	err := xml.Unmarshal(data, &v)
	if err != nil {
		t.Fatal(err)
	}
	if string(v.V) != "Hello, world!" {
		t.Fail()
	}
}
