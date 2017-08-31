package main

import (
	"testing"
)

func TestKeyToId(t *testing.T) {
	var bing, id string
	var ok bool

	mx := getKeyToNameTargets(availableTargets)
	my := getNameToIdTargets(availableTargets)

	if bing, ok = mx['b']; !ok && (bing != "bing") {
		t.Error("'b' key doesn't refer to 'bing'")
	}

	if id, ok = my[bing]; id != "#t101" {
		t.Error("'bing' ID is not correct")
	}
}
