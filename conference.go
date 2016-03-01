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
		for _, r := range d.rooms {
			if c.rooms[r] == nil {
				c.rooms[r] = &Room{name: r}
			}
		}
		for _, e := range d.events {
			e.day = d.date
			c.events[e.id] = e
			c.rooms[e.room.name].addEvent(e)
			e.room = c.rooms[e.room.name]
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
		Title            string
		Subtitle         string
		Venue            string
		City             string
		Start            string
		End              string
		Days             int
		Release          string
		DayChange       string
		TimeslotDuration string
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
