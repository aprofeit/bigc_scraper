package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ShopChecker struct {
	Concurrency int
	ShopIDs     chan string
	SaveFile    *os.File
}

func NewShopChecker(concurrency int, shopIDs chan string, saveFile *os.File) *ShopChecker {
	return &ShopChecker{
		Concurrency: concurrency,
		ShopIDs:     shopIDs,
		SaveFile:    saveFile,
	}
}

func (c *ShopChecker) Work() {
	for i := 0; i <= c.Concurrency; i++ {
		go func() {
			for {
				c.CheckShopURL(<-c.ShopIDs)
			}
		}()
	}
}

func (c *ShopChecker) CheckShopURL(shopID string) {
	url := fmt.Sprintf("http://store-%s.mybigcommerce.com", shopID)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Printf("found shop %s\n", shopID)
		_, err := c.SaveFile.WriteString(fmt.Sprintf("%s\n", shopID))
		if err != nil {
			log.Fatalf("saving shop shopID %s: %v", shopID, err)
		}
	}
}
