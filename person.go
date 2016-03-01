package ottaconf

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
