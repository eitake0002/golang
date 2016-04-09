package main

import (
  "fmt"
  "log"
  "time"
  "os"
  "bufio"
  "reflect"
)

type MyWrapper struct{
  filename string
}

func (f *MyWrapper) ReadFile() []string {
  fp, _   := os.Open(f.filename)
  reader  := bufio.NewScanner(fp)
  content := []string{}
  for reader.Scan(){
    //fmt.Println(reader.Text())
    content = append(content, reader.Text())
  }
  //fmt.Println(content)
  return content
}

func main(){
  start := time.Now()
  log.Print("Start")
  fmt.Print("Start")

  wrapper := MyWrapper{filename: "hello.go"}
  fmt.Println(reflect.TypeOf(wrapper))
  content := wrapper.ReadFile()
  fmt.Println(content)

  end := time.Now()
  log.Printf("%f sec\n", (end.Sub(start)).Seconds())
}
