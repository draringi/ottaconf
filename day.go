package ottaconf

type Day struct {
	index        int
	date         *Date
	eventsByRoom map[string][]*Event
}
