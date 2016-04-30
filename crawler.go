package main

import(
  "log"
  _ "fmt"
  "time"
  _ "net/http"
  _ "io/ioutil"
  _ "bufio"
  _ "os"
  _ "os/exec"

  "github.com/PuerkitoBio/goquery"
)

func Scraping(url string) []string {
  doc, err := goquery.NewDocument(url)
  if err != nil {
    log.Fatal(err)
  }

  var url_list []string
  doc.Find("a").Each(func(_ int, s *goquery.Selection){
    url, _ := s.Attr("href")
    url_list = append(url_list, url) 
  })
  return url_list
}

func main() {
  log.Println("Start")
  start := time.Now();

  url_list := Scraping("http://google.com")
  for _, v := range url_list {
    log.Println(v)
  }

  end := time.Now();
  log.Printf("%f秒\n",(end.Sub(start)).Seconds())
}
