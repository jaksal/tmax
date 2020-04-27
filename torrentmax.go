package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func torrentmaxWork(conf *Config) {
	//log.Println("check start", conf)

	for _, s := range conf.Search {

		items, err := getList(conf.Site+"/max/"+s.Category, func(body io.Reader) ([]*Item, error) {

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			var result []*Item
			doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
				// For each item found, get the band and title
				//id, _ := strconv.Atoi(s.Find(".wr-num").Text())
				link, _ := s.Find(".wr-subject > a").Attr("href")
				id, _ := strconv.Atoi(link[strings.LastIndex(link, "/")+1:])

				title := strings.TrimSpace(s.Find(".wr-subject > a").Text())

				result = append(result, &Item{
					ID:    id,
					Title: title,
					//Link"," link,
				})
			})

			return result, nil
		})
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
			link := fmt.Sprintf("%s/link?bo_table=%s&wr_id=%d&no=1", conf.Site, s.Category, i.ID)
			magnet, err := getRedirectLink(link)
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
