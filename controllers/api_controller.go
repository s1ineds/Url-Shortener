package controllers

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type CsvRecord struct {
	OriginalUrl string `json:"originalUrl"`
	ShortenUrl  string `json:"shortenUrl"`
}

func GetUrls(c *gin.Context) {
	csvReader := CsvReader{filename: "data/data.csv"}

	records := csvReader.ReadCsv()

	var sliceOfRecords []CsvRecord

	for _, line := range records {
		tmpObj := CsvRecord{OriginalUrl: line[0], ShortenUrl: line[1]}
		sliceOfRecords = append(sliceOfRecords, tmpObj)
	}

	resultSlice, err := json.MarshalIndent(sliceOfRecords, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(resultSlice))
}

func GetShortUrl(c *gin.Context) {
	var sliceOfRecs []CsvRecord

	shortUrl := generateShortURL()
	shortUrl = path.Join("exmpl.cm", shortUrl)

	file, err := os.OpenFile("data/data.csv", os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	record := append(record, c.Param("originalUrl"), shortUrl)

	csvWriter.Write(record)
	csvWriter.Flush()

	sliceOfRecs = append(sliceOfRecs, CsvRecord{OriginalUrl: c.Param("originalUrl"), ShortenUrl: shortUrl})

	resultSlice, err := json.MarshalIndent(sliceOfRecs, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(resultSlice))
}
