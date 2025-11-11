package main

import (
	"encoding/csv"
	"os"
)

func loadrecords(csvpath string) ([][]string, error) {
	file, err := os.Open(csvpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	return reader.ReadAll()
}
