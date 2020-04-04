package main

import (
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	err := initDB("test.db")
	defer os.Remove("test.db")
	if err != nil {
		t.Fatal(err)
	}

	if v, err := existDB("test123"); err != nil {
		t.Error(err)
	} else {
		if v != "" {
			t.Error("mismatch", v)
		}
	}

	if err := insertDB("test123", "test12333"); err != nil {
		t.Error(err)
	}

	if v, err := existDB("test123"); err != nil {
		t.Error(err)
	} else {
		if v != "test12333" {
			t.Error("mismatch", v)
		}
	}

	deleteDB("test123")

	if v, err := existDB("test123"); err != nil {
		t.Error(err)
	} else {
		if v != "" {
			t.Error("mismatch", v)
		}
	}

}
