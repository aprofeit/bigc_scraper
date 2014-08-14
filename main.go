package main

import (
	"bufio"
	"log"
	"os"
)

func setShopIDAtIndexToChar(shopID []byte, index int, char byte, shopIDs chan string) {
	shopID[index] = char

	if index < len(shopID)-1 {
		buildShopIDAtIndex(shopID, index+1, shopIDs)
	} else {
		shopIDs <- string(shopID)
	}
}

func buildShopIDAtIndex(shopID []byte, index int, shopIDs chan string) {
	for i := 48; i <= 57; i++ {
		setShopIDAtIndexToChar(shopID, index, byte(i), shopIDs)
	}

	for i := 97; i <= 122; i++ {
		setShopIDAtIndexToChar(shopID, index, byte(i), shopIDs)
	}
}

func main() {
	shopIDs := make(chan string, 2000)

	saveFile, err := os.Open("logs/shops.txt")
	if err != nil {
		log.Fatal("creating file: %v", err)
	}
	defer saveFile.Close()

	scanner := bufio.NewScanner(saveFile)
	var lastShopID string
	for scanner.Scan() {
		lastShopID = scanner.Text()
	}
	log.Printf("starting at shop id %s", lastShopID)
	log.Printf("is %v", []byte(lastShopID))

	shopChecker := NewShopChecker(1000, shopIDs, saveFile)
	shopChecker.Work()

	for shopLength := 5; shopLength <= 7; shopLength++ {
		shopID := make([]byte, shopLength)

		buildShopIDAtIndex(shopID, 0, shopIDs)
	}
}
