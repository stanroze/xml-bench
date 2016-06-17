package stanza

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ThomsonReutersEikon/etree"
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

func etreeParse() {
	doc := etree.NewDocument()
	doc.ReadFromString(data)
	doc.FindElement("//Email/Addr")
}

func xmlMarshal() {
	v := Result{Name: "none", Phone: "none"}
	xml.Unmarshal([]byte(data), &v)

}

func xmlDecode() {
	d := xml.NewDecoder(strings.NewReader(data))
	for {
		s, err := d.Token()
		if err != nil || s == nil {
			break
		}
		switch se := s.(type) {
		case xml.StartElement:
			el := New(se, d)
			el.Find("//Person/Email/Addr")
		}
	}
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

func BenchmarkMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xmlMarshal()
	}
}
