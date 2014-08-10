package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var urls chan string
var checked int

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

	log.Printf("[%v] %v", resp.StatusCode, url)
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

func init() {
	urls = make(chan string, 1000)
}

func reportChecked() {
	secondsPassed := 0
	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		secondsPassed += 10
		log.Printf("%v checked per second", float64(checked)/float64(secondsPassed))
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		go checkShops()
	}

	go reportChecked()

	for shopLength := 5; shopLength <= 7; shopLength++ {
		shopURL := make([]byte, shopLength)

		buildURLAtIndex(shopURL, 0)
	}
}
