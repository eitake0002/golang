package main

import (
  "fmt"
  "log"
  "time"
)

func test_func(ch chan bool, num int) {
  log.Print("Start Sleep", num)
  time.Sleep(5 * time.Second)
  log.Print("End Sleep")
  ch <- true
}

func main(){
  fmt.Println("Start")
  time.Sleep(1 * time.Second)
  ch := make(chan bool)
	go test_func(ch, 0)
  for i := 0; i < 5; i++ {
    go test_func(ch, i)
  }
  <-ch
  log.Print("End all")
}
