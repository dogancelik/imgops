package main

import (
	"testing"
)

func TestTargetSplit(t *testing.T) {
	QueryMap = getIdMap()

	queryList := getQueryList("google")
	if len(queryList) != 1 {
		t.Errorf("Expected 1 query: %v", queryList)
	}

	queryList = getQueryList("google, yandex, wrong")
	if len(queryList) != 2 {
		t.Error("Expected 2 queries")
	}
}
