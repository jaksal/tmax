package main

import "testing"

func TestTransmission(t *testing.T) {
	initTransmission()
	err := checkState()
	if err != nil {
		t.Error(err)
	}

	err = addMagnet(
		"magnet:?xt=urn:btih:3b6614d4a730d79b60cfaab67f371345db4a68e5",
		"/shares/Public/down/TV",
	)
	if err != nil {
		t.Error(err)
	}

}
