package ottaconf

import "encoding/xml"

type Day struct {
	index        int
	date         *Date
	eventsByRoom map[string][]*Event
	events       []*Event
}

func (d *Day) UnmarshalXML(dec *xml.Decoder, start xml.StartElement) error {
	var data struct {
		Index  int      `xml:"index,attr"`
		Date   string   `xml:"date,attr"`
		Events []*Event `xml:"room>event"`
	}
	err := dec.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	d.index = data.Index
	d.date, err = AtoDate(data.Date)
	if err != nil {
		return err
	}
	d.events = data.Events
	return nil
}
