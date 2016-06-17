package stanza

import "encoding/xml"

var whiteList map[string]struct{} = map[string]struct{}{
	"//Person/Email/Addr": struct{}{},
	"//iq/Bind/Jid":       struct{}{},
}

type Element struct {
	lookup     map[string]string
	Space, Tag string
	Attr       []Attr
	Child      []Token
}

type Attr struct {
	Space, Tag string
	Value      string
}

type CharData struct {
	Data string
}

type Token interface {
}

func (e *Element) Find(path string) string {
	val, ok := e.lookup[path]
	if !ok {
		return "NO BITCH"
	}

	return val
}
func New(start xml.StartElement, d *xml.Decoder) Element {
	m := map[string]string{}
	el := createElement("/", start, d, m)
	el.lookup = m
	return el
}

func createElement(parent string, start xml.StartElement, d *xml.Decoder, m map[string]string) Element {
	e := Element{
		Space: start.Name.Space,
		Tag:   start.Name.Local,
	}

	parent = parent + "/" + e.Tag

	for _, attr := range start.Attr {
		e.Attr = append(
			e.Attr,
			Attr{
				Space: attr.Name.Space,
				Tag:   attr.Name.Local,
				Value: attr.Value,
			},
		)
	}

	e.Child = createChildren(parent, d, m)
	return e
}

func createChildren(parent string, d *xml.Decoder, m map[string]string) (children []Token) {
	for {
		token, err := d.Token()
		if err != nil {
			return
		}

		switch elem := token.(type) {
		case xml.StartElement:
			el := createElement(parent, elem, d, m)
			children = append(children, el)
		case xml.EndElement:
			return
		case xml.CharData:
			data := string(elem)
			if len(data) <= 0 {
				return
			}

			_, add := whiteList[parent]
			if add {
				m[parent] = data
			}
			children = append(children, CharData{Data: data})
		}
	}
}
