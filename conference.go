package ottaconf

import (
	"encoding/xml"
	"io"
)

type Conference struct {
	title            string
	subtitle         string
	venue            string
	city             string
	start            *Date
	end              *Date
	length           int
	release          string
	dayChange        Time
	timeslotDuration Time
	events           map[int]*Event
	people           map[int]*Person
	rooms            map[string]*Room
	days             map[*Date]*Day
}

func Parse(r io.Reader) (*Conference, error) {
	p := xml.NewDecoder(r)
	var schedule struct {
		Conference *Conference
		Day        []*Day
	}
	err := p.Decode(&schedule)
	if err != nil {
		return nil, err
	}
	c := schedule.Conference
	for _, d := range schedule.Day {
		c.days[d.date] = d
		for _, e := range d.events {
			e.day = d.date
			c.events[e.id] = e
			r, ok := c.rooms[e.room.name]
			if ok {
				e.room = r
			} else {
				r = e.room
				c.rooms[r.name] = r
			}
			r.addEvent(e)
			for i, p := range e.people {
				person, ok := c.people[p.id]
				if ok {
					e.people[i] = person
				} else {
					person = p
					c.people[p.id] = person
				}
				person.addEvent(e)
			}
		}
	}
	return c, nil
}

func (c *Conference) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data struct {
		XMLName          xml.Name `xml:"conference"`
		Title            string   `xml:"title"`
		Subtitle         string   `xml:"subtitle"`
		Venue            string   `xml:"venue"`
		City             string   `xml:"city"`
		Start            string   `xml:"start"`
		End              string   `xml:"end"`
		Days             int      `xml:"days"`
		Release          string   `xml:"release"`
		DayChange        string   `xml:"day_change"`
		TimeslotDuration string   `xml:"timeslot_duration"`
	}
	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	c.title = data.Title
	c.subtitle = data.Subtitle
	c.venue = data.Venue
	c.city = data.City
	c.start, err = AtoDate(data.Start)
	if err != nil {
		return err
	}
	c.end, err = AtoDate(data.End)
	if err != nil {
		return err
	}
	c.length = data.Days
	c.release = data.Release
	c.dayChange, err = AtoTime(data.DayChange)
	if err != nil {
		return err
	}
	c.timeslotDuration, err = AtoTime(data.TimeslotDuration)
	if err != nil {
		return err
	}
	return nil
}
