package controllers

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvReader struct {
	filename string
}

func (c *CsvReader) ReadCsv() [][]string {
	file, err := os.OpenFile(c.filename, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	fields, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return fields
}
