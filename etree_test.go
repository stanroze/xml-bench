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

func runEtree(n int) {
	for i := 0; i < n; i++ {
		etreeParse()
	}
}

func runXmlDecode(n int) {
	for i := 0; i < n; i++ {
		xmlDecode()
	}
}

func runXmlMarshal(n int) {
	for i := 0; i < n; i++ {
		xmlMarshal()
	}
}

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

func BenchmarkXmlDecode1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlDecode(1000)
	}
}

func BenchmarkXmlDecode10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlDecode(10000)
	}
}

func BenchmarkXmlDecode100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlDecode(100000)
	}
}

func BenchmarkEtree(b *testing.B) {

	for n := 0; n < b.N; n++ {
		etreeParse()
	}
}

func BenchmarkEtree1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runEtree(1000)
	}
}

func BenchmarkEtree10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runEtree(10000)
	}
}

func BenchmarkEtree100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runEtree(100000)
	}
}

func BenchmarkXmlMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xmlMarshal()
	}
}

func BenchmarkXmlMarshal1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlMarshal(1000)
	}
}

func BenchmarkXmlMarshal10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlMarshal(10000)
	}
}

func BenchmarkXmlMarshal100000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runXmlMarshal(100000)
	}
}
