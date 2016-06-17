package stanza

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/ThomsonReutersEikon/etree"
	"github.com/stretchr/testify/assert"
)

type Iq struct {
	To        string `xml:"to,attr"`
	From      string `xml:"from,attr"`
	Type      string `xml:"type,attr"`
	ID        string `xml:"id,attr"`
	Bind      Bind
	Extension []AnyHolder `xml:",any"`
}

type AnyHolder struct {
	XMLName xml.Name
	XML     string `xml:",innerxml"`
}

type Bind struct {
	Jid string
}

var data = `
<iq to='juliet@capulet.com/core' type='result' id='bind-1'>
     <Bind>
       <Jid>stan.test@capulet.com/core</Jid>
     </Bind>
	 <Extra>
	 	<S>
		 	<T>Hello World</T>
		 </S>
		 <S>
		 	<T>Hello World</T>
		 </S>
		 <S>
		 	<T>Hello World</T>
		 </S>
		 <S>
		 	<T>Hello World</T>
		 </S>
		 <S>
		 	<T>Hello World</T>
		 </S>
		 <S>
		 	<T>Hello World</T>
		 </S>
	 </Extra>
   </iq>
`

func etreeParse() string {
	doc := etree.NewDocument()
	doc.ReadFromString(data)
	e := doc.FindElement("//Bind/Jid")
	return e.Text()
}

func xmlMarshal() string {
	iq := Iq{}
	xml.Unmarshal([]byte(data), &iq)
	return iq.Bind.Jid
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
			return el.Find("//iq/Bind/Jid")
		}
	}

	return ""
}

func TestXmlDecode(t *testing.T) {
	want := "stan.test@capulet.com/core"
	what := xmlDecode()
	assert.Equal(t, what, want, "XmlDecode expected a different result")
}

func TestEtree(t *testing.T) {
	want := "stan.test@capulet.com/core"
	what := etreeParse()
	assert.Equal(t, what, want, "Etree expected a different result")
}

func TestXmlMarshal(t *testing.T) {
	want := "stan.test@capulet.com/core"
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
