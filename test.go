package main

import (
  "log"
)

type MyStruct struct {
  name string
  age  int
}

func (m *MyStruct) MyFunc() {
  log.Println("MyFunc()")
  m.age = 10
}

func main() {
  var my_struct MyStruct
  my_struct.MyFunc()
  log.Println(my_struct.age)
}
