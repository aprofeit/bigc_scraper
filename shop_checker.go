package main

import (
	"bufio"

	"os"
)

type ShopChecker struct {
	Concurrency int
	ShopIDs     chan string
	SaveFile    *os.File
	MemSave     []string
}

func NewShopChecker(concurrency int, shopIDs chan string, saveFile *os.File, memSave []string) *ShopChecker {
	return &ShopChecker{
		Concurrency: concurrency,
		ShopIDs:     shopIDs,
		SaveFile:    saveFile,
		MemSave:     memSave,
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
	c.MemSave = append(c.MemSave, shopID)
}

func (c *ShopChecker) LastShopID() []byte {
	scanner := bufio.NewScanner(c.SaveFile)
	var lastShopID string
	for scanner.Scan() {
		lastShopID = scanner.Text()
	}

	if len(lastShopID) == 0 {
		return []byte{48, 48, 48, 48, 47}
	} else {
		return []byte(lastShopID)
	}
}
