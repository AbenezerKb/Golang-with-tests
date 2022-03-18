package code

import (
	"fmt"
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	fmt.Printf("checking: %v\n", url)
	if url == "waat://furhurterwe.geds" {

		return false
	}

	return true
}

func TestCheckWebsite(t *testing.T) {

	websites := []string{
		"http://google.com",
		"http://youtube.com",
		"waat://furhurterwe.geds",
	}
	result := checkwebsites(mockWebsiteChecker, websites)

	want := map[string]bool{
		"http://google.com":       true,
		"http://youtube.com":      true,
		"waat://furhurterwe.geds": false}

	got := result

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}
