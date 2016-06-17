package stanza

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ThomsonReutersEikon/etree"
	"github.com/stretchr/testify/assert"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}
type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}

var data = `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
      		<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
                <Value>Friends</Value>
				<Value>Squash</Value>
                <Value>Friends</Value>
				<Value>Squash</Value>
                <Value>Friends</Value>
				<Value>Squash</Value>
                <Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`

func etreeParse() string {
	doc := etree.NewDocument()
	doc.ReadFromString(data)
	e := doc.FindElement("//Email/Addr")
	return e.Text()
}

func xmlMarshal() string {
	v := Result{Name: "none", Phone: "none"}
	xml.Unmarshal([]byte(data), &v)
	return v.Email[0].Addr
}

func xmlDecode() string {
	d := xml.NewDecoder(strings.NewReader(data))
	for {
		s, err := d.Token()
		if err != nil || s == nil {
			break
		}
		switch se := s.(type) {
		case xml.StartElement:
			el := New(se, d)
			return el.Find("//Person/Email/Addr")
		}
	}

	return ""
}

func TestXmlDecode(t *testing.T) {
	want := "gre@example.com"
	what := xmlDecode()
	assert.Equal(t, what, want, "XmlDecode expected a different result")
}

func TestEtree(t *testing.T) {
	want := "gre@example.com"
	what := etreeParse()
	assert.Equal(t, what, want, "Etree expected a different result")
}

func TestXmlMarshal(t *testing.T) {
	want := "gre@example.com"
	what := xmlMarshal()
	assert.Equal(t, what, want, "XmlMarshal expected a different result")
}

func BenchmarkXmlDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xmlDecode()
	}
}

func BenchmarkEtree(b *testing.B) {

	for n := 0; n < b.N; n++ {
		etreeParse()
	}
}

func BenchmarkXmlMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xmlMarshal()
	}
}
