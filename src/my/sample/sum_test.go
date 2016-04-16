package sample

import (
  "testing"
)

func TestSum(t *testing.T){
  actual   := Sum(10, 20)
  expected := 0
  if actual != expected {
    t.Errorf("got %v expected %v", actual, expected)
  }
}
