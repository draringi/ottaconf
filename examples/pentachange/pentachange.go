package main

import (
	"flag"
	"fmt"
	"github.com/draringi/ottaconf"
	"github.com/draringi/ottaconf/ottadiff"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		fmt.Printf("Usage: %s  OLD NEW\n", os.Args[0])
		return
	}
	oldPath := args[0]
	newPath := args[1]
	oldSource, err := os.Open(oldPath)
	if err != nil {
		fmt.Printf("Unable to open file: %v", err)
		return
	}
	defer oldSource.Close()
	newSource, err := os.Open(newPath)
	if err != nil {
		fmt.Printf("Unable to open file: %v", err)
		return
	}
	defer newSource.Close()
	oldSchedule, err := ottaconf.Parse(oldSource)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	newSchedule, err := ottaconf.Parse(newSource)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	changes := ottadiff.Diff(oldSchedule, newSchedule)
	for _, change := range changes {
		fmt.Print(change)
	}
}
