package main

import (
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
	for char := 48; char <= 57; char++ {
		setShopIDAtIndexToChar(shopID, index, byte(char), shopIDs)
	}

	for char := 97; char <= 122; char++ {
		setShopIDAtIndexToChar(shopID, index, byte(char), shopIDs)
	}
}

func main() {
	shopIDs := make(chan string, 2000)

	saveFile, err := os.OpenFile("logs/possible_shops.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("error opening save file: %v", err)
	}
	defer saveFile.Close()

	shopChecker := NewShopChecker(1000, shopIDs, saveFile)
	shopChecker.Work()
	lastShopID := shopChecker.LastShopID()
	lastShopID[len(lastShopID)-1]++

	for shopLength := len(lastShopID); shopLength <= 7; shopLength++ {
		shopID := make([]byte, shopLength)

		buildShopIDAtIndex(shopID, 0, shopIDs)
	}
}
