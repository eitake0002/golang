package main

import (
  "log"
  "my/sample"
)

func main() {
  log.Println("Test")
  sample.Sayhello()
  sum := sample.Sum(10, 20)
  log.Println(sum)
}
