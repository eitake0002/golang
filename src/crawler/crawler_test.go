package crawler

import (
	"testing"
)

func TestScraping(t *testing.T) {
	test_arg := "http://yahoo.co.jp"
	actual := len(Scraping(test_arg))
	expected_more_than := 1

	if actual < expected_more_than {
		t.Errorf("got %v urls want more than %v", actual, expected_more_than)
	}
}
