package ottaconf

type Event struct {
	id          int
	room        *Room
	day         *Date
	start       Time
	duration    Time
	title       string
	subtitle    string
	slug        string
	track       *Track
	eventType   *EventType
	language    string
	abstract    string
	description string
	people      []*Person
	links       []*Link
}
