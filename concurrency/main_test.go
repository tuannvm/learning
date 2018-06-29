package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "an url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func mockWebsiteChecker(url string) bool {
	if url == "http://fail.webasdasd" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://gamek.vn",
		"http://fail.webasdasd",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	expectedResults := map[string]bool{
		"http://google.com":     true,
		"http://gamek.vn":       true,
		"http://fail.webasdasd": false,
	}

	if !reflect.DeepEqual(expectedResults, actualResults) {
		t.Fatalf("got %v want %v", actualResults, expectedResults)
	}
}
