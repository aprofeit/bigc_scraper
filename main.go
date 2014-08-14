package main

import (
	"fmt"
	"log"
	"os"
)

func setURLAtIndexToChar(url []byte, index int, char byte, urls chan string) {
	url[index] = char

	if index < len(url)-1 {
		buildURLAtIndex(url, index+1, urls)
	} else {
		shopURL := fmt.Sprintf("http://store-%v.mybigcommerce.com", string(url))
		urls <- shopURL
	}
}

func buildURLAtIndex(url []byte, index int, urls chan string) {
	for i := 48; i <= 57; i++ {
		setURLAtIndexToChar(url, index, byte(i), urls)
	}

	for i := 97; i <= 122; i++ {
		setURLAtIndexToChar(url, index, byte(i), urls)
	}
}

func main() {
	urls := make(chan string, 2000)

	saveFile, err := os.Create("logs/shops.txt")
	if err != nil {
		log.Fatal("creating file: %v", err)
	}
	defer saveFile.Close()

	shopChecker := NewShopChecker(1000, urls, saveFile)
	shopChecker.Work()

	for shopLength := 5; shopLength <= 7; shopLength++ {
		shopURL := make([]byte, shopLength)

		buildURLAtIndex(shopURL, 0, urls)
	}
}
