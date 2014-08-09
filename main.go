package main

import log "github.com/Sirupsen/logrus"

func main() {
	for i := 48; i <= 57; i++ {
		log.Infof("%v: %v", i, string(i))
	}

	for i := 97; i <= 122; i++ {
		log.Infof("%v: %v", i, string(i))
	}

	shop := make([]byte, 2)
	shop[0] = 49
	log.Infof("%v", string(shop))
}
