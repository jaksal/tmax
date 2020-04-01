package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("start tmax")

	magnet, _ := getContents("https://torrentmax.net/link?bo_table=VARIETY&wr_id=13795&no=1")

	fmt.Println(magnet)
}
