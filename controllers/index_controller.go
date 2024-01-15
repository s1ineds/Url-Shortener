package controllers

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var record []string = make([]string, 0)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}

func PostIndexForm(c *gin.Context) {
	postUrl := c.PostForm("url")
	var newUrl string

	file, err := os.OpenFile("data/data.csv", os.O_WRONLY|os.O_APPEND, 0775)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	for {
		shortenUrl := generateShortURL()
		newUrl = path.Join("exmpl.cm", shortenUrl)
		record = append(record, postUrl, newUrl)
		if !checkDuplicates(record) {
			csvWriter := csv.NewWriter(file)
			err := csvWriter.Write(record)
			if err != nil {
				fmt.Println("[CSV Writer] " + err.Error())
			}
			csvWriter.Flush()
			break
		}
	}

	c.HTML(http.StatusOK, "index.html", gin.H{"url": newUrl})
}

func RedirectIndexControlller(c *gin.Context) {
	url := path.Join("exmpl.cm", c.Param("urlId"))
	var redirectTo string

	file, err := os.OpenFile("data/data.csv", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			if records[i][1] == url {
				redirectTo = records[i][0]
				break
			}
		}
	}

	if strings.Contains(redirectTo, "https://") {
		c.Redirect(http.StatusMovedPermanently, redirectTo)
	} else {
		c.Redirect(http.StatusMovedPermanently, "https://"+redirectTo)
	}
}

func generateShortURL() string {
	var runes []rune = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K',
		'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'V', 'X', 'Y', 'Z', 'a', 'b',
		'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q',
		'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	var resultURL string
	for i := 0; i < 6; i++ {
		randNum := rand.Intn(len(runes) - 1)
		resultURL += string(runes[randNum])
	}

	return resultURL
}

func checkDuplicates(record []string) bool {
	file, err := os.Open("data/data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	fields, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	if len(fields) > 0 {
		for i := 0; i < len(fields); i++ {
			for j := 0; j < len(fields[i]); j++ {
				if fields[i][j] == record[1] {
					return true
				}
			}
		}
	}
	return false
}
