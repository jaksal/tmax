package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("start tmax")

	category := "VARIETY"

	conf, err := loadConfig("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	// init db
	if err := initDB(); err != nil {
		log.Fatal(err)
	}

	// init transmission
	if err := initTransmission(); err != nil {
		log.Fatal(err)
	}

	for _, s := range conf.Search {
		items, err := getList("https://torrentmax.net/max/" + s.Category)
		if err != nil {
			log.Fatal(err)
		}

		for _, i := range items {
			// 1. find

			// 2. check db
			if exist, err := existDB(i.Title); err != nil {
				continue
			} else {
				if exist {
					continue
				}
			}

			// 3. get magnet link
			link := fmt.Sprintf("https://torrentmax.net/link?bo_table=%s&wr_id=%d&no=1", category, i.ID)
			magnet, err := getMagnet(link)
			if err != nil {
				continue
			}
			// 4. add transmission
			if err := addMagnet(magnet, s.Save); err != nil {
				continue
			}
			log.Println("get magnet", i, magnet)
		}
	}

}
