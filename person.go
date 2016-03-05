package ottaconf

import "encoding/xml"

// Person stores information about people associated with the conference, including their name, id, and events hosted/associated with.
type Person struct {
	id     int
	name   string
	events []*Event
}

// Internal method for easily associating events with people
func (p *Person) addEvent(e *Event) {
	p.events = append(p.events, e)
}

// String provides a human readable version of a person, that is, their name.
func (p *Person) String() string {
	return p.name
}

// UnmarshalXML provides an interface to unmarshal XML encoded data about a
// person into the internal representation.
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

// Name returns the name of the person represented by the Person Object
func (p *Person) Name() string {
	return p.name
}

// ID returns the unique id number used to identify the person in the schedule.
func (p *Person) ID() int {
	return p.id
}

// Events returns a copy of the list of events the person is hosting/associated with.
func (p *Person) Events() []*Event {
	e := make([]*Event, len(p.events))
	copy(e, p.events)
	return e
}
