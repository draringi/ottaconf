package ottaconf

type Room struct {
	name   string
	events map[int]*Event
}
