package ottaconf

// Room stores information about various rooms used during the conference
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

// String returns the name of the room
func (r *Room) String() string {
	return r.name
}

// Name returns the name of the room
func (r *Room) Name() string {
	return r.name
}

// Events returns a list of events taking place in this room during the conference
func (r *Room) Events() []*Event {
	eventList := make([]*Event, len(r.events))
	var i int
	for _, e := range r.events {
		eventList[i] = e
		i++
	}
	return eventList
}
