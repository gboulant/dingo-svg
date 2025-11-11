package main

import (
	"fmt"
)

func demo01_table() error {
	notes, err := LoadNotes("guitarneck.csv")
	if err != nil {
		return err
	}
	fmt.Println(notes)
	return nil
}

func demo02_sketch() error {
	notes, err := LoadNotes("guitarneck.csv")
	if err != nil {
		return err
	}
	return makesketch(notes, "output.guitarneck.svg")
}

func main() {
	demo01_table()
	demo02_sketch()
}
