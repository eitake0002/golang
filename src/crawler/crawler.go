package crawler

import (
	_ "bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	_ "os/exec"
	"regexp"

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

// Spidering links with depth.
// Wrappering Scraping().
func Spidering(domain string, depth int) [][]string {
	var url_list [][]string
	for i := 0; i < depth; i++ {
		log.Println(i)
		var urls []string
		if i == 0 {
			urls = Scraping(domain)
		} else {
			for _, v := range url_list[i-1] {
				urls = Scraping(v)
			}
		}
		url_list = append(url_list, urls)
	}
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

type GetRequestStruct struct {
	BodyList []string
}

// Exec get request method for ParallelGetRequest().
func (g *GetRequestStruct) GetRequest(url string, c chan int) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	g.BodyList = append(g.BodyList, string(body))
	c <- 1
}

// Exec get request with groutine.
// Using GetRequest().
func (g *GetRequestStruct) ParallelGetRequest(url_list []string) {
	ch := make(chan int, 10)
	for _, v := range url_list {
		go g.GetRequest(v, ch)
	}
	for i := 0; i < len(url_list); i++ {
		<-ch
	}
}

// Output swithcing with option.
func OutputResult(output int, url_list []string, body_list []string) {
	if output == 1 {
		OutputSingleArray(url_list)
	} else if output == 2 {
		OutputSingleArray(body_list)
	}
}

