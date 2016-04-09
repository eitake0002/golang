package main

import (
  "fmt"
  "log"
  "time"
  "os"
  "bufio"
  "reflect"
  "net/http"
  "io/ioutil"
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

func (f *MyWrapper) GetRequest(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    log.Print(err)
  }
  log.Print(resp)
  body, err := ioutil.ReadAll(resp.Body)
  //log.Print(body)
  log.Print(reflect.TypeOf(body))
  return string(body)
}

func main(){
  start := time.Now()
  log.Print("Start")
  fmt.Print("Start")

  wrapper := MyWrapper{filename: "hello.go"}
  fmt.Println(reflect.TypeOf(wrapper))
  content := wrapper.ReadFile()
  fmt.Println(content)
  body := wrapper.GetRequest("http://your-news.me")
  log.Print(body)

  end := time.Now()
  log.Printf("%f sec\n", (end.Sub(start)).Seconds())
}
