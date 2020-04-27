package main

import (
	"io"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseQQLink(t *testing.T) {
	link := "https://torrentqq27.com/torrent/med/160279.html"
	ll, _ := getBody(link, func(body io.Reader) (string, error) {
		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(body)
		if err != nil {
			t.Fatal(err)
		}

		result := strings.Split(doc.Find(".view-torrent > table > tbody > tr > td > ul > li").Eq(0).Text(), " ")[2]

		return result, nil
	})
	t.Log(ll)

}
