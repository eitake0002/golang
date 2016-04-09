package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
  "github.com/PuerkitoBio/gocrawl"
  "github.com/PuerkitoBio/goquery"
)

func test_func(ch chan bool, num int) {
  log.Print("Start Sleep", num)
  log.Print("End Sleep")
  ch <- true
}

type ExampleExtender struct{
  gocrawl.DefaultExtender
}

func (x *ExampleExtender) Visit(ctx *gocrawl.URLContext, res *http.Response, doc *goquery.Document) (interface{}, bool) {
  // Use the goquery document or res.Body to manipulate the data
  // Return nil and true - let gocrawl find the links
  // fmt.Println(res)
  return nil, true
}

func main(){
  start := time.Now()
  fmt.Println("Start")
  opts := gocrawl.NewOptions(new(ExampleExtender))
  opts.MaxVisits = 5
  opts.LogFlags  = gocrawl.LogAll
  fmt.Println(opts)
  c := gocrawl.NewCrawlerWithOptions(opts)
  c.Run("http://normalizer.wp-x.jp/")
  //c.Run("http://your-news.me/")
  //fmt.Println(result)
  //ch := make(chan bool)
  //  go test_func(ch, 0)
  //for i := 0; i < 1; i++ {
  //go test_func(ch, i)
  //}
  //<-ch
  //log.Print("End all")
  end := time.Now()
  fmt.Printf("%f sec\n", (end.Sub(start)).Seconds())
}
