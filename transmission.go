package main

import (
	"errors"
	"log"

	"github.com/hekmon/transmissionrpc"
)

var tr *transmissionrpc.Client

func initTransmission(addr, uid, pwd string) error {
	var err error
	tr, err = transmissionrpc.New(addr, uid, pwd, nil)
	if err != nil {
		log.Println("transmissionrpc init fail", err)
		return err
	}
	return nil
}

func addMagnet(link string, save string) error {
	if tr == nil {
		return errors.New("transmissionrpc not init")
	}

	t, err := tr.TorrentAdd(&transmissionrpc.TorrentAddPayload{
		Filename:    &link,
		DownloadDir: &save,
	})
	if err != nil {
		log.Println("torrent add error", err)
		return err
	}

	log.Println("add torrent", *t.ID, *t.Name, link, save)

	return nil
}

func checkState() error {
	if tr == nil {
		return errors.New("transmissionrpc not init")
	}

	torrents, err := tr.TorrentGetAll()
	if err != nil {
		log.Println("get torrent list error", err)
		return err
	}

	var ids []int64
	for _, t := range torrents {
		if *t.Status == transmissionrpc.TorrentStatusSeed {
			log.Println("complete download", *t.ID, *t.Name)
			ids = append(ids, *t.ID)
		}
	}
	if len(ids) > 0 {
		err := tr.TorrentStopIDs(ids)
		if err != nil {
			log.Println("torrent stop error", err)
		}
	}
	return nil
}
