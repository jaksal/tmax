package main

import (
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(io.MultiWriter(&lumberjack.Logger{
		Filename:   "tmax.log",
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}, os.Stdout))

	conf, err := loadConfig("conf.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("check tmax", conf)

	// init db
	if err := initDB("tmax.db"); err != nil {
		log.Fatal(err)
	}

	// init transmission
	if err := initTransmission(conf.Transmission.URL, conf.Transmission.UserID, conf.Transmission.Passwd); err != nil {
		log.Fatal(err)
	}

	tmaxWork(conf)
}
