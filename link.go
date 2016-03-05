package ottaconf

import "encoding/xml"

// Link represents a hyperlink object, as stored in the schedule.
// It contains a URL and descriptive text.
type Link struct {
	url  string
	text string
}

// UnmarshalXML provides an interface to unmarshal XML encoded data about a
// link into the internal representation.
func (l *Link) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var data struct {
		Href string `xml:"href,attr"`
		Text string `xml:",chardata"`
	}
	err := d.DecodeElement(&data, &start)
	if err != nil {
		return err
	}
	l.url = data.Href
	l.text = data.Text
	return nil
}

// URL returns the url that the hyperlink points to.
func (l *Link) URL() string {
	return l.url
}

// Description returns a decription of what the hyperlink points to
func (l *Link) Description() string {
	return l.text
}

// Link returns the url and description pair
func (l *Link) Link() (url, description string) {
	return l.url, l.text
}
