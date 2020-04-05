package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var client = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 1024,
	},
	Timeout: time.Duration(60) * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

// Item ...
type Item struct {
	ID    int
	Title string
	// Link  string
}

func getList(site string) ([]*Item, error) {
	// log.Println("get list ", site)
	res, err := client.Get(site)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Println("status code error", res.StatusCode, res.Status)
		return nil, fmt.Errorf("status code error %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
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
			//Link:  link,
		})
	})

	return result, nil
}

func getMagnet(site string) (string, error) {
	//log.Println("get magnet", site)

	res, err := client.Get(site)
	if err != nil {
		log.Println("get error", err)
		return "", err
	}
	if res.StatusCode != http.StatusFound {
		log.Println("status code error", res.StatusCode)
		return "", errors.New("invalid status code")
	}

	return res.Header.Get("location"), nil
}
