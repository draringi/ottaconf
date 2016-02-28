package ottaconf

type Conference struct {
	title            string
	subtitle         string
	venue            string
	city             string
	start            *Date
	end              *Date
	days             int
	release          string
	dayChange        Time
	timeslotDuration Time
	events           map[int]*Event
}
