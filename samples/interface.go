package main

import (
  "log"
)

type WrapInterface interface {
  Calcurator
}

type Calcurator interface {
  Sum() int
  Add() int
}

type MyCalc struct {
  name string
  age  int
}

func (m *MyCalc) Sum() int {
  log.Println(m.age)
  return m.age
}

func (m *MyCalc) Add() int {
  log.Println(m.name)
  return m.age
}

func (m *MyCalc) Sub() int {
  return m.age
}

type MyStruct struct {
  name string
  age  int
}

func (m *MyStruct) Sum() int {
  return m.age
}
func (m *MyStruct) Add() int {
  return m.age
}

func main() {
  log.Println("Test")
  calc := MyCalc{"Taro", 20}
  var i Calcurator = &calc
  i.Sum()
  i.Add()

  my_struct := MyStruct{"test", 9999}
  var my_interface Calcurator = &my_struct
  log.Println(my_interface.Sum())
  log.Println(my_interface.Add())

}
