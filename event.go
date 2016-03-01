package ottaconf

import "encoding/xml"

type Event struct {
	id          int
	room        *Room
	day         *Date
	start       Time
	duration    Time
	title       string
	subtitle    string
	slug        string
	track       string
	eventType   string
	language    string
	abstract    string
	description string
	people      []*Person
	links       []*Link
}

func (e *Event) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data struct {
		Id          int `xml:"id,attr"`
		Room        string
		Start       string
		Duration    string
		Slug        string
		Title       string
		Subtitle    string
		Track       string
		Type        string
		Language    string
		Abstract    string
		Description string
		People      []*Person `xml:"persons>person"`
		Links       []*Link   `xml:"links>link"`
	}
	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	e.id = data.Id
	e.room = &Room{name: data.Room}
	e.start, err = AtoTime(data.Start)
	if err != nil {
		return err
	}
	e.duration, err = AtoTime(data.Duration)
	if err != nil {
		return err
	}
	e.slug = data.Slug
	e.title = data.Title
	e.subtitle = data.Subtitle
	e.track = data.Track
	e.language = data.Language
	e.abstract = data.Abstract
	e.description = data.Description
	e.people = data.People
	e.links = data.Links
	return nil
}

func (e *Event) String() string {
	return e.title
}
