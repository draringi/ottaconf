package ottaconf

import "encoding/xml"

// Day stores information on the various days within the conference's schedule, as stored in the XML file describing the conference.
type Day struct {
	index        int
	date         *Date
	eventsByRoom map[string][]*Event
	events       []*Event
}

// UnmarshalXML provides an interface to unmarshal XML encoded data about a
// day into the internal representation.
func (d *Day) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	d.eventsByRoom = make(map[string][]*Event)
	var data struct {
		Index  int      `xml:"index,attr"`
		Date   string   `xml:"date,attr"`
		Events []*Event `xml:"room>event"`
	}
	err := dec.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	d.index = data.Index
	d.date, err = AtoDate(data.Date)
	if err != nil {
		return err
	}
	d.events = data.Events
	for _, e := range d.events {
		r := e.room.name
		eList := append(d.eventsByRoom[r], e)
		d.eventsByRoom[r] = eList
	}
	return nil
}

// Events provides a list of the day's events.
func (d *Day) Events() []*Event {
	retValue := make([]*Event, len(d.events))
	copy(d.events, retValue)
	return retValue
}
