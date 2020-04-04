package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

var db *bolt.DB
var bucket = []byte("tmax")

func initDB(file string) error {
	var err error
	db, err = bolt.Open(file, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

}
func closeDB() {
	if db != nil {
		db.Close()
	}
}
func insertDB(name, link string) error {
	if db == nil {
		return errors.New("not init db")
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put([]byte(name), []byte(link))
		return err
	})
}

func existDB(name string) (string, error) {
	var result string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		v := b.Get([]byte(name))
		result = string(v)
		//log.Println("get data", name, result, v)
		return nil
	})

	return result, nil
}

func deleteDB(name string) error {
	if db == nil {
		return errors.New("not init db")
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Delete([]byte(name))
		return err
	})

	return nil
}
