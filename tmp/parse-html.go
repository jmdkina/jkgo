package main

import (
	"jk/jklog"

	"github.com/goquery"
)

func ExampleScrape() {
	doc, err := goquery.NewDocument("http://www.rfcreader.com/#rfc3550")

	// doc, err := goquery.NewDocument("http://192.168.6.151")
	if err != nil {
		jklog.L().Errorln("error open the document:", err)
	}
	table := doc.Find("div")
	value := table.Text()
	jklog.L().Infoln("the table: ", value)
	// doc.Find("table").Each(func(i int, s *goquery.Selection) {
	// jklog.L().Infof("Review %d: %s \n", i, s.Text())
	// })
}

func main() {
	ExampleScrape()
}
