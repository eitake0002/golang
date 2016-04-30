package main

import (
	_ "bufio"
	_ "fmt"
	_ "io/ioutil"
	"log"
	_ "net/http"
	_ "os"
	_ "os/exec"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Scraping(url string) []string {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var url_list []string
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		r := regexp.MustCompile(`(https://|http://)`)
		if r.MatchString(url) {
			url_list = append(url_list, url)
		}
	})
	return url_list
}

func Output2Array(ary [][]string) {
  for _, v := range ary {
    for _, v2 := range v {
      log.Println(v2)
    }
  }
}

func main() {
	log.Println("Start")
	start := time.Now()

	var url_list [][]string
	for i := 0; i < 3; i++ {
		log.Println(i)
		var urls []string
		if i == 0 {
			urls = Scraping("http://yahoo.co.jp")
		} else {
			for _, v := range url_list[i-1] {
				urls = Scraping(v)
			}
		}
		url_list = append(url_list, urls)
	}

    Output2Array(url_list)

	end := time.Now()
	log.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
