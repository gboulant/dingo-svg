package main

import "fmt"

type Note struct {
	StringNumber int
	FretNumber   int
	Name         string
}

type Notes [][]Note

func (notes Notes) String() string {
	var table string

	nbstrings := len(notes)
	nbfrets := len(notes[0])
	var header0 string = "   | "
	var header1 string = "----"
	for j := range nbfrets {
		note := notes[0][j]
		header0 += fmt.Sprintf("F%.2d   ", note.FretNumber)
		header1 += "------"
	}
	table += fmt.Sprintln(header0)
	table += fmt.Sprintln(header1)

	for i := range nbstrings {
		stringNotes := notes[i]
		stringNumber := stringNotes[0].StringNumber
		var line string = fmt.Sprintf("S%d | ", stringNumber)
		for j := range nbfrets {
			note := stringNotes[j]
			line += fmt.Sprintf("%-6s", note.Name)
		}
		table += fmt.Sprintln(line)
	}
	return table
}

func NewNotes(nbstrings, nbfrets int) Notes {
	var notes Notes = make(Notes, nbstrings)
	for i := range nbstrings {
		notes[i] = make([]Note, nbfrets)
	}
	return notes
}

func LoadNotes(csvpath string) (Notes, error) {
	records, err := loadrecords(csvpath)
	if err != nil {
		return nil, err
	}
	header := records[0]
	nbfrets := len(header) - 1
	nbstrings := len(records) - 1
	notes := NewNotes(nbstrings, nbfrets)

	for i := range nbstrings {
		for j := range nbfrets {
			notes[i][j] = Note{StringNumber: i + 1, FretNumber: j, Name: records[i+1][j+1]}
		}
	}
	return notes, nil
}
