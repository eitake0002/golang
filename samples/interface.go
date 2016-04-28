package main

import (
  "log"
)

// Set interface
type Calc interface {
  Sum(int) int
  Add() int
}

// Set struct
type MyCalc struct {
  Name string
  Age  int
}

// Connect to interface.
func (m MyCalc) Sum(i int) int {
  log.Println(i)
  return i
}
// Connect to interface. 
func (m MyCalc) Add() int {
  return m.Age
}

func main() {
  log.Println("Start Interface Sample")

  // MyCalc struct must be pointer.
  var mc MyCalc
  var cal Calc
  cal = mc
  cal.Sum(10)
}
