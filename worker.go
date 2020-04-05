package main

import (
	"fmt"
	"log"
	"strings"
)

func findStr(src string, fs [][]string) bool {
	//log.Println("find...", src, fs)

	upperSrc := strings.ToUpper(src)

	// or check
	for _, fo := range fs {
		// and check..
		exist := func(fas []string) bool {
			for _, s := range fas {
				if !strings.Contains(upperSrc, strings.ToUpper(s)) {
					return false
				}
			}
			return true
		}(fo)

		if exist {
			// log.Println("match", src, fo)
			return true
		}
	}
	// log.Println("mismatch", src, fs)
	return false
}

func tmaxWork(conf *Config) {
	//log.Println("check start", conf)

	for _, s := range conf.Search {

		items, err := getList("https://torrentmax.net/max/" + s.Category)
		if err != nil {
			log.Println("get board list error", err)
			return
		}

		for _, i := range items {
			// 1. find
			if !findStr(i.Title, s.Finds) {
				continue
			}

			// log.Println("check db ", i)

			// 2. check db
			if exist, err := existDB(i.Title); err != nil {
				log.Println("db check error", err)
				continue
			} else {
				if exist != "" {
					// log.Println("aleady exist db", i.Title)
					continue
				}
			}

			// 3. get magnet link
			link := fmt.Sprintf("https://torrentmax.net/link?bo_table=%s&wr_id=%d&no=1", s.Category, i.ID)
			magnet, err := getMagnet(link)
			if err != nil {
				continue
			}
			//log.Println("get link", link, magnet)

			// 4. add transmission
			if err := addMagnet(magnet, s.Save); err != nil {
				continue
			}

			// 5. add db
			insertDB(i.Title, magnet)

			log.Println("new magnet", i, magnet)
		}
	}

}
