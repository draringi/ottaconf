package ottaconf

import "encoding/xml"

type Person struct {
	id     int
	name   string
	events []*Event
}

func (p *Person) addEvent(e *Event) {
	p.events = append(p.events, e)
}

func (p *Person) String() string {
	return p.name
}

func (p *Person) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data struct {
		XMLName xml.Name `xml:"person"`
		Id      int      `xml:"id,attr"`
		Name    string   `xml:",chardata"`
	}
	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	p.id = data.Id
	p.name = data.Name
	return nil
}
