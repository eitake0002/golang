// Main package for crawling
package main

import (
	_ "bufio"
	"flag"
	"fmt"
	_ "io/ioutil"
	"log"
	_ "net/http"
	_ "os"
	_ "os/exec"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Scraping link url from anchor tag.
// Using "goquery" library.
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

// To output single string array.
func OutputSingleArray(single_url_list []string) {
	for _, v := range single_url_list {
		fmt.Println(v)
	}
}

// To ouput double string array.
func OutputDoubleArray(ary [][]string) {
	for _, v := range ary {
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}

// Formatting [][]string to []string url_list.
func FmtToSingleArray(url_list [][]string) []string {
	var single_url_list []string
	for _, v := range url_list {
		for _, v2 := range v {
			single_url_list = append(single_url_list, v2)
		}
	}
	return single_url_list
}

func main() {
	log.Println("Start")
	start := time.Now()

	depth := flag.Int("l", 3, "link depth")
	domain := flag.String("d", "http://yahoo.co.jp", "first domain")
	flag.Parse()

	var url_list [][]string
	for i := 0; i < *depth; i++ {
		log.Println(i)
		var urls []string
		if i == 0 {
			urls = Scraping(*domain)
		} else {
			for _, v := range url_list[i-1] {
				urls = Scraping(v)
			}
		}
		url_list = append(url_list, urls)
	}

	single_url_list := FmtToSingleArray(url_list)

	OutputSingleArray(single_url_list)

	end := time.Now()
	log.Printf("%f SEC\n", (end.Sub(start)).Seconds())
}
