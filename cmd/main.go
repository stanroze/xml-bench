package main

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/ThomsonReutersEikon/etree"
	"github.com/stanroze/stanza"
)

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

func main() {
	d := xml.NewDecoder(strings.NewReader(data))

	for {
		s, err := d.Token()
		if err != nil || s == nil {
			break
		}
		switch se := s.(type) {
		case xml.StartElement:
			el := stanza.New(se, d)
			fmt.Println("elementfind", el.Find("//Person/Email/Addr"))
		}
	}

	doc := etree.NewDocument()
	doc.ReadFromString(data)
	e := doc.FindElement("//Email/Addr")
	fmt.Println("email", e.Text())

	fmt.Println("I'm out sissy")
}
