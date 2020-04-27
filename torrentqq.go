package main

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func torrentqqWork(conf *Config) {

	for _, s := range conf.Search {

		items, err := getList(conf.Site+"/torrent/newest.html?board="+s.Category, func(body io.Reader) ([]*Item, error) {

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(body)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			var result []*Item
			doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
				se := s.Find(".wr-subject > a").Eq(1)

				title := strings.TrimSpace(se.Text())
				link, _ := se.Attr("href")

				//log.Println(title, link)

				result = append(result, &Item{
					Title: title,
					Link:  link,
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
			magnet, err := getBody(i.Link, func(body io.Reader) (string, error) {
				// Load the HTML document
				doc, err := goquery.NewDocumentFromReader(body)
				if err != nil {
					return "", err
				}

				result := "magnet:?xt=urn:btih:" + strings.Split(doc.Find(".view-torrent > table > tbody > tr > td > ul > li").Eq(0).Text(), " ")[2]

				return result, nil
			})

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
