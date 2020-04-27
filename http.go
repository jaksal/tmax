package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var client *http.Client

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1024,
		},
		Timeout: time.Duration(60) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: jar,
	}

}

// Item ...
type Item struct {
	ID    int
	Title string
	Link  string
}

func getList(site string, parse func(io.Reader) ([]*Item, error)) ([]*Item, error) {

	res, err := client.Get(site)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("status code error", res.StatusCode, res.Status)
		return nil, fmt.Errorf("status code error %d %s", res.StatusCode, res.Status)
	}

	return parse(res.Body)
}

func getBody(site string, parse func(io.Reader) (string, error)) (string, error) {

	res, err := client.Get(site)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("status code error", res.StatusCode, res.Status)
		return "", fmt.Errorf("status code error %d %s", res.StatusCode, res.Status)
	}

	return parse(res.Body)
}

func getRedirectLink(site string) (string, error) {
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
