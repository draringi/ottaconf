package ottaconf

type Room struct {
	name   string
	events map[int]*Event
}

func (r *Room) addEvent(e *Event) {
	if r.events == nil {
		r.events = make(map[int]*Event)
	}
	r.events[e.id] = e
}

func (r *Room) String() string {
	return r.name
}
