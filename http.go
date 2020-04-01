package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const site = "https://torrentmax.net/max/VARIETY"

var client = &http.Client{
	Timeout: time.Second * 10,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func getList() error {
	res, err := client.Get(site)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		fmt.Println("=================================================================")
		id := s.Find(".wr-num").Text()
		fmt.Println("id", id)

		link, _ := s.Find(".wr-subject > a").Attr("href")
		fmt.Println("link", link)
		fmt.Println("title", strings.TrimSpace(s.Find(".wr-subject > a").Text()))
	})

	return nil
}

func getContents(site string) (string, error) {

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
