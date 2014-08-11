package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aybabtme/uniplot/spark"
)

var urls chan string
var checked int
var saveFile *os.File

func checkShops() {
	for {
		url := <-urls
		checkShopURL(url)
	}
}

func checkShopURL(url string) {
	resp, err := http.Get(url)

	checked++
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		saveURL(url)
	}
}

func saveURL(url string) {
	saveFile.WriteString(fmt.Sprintf("%s\n", url))
}

func setURLAtIndexToChar(url []byte, index int, char byte) {
	url[index] = char

	if index < len(url)-1 {
		buildURLAtIndex(url, index+1)
	} else {
		shopURL := fmt.Sprintf("http://store-%v.mybigcommerce.com", string(url))
		urls <- shopURL
	}
}

func buildURLAtIndex(url []byte, index int) {
	for i := 48; i <= 57; i++ {
		setURLAtIndexToChar(url, index, byte(i))
	}

	for i := 97; i <= 122; i++ {
		setURLAtIndexToChar(url, index, byte(i))
	}
}

func reportToSpark(sprk *spark.SparkStream) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		<-ticker.C
		sprk.Add(float64(checked))
		checked = 0
	}
}

func init() {
	urls = make(chan string, 5000)
}

func main() {
	for i := 0; i < 1000; i++ {
		go checkShops()
	}

	sprk := spark.Spark(1 * time.Second)
	sprk.Start()
	go reportToSpark(sprk)

	fileName := strconv.FormatInt(time.Now().Unix(), 10)
	file, err := os.Create(fmt.Sprintf("shops/%s", fileName))
	if err != nil {
		log.Fatal("creating file: %v", err)
	}
	saveFile = file

	for shopLength := 5; shopLength <= 7; shopLength++ {
		shopURL := make([]byte, shopLength)

		buildURLAtIndex(shopURL, 0)
	}
}
