/*lsevents is a simple tool for listing all events at a conference.
It uses ottaconf to parse an xml file, either located on the local machine,
or on a website.
The idea originates from the early lists of events that BSDCan puts out, and
the list of events and talkers on the back of the BSDCan shirt.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/draringi/ottaconf"
	"io"
	"net/http"
	"os"
)

type msg struct {
	title     string
	people    string
	eventType string
}

func main() {
	isURL := false
	flag.BoolVar(&isURL, "U", false, "Tells lsevents that the path provided is a URL, allowing for")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("Usage: %s [-U] PATH|URL\n", os.Args[0])
		return
	}
	path := args[0]
	var source io.ReadCloser
	var err error
	if isURL {
		c := new(http.Client)
		r, err := c.Get(path)
		if err != nil {
			fmt.Printf("Error reading data: %v\n", err)
			return
		}
		source = r.Body
	} else {
		source, err = os.Open(path)
		if err != nil {
			fmt.Printf("Unable to open file: %v", err)
			return
		}
	}
	defer source.Close()
	conf, err := ottaconf.Parse(source)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	events := conf.Events()
	var lenTitle, lenPeople, lenType int
	var eventPrints []*msg
	for _, e := range events {
		m := new(msg)
		m.title = e.Title()
		if len(m.title) > lenTitle {
			lenTitle = len(m.title)
		}
		var peopleString string
		peopleList := e.People()
		for i, p := range peopleList {
			peopleString += p.Name()
			if i+1 < len(peopleList) {
				peopleString += ", "
			}
		}
		if len(peopleString) > lenPeople {
			lenPeople = len(peopleString)
		}
		m.people = peopleString
		m.eventType = e.Type()
		if len(m.eventType) > lenType {
			lenType = len(m.eventType)
		}
		eventPrints = append(eventPrints, m)
	}
	{
		space := lenTitle - 5
		sep := space >> 1
		printSpace(sep + 1)
		fmt.Printf("Title")
		if (sep << 1) < space {
			sep++
		}
		printSpace(sep + 1)
		fmt.Printf("|")
		space = lenPeople - 6
		sep = space >> 1
		printSpace(sep + 1)
		fmt.Printf("People")
		if (sep << 1) < space {
			sep++
		}
		printSpace(sep + 1)
		fmt.Printf("|")
		space = lenType - 4
		sep = space >> 1
		printSpace(sep + 1)
		fmt.Printf("Type\n")
		printDash(lenTitle + 2)
		fmt.Printf("+")
		printDash(lenPeople + 2)
		fmt.Printf("+")
		printDash(lenType + 2)
		fmt.Printf("\n")
	}
	for _, m := range eventPrints {
		fmt.Printf(" %s", m.title)
		printSpace(lenTitle - len(m.title))
		fmt.Printf(" | %s", m.people)
		printSpace(lenPeople - len(m.people))
		fmt.Printf(" | %s\n", m.eventType)
	}
}

func printSpace(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf(" ")
	}
}

func printDash(count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("-")
	}
}
