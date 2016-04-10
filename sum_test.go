package MyMath

import (
  "testing"
  "./my_sum"
)

func TestSum(t *testing.T){
  actual := MySum.Sum(10,20)
  expected := 30
  if actual != expected {
    t.Errorf("got %v and %v", actual, expected)
  }
  actual2 := MySum.Multi(10, 10)
  expected2 := 10
  if actual2 != expected2 {
    t.Errorf("got %v expected %v", actual, expected)
  }
}

