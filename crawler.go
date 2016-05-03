// Main package for crawling
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/PuerkitoBio/goquery"

	"crawler"
)

// To output single string array.
func OutputSingleArray(single_url_list []string) bool {
	for _, v := range single_url_list {
		fmt.Println(v)
	}
	return true
}

// To ouput double string array.
func OutputDoubleArray(ary [][]string) {
	for _, v := range ary {
		for _, v2 := range v {
			fmt.Println(v2)
		}
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

func main() {
	log.Println("Start")
	start := time.Now()

	depth := flag.Int("l", 3, "link depth")
	domain := flag.String("d", "http://yahoo.co.jp", "first domain")
	output := flag.Int("o", 0, "0: no output, 1: output url_list, 2: output body_list")
	flag.Parse()

	url_list := crawler.Spidering(*domain, *depth)

	single_url_list := crawler.FmtToSingleArray(url_list)

	get_struct := crawler.GetRequestStruct{}
	get_struct.ParallelGetRequest(single_url_list)

	crawler.OutputResult(*output, single_url_list, get_struct.BodyList)

	end := time.Now()
	log.Printf("%f SEC\n", (end.Sub(start)).Seconds())
}
