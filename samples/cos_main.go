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
  "strconv"
  "github.com/yukihir0/gec"
  "mecab"
)

type MyWrapper struct{
  filename string
}

func (f *MyWrapper) ReadFile() []string {
  fp, _      := os.Open(f.filename)
  reader     := bufio.NewScanner(fp)
  max_val, _ := strconv.Atoi(os.Args[1])
  line_num := 0
  for reader.Scan(){
    if line_num < max_val {
      line_num++
    } else {
      break
    }
  }

  fp2, _  := os.Open(f.filename)
  reader2 := bufio.NewScanner(fp2)
  url_list := make([]string, line_num)
  i := 0
  for reader2.Scan(){
    if i < max_val {
      log.Println(i, reader2.Text())
      url_list[i] = reader2.Text()
    } else {
      log.Println("breaked", i)
      break
    }
    i++
  }
  log.Println(len(url_list))
  return url_list
}

func (f *MyWrapper) GetRequest(url string, ch chan int) string {
  resp, err := http.Get(url)
  if err != nil {
    log.Print(err)
  }
  body, err := ioutil.ReadAll(resp.Body)
  log.Println(body[0:10])
  text := string(body)
  opt  := gec.NewOption()
  opt.Threashold = 150
  _, title := gec.Analyse(text, opt)
  ch <- 1
  log.Println(title)
  mecab_parsed, _ := mecab.Parse(title)
  log.Println(mecab_parsed)
  return title
}

func main(){
  start := time.Now()
  log.Println("Start")
  fmt.Println("Start")

  ch := make(chan int, 10)
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
