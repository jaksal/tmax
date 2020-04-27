package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   "tmax.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}, os.Stdout))

	cur, _ := os.Executable()
	conf, err := loadConfig(filepath.Dir(cur) + "/conf.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("check tmax", conf)

	// init db
	if err := initDB(filepath.Dir(cur) + "/tmax.db"); err != nil {
		log.Fatal(err)
	}

	// init transmission
	if err := initTransmission(conf.Transmission.URL, conf.Transmission.UserID, conf.Transmission.Passwd); err != nil {
		log.Fatal(err)
	}
	checkState()

	switch {
	case strings.Contains(conf.Site, "torrentmax"):
		torrentmaxWork(conf)
	case strings.Contains(conf.Site, "torrentqq"):
		torrentqqWork(conf)
	default:
		log.Fatal("not found torrentsite", conf.Site)
	}

}
