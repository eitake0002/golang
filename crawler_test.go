package main

import (
	"testing"
)

func TestScraping(t *testing.T) {
	test_arg := []string{"aaa", "bbb", "ccc"}
	actual := OutputSingleArray(test_arg)
	expected := true
	if actual != expected {
		t.Errorf("got %v want %v", actual, expected)
	}
}
