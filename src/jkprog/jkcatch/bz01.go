package main

import (
	"github.com/anaskhan96/soup"
	"jk/jklog"
	"net/url"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
)

type BZInfo struct {
	url    string
	imgs   []string

	from   int
	end    int
}

func NewBZ(url string) *BZInfo {
	return &BZInfo{
		url: url,
	}
}

func (bz *BZInfo) setPage(from, end int) {
	bz.from = from
	bz.end = end
}

func (bz *BZInfo) extraceDomain() string {
	url, err := url.Parse(bz.url)
	if err != nil {
		jklog.L().Errorln(err)
		return ""
	}
	return url.Scheme + "://" + url.Host
}

func (bz *BZInfo) extractName(fileurl string) string {
	url, err := url.Parse(fileurl)
	if err != nil {
		jklog.L().Errorln(err)
		return ""
	}
	index := strings.LastIndex(url.Path, "/")
	return url.Path[index+1:]
}

func (bz *BZInfo) downloadImageNow(img string) error {
	jklog.L().Infoln("download : ", img)
	d, e := soup.Get(img)
	if e != nil {
		return e
	}
	name := bz.extractName(img)
	if len(name) == 0 {
		return e
	}
	jklog.L().Infoln("name", name)
	ioutil.WriteFile("out/" + name, []byte(d), os.ModePerm)
	return nil
}

func (bz *BZInfo) queryImagePreview(link string) error {
	resp, err := soup.Get(link)
	if err != nil {
		jklog.L().Errorln(err)
		return err
	}
	doc := soup.HTMLParse(resp)
	alink := doc.Find("div", "class", "img-container-desktop")
	if alink.Error != nil {
		jklog.L().Errorln(alink.Error)
		return alink.Error
	}
	blink := alink.Find("a").Attrs()["href"]
	jklog.L().Debugln("find real link ", blink)
	bz.downloadImageNow(blink)
	//bz.imgs = append(bz.imgs, blink)
	return nil
}

func (bz *BZInfo) query() string {
	i := bz.from
	for i = bz.from; i < bz.end; i++ {
		jklog.L().Infof("index : %d\n", i)
		queryurl := bz.url + "&page=" + strconv.Itoa(i)
		resp, err := soup.Get(queryurl)
		if err != nil {
			continue
		}
		//jklog.L().Infoln(resp)
		doc := soup.HTMLParse(resp)
		links := doc.FindAll("div", "class", "thumb-container-big")
		jklog.L().Infoln("len links ", len(links), " of url ", queryurl)

		for _, link := range links {
			aurl := link.Find("a").Attrs()["href"]
			host := bz.extraceDomain()
			tourl := host + "/" + aurl
			jklog.L().Infoln("query: ", tourl)
			bz.queryImagePreview(tourl)
		}
	}
	return ""
}