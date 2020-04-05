package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetLink(t *testing.T) {
	category := "DRAMA"
	id := 12119

	link := fmt.Sprintf("https://torrentmax.net/link?bo_table=%s&wr_id=%d&no=1", category, id)
	magnet, err := getMagnet(link)
	if err != nil {
		t.Error(err)
	}
	t.Log(magnet)

}

func TestParseLink(t *testing.T) {
	link := "https://torrentmax.net/max/DRAMA/12147"

	t.Log(link[strings.LastIndex(link, "/")+1:])

}
