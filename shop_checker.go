package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ShopChecker struct {
	Concurrency int
	URLS        chan string
	SaveFile    *os.File
}

func NewShopChecker(concurrency int, urls chan string, saveFile *os.File) *ShopChecker {
	return &ShopChecker{
		Concurrency: concurrency,
		URLS:        urls,
		SaveFile:    saveFile,
	}
}

func (c *ShopChecker) Work() {
	for i := 0; i <= c.Concurrency; i++ {
		go func() {
			for {
				c.CheckShopURL(<-c.URLS)
			}
		}()
	}
}

func (c *ShopChecker) CheckShopURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Printf("found shop %s\n", url)
		_, err := c.SaveFile.WriteString(fmt.Sprintf("%s\n", url))
		if err != nil {
			log.Fatalf("saving shop url %s: %v", url, err)
		}
	}
}
