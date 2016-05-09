package mongo

import (
	"testing"
)

func TestInsert(t *testing.T) {
	url := "http://yahoo.co.jp"
	actual := Insert(url)

	if actual != nil {
		t.Errorf("got %v urls want more than %v", actual, nil)
	}
}
