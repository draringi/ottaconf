package ottaconf

import "encoding/xml"

type Link struct {
	url  string
	text string
}

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

func (l *Link) URL() string {
	return l.url
}

func (l *Link) Description() string {
	return l.text
}
