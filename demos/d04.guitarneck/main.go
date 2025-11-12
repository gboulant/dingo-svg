package main

import (
	"fmt"
)

const (
	csvnamespath = "datafiles/guitarneck.names.csv"
	csvfreqspath = "datafiles/guitarneck.frequencies.csv"
)

func demo01_table() error {
	notes, err := LoadNotesData(csvnamespath, csvfreqspath)
	if err != nil {
		return err
	}
	fmt.Println(notes)
	return nil
}

func demo02_sketch() error {
	notes, err := LoadNotesData(csvnamespath, csvfreqspath)
	if err != nil {
		return err
	}
	return makesketch(notes, "output.guitarneck.svg")
}

func main() {
	demo01_table()
	demo02_sketch()
}
