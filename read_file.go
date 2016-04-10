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
  "github.com/yukihir0/gec"
)

type MyWrapper struct{
  filename string
}

func (f *MyWrapper) ReadFile() []string {
  fp, _    := os.Open(f.filename)
  reader   := bufio.NewScanner(fp)
  line_num := 0
  for reader.Scan(){
    line_num++
  }

  fp2, _  := os.Open(f.filename)
  reader2 := bufio.NewScanner(fp2)
  url_list := make([]string, line_num)
  i := 0
  for reader2.Scan(){
    url_list[i] = reader2.Text()
    i++
  }
  return url_list
}

func (f *MyWrapper) GetRequest(url string, ch chan string) string {
  resp, err := http.Get(url)
  if err != nil {
    log.Print(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  text := string(body)
  opt  := gec.NewOption()
  opt.Threashold = 150
  _, title := gec.Analyse(text, opt)
  ch <- url
  log.Println(title)
  return title
}

func main(){
  start := time.Now()
  log.Println("Start")
  fmt.Println("Start")
  
  ch := make(chan string)
  wrapper  := MyWrapper{filename: "url_list.csv"}
  url_list := wrapper.ReadFile()

  for _, v := range url_list {
    go wrapper.GetRequest(v, ch)
  }

  for i, v := range url_list {
    fmt.Println("finished :",i,v)
    <-ch
  }

  end := time.Now()
  log.Printf("Exec Time : %f sec\n", (end.Sub(start)).Seconds())
  var _ = reflect.TypeOf(1)
}
