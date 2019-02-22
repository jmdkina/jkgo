package main

import (
	"github.com/anaskhan96/soup"
	"jk/jklog"
	"net/url"
	"strings"
	"io/ioutil"
	"os"
	"strconv"
	"math/rand"
	"time"
)

type BZInfo struct {
	url    string
	imgs   []string

	path   string

	from   int
	end    int
}

func NewBZ(url string, path string) *BZInfo {
	return &BZInfo{
		url: url,
		path: path,
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
	ioutil.WriteFile(bz.path + "/" + name, []byte(d), os.ModePerm)
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

func (bz *BZInfo) queryRandom(max int) error {
	rand.Seed( time.Now().UTC().UnixNano())
	jklog.L().Debugln("random max ", max)
	for i := 0; i < max; i++ {
		e := bz.queryRandomI()
		if e != nil {
			jklog.L().Errorln("index ", i, " : ", e)
		}
	}
	return nil
}

func (bz *BZInfo) queryRandomI() error {
	jklog.L().Infoln("query random")
    pageindex := rand.Intn(200)
    queryurl := bz.url + "&page=" + strconv.Itoa(pageindex)
    jklog.L().Infoln("query url ", queryurl)
	resp, err := soup.Get(queryurl)
	if err != nil {
		return err
	}
	//jklog.L().Infoln(resp)
	doc := soup.HTMLParse(resp)
	links := doc.FindAll("div", "class", "thumb-container-big")
	jklog.L().Infoln("len links ", len(links), " of url ", queryurl)

	linkrandom := rand.Intn(len(links)-1)
	link := links[linkrandom]
	jklog.L().Infoln("query random link index ", linkrandom)
	aurl := link.Find("a").Attrs()["href"]
	host := bz.extraceDomain()
	tourl := host + "/" + aurl
	jklog.L().Infoln("query: ", tourl)
	e := bz.queryImagePreview(tourl)

	return e
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