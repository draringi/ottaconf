package ottaconf

import (
	"encoding/xml"
	"io"
)

// Conference is the main structure used, storing information on the conference, as well as what version of the schedule is contained.
// In addition, it provides indexing of events, rooms and event hosts.
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

// Parse parses a source for a xml description of the conference schedule, producing the internal representation, while index the
// relevent information.
func Parse(r io.Reader) (*Conference, error) {
	p := xml.NewDecoder(r)
	var schedule struct {
		XMLName    xml.Name    `xmll:"schedule"`
		Conference *Conference `xml:"conference"`
		Day        []*Day      `xml:"day"`
	}
	err := p.Decode(&schedule)
	if err != nil {
		return nil, err
	}
	c := schedule.Conference
	c.events = make(map[int]*Event)
	c.people = make(map[int]*Person)
	c.rooms = make(map[string]*Room)
	c.days = make(map[*Date]*Day)
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

// UnmarshalXML provides an interface to unmarshal XML encoded data about a conference into
// the internal representation.
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

// Title provides the official title of the conference
func (c *Conference) Title() string {
	return c.title
}

// Subtitle provides a short description of the conference
func (c *Conference) Subtitle() string {
	return c.subtitle
}

// Venue provides the general location of the conference
func (c *Conference) Venue() string {
	return c.venue
}

// City provides the city where the conference takes place
func (c *Conference) City() string {
	return c.city
}

// StartDate provides the date of the fist event scheduled for the conference
func (c *Conference) StartDate() *Date {
	return c.start
}

// EndDate provides the date of the last event scheduled for the conference
func (c *Conference) EndDate() *Date {
	return c.end
}

// Length provides the number of days scheduled the conference spans
func (c *Conference) Length() int {
	return c.length
}

// Release provides the schedule version in use
func (c *Conference) Release() string {
	return c.release
}

// DayChange provides the time of day the schedule considers the date to change
func (c *Conference) DayChange() Time {
	return c.dayChange
}

// TimeSlotDuration provides the smallest unit in time used to measure the duration
// and time between events
func (c *Conference) TimeSlotDuration() Time {
	return c.timeslotDuration
}
