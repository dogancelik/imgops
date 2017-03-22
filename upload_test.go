package main

import (
	"strings"
	"testing"
)

const testUrl string = "https://encrypted.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"
const testFile string = "test.png"

func TestFileUpload(t *testing.T) {
	var err error

	if _, err = UploadFile(testFile, ""); err != nil {
		t.Error(err)
	}

	if _, err = UploadFile(testFile, "google"); err != nil {
		t.Error(err)
	}

	if _, err = UploadURL(testFile, "wrong"); !strings.Contains(err.Error(), "No link") {
		t.Error(err)
	}
}

func TestUrlUpload(t *testing.T) {
	var err error

	if _, err = UploadURL(testUrl, ""); err != nil {
		t.Error(err)
	}

	if _, err = UploadURL(testUrl, "google, yandex"); err != nil {
		t.Error(err)
	}

	if _, err = UploadURL(testUrl, "wrong"); !strings.Contains(err.Error(), "No link") {
		t.Error(err)
	}
}
