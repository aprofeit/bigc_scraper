package main

import log "github.com/Sirupsen/logrus"

func buildUrlAtIndex(url []byte, index int, char int) {
	url[index] = byte(char)

	log.Info(string(url))
}

func findShops(url []byte) {
	for urlIndex := 0; urlIndex < len(url); urlIndex++ {
		for i := 48; i <= 57; i++ {
			buildUrlAtIndex(url, urlIndex, i)
		}

		for i := 97; i <= 122; i++ {
			buildUrlAtIndex(url, urlIndex, i)
		}
	}
}

func main() {
	for shopLength := 1; shopLength <= 1; shopLength++ {
		shopURL := make([]byte, shopLength)

		findShops(shopURL)
	}
}
