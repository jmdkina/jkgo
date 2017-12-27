package simpleserver

import (
	"golanger.com/utils"
	"jk/jklog"
	. "simpleserver"
	// "jkdbs"
	"image/draw"
	"io"
	img "jk/jkimage"
	"net/http"
	"os"
	. "simpleserver/dbs"
	"time"
	// "strconv"
)

type JmdkinaAdd struct {
	Base
}

func NewJmdkinaAdd(path string) *JmdkinaAdd {
	j := &JmdkinaAdd{}
	j.SetPath(path)
	j.SetFunc("/jmdkinaadd", j)
	return j
}

func (s *JmdkinaAdd) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/jmdkina/add.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *JmdkinaAdd) Post(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code != "jmdkinaadd" {
		return
	}
	cmd := r.FormValue("cmd")
	jklog.L().Debugf("jmdkinaadd page -- cmd %s\n", cmd)
	switch cmd {
	case "add":
		s.saveImages(r)
		s.WriteSerialData(w, "", 200)
		break
	}
}

func (s *JmdkinaAdd) encodeImage(basepath, path, name string) error {
	jklog.L().Infof("encode image %s %s %s\n", basepath, path, name)
	en := img.JKEncoderImage{
		Scale: 50,
	}
	srcpath := basepath + "/q60/" + path + "/" + name
	dstpath := basepath + "/q10/" + path + "/" + name
	os.MkdirAll(basepath+"/q10/"+path, os.ModeDir|os.ModePerm)
	err := en.JK_reducejpeg(dstpath, srcpath, draw.Src, 50)
	if err != nil {
		jklog.L().Errorln("encode img error : ", err)
		return err
	}
	return nil
}

func (s *JmdkinaAdd) saveImages(r *http.Request) {
	path := r.FormValue("path")
	tag := r.FormValue("tag")
	author := r.FormValue("author")
	content := r.FormValue("content")
	r.ParseMultipartForm(32 << 20)

	data := []utils.M{}

	// The pics come here is reduce to level 60
	// this program only save it ,and then reduce it to level 10
	baselevel := "/q60/"

	bconfig := GlobalBaseConfig()
	basepath := bconfig.PicsPath + baselevel + path
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		fhs := r.MultipartForm.File["jmdkinaaddfile"]
		num := len(fhs)
		jklog.L().Debugf("length of files %d\n", num)

		if num == 0 {
			return
		}

		err := os.MkdirAll(basepath, os.ModeDir|os.ModePerm)
		if err != nil {
			jklog.L().Errorf("Mkdir dir %s fail %v\n", path, err)
			return
		}

		for _, v := range fhs {
			filefullpath := basepath + "/" + v.Filename
			jklog.L().Infof("Add file name %s\n", filefullpath)
			fread, err := v.Open()
			defer fread.Close()
			if err != nil {
				jklog.L().Errorf("upload file %s fail %v\n", filefullpath, err)
				continue
			}
			fsave, err := os.Create(filefullpath)
			defer fsave.Close()
			if err != nil {
				jklog.L().Errorf("save file %s fail %v\n", filefullpath, err)
				continue
			}
			io.Copy(fsave, fread)

			// encode image
			s.encodeImage(bconfig.PicsPath, path, v.Filename)

			// Generate utils.M to save to dbs
			item := utils.M{
				"author":     author,
				"path":       path,
				"tag":        tag,
				"createtime": time.Now().Unix(),
				"updatetime": time.Now().Unix(),
				"name":       v.Filename,
				"content":    content,
			}
			data = append(data, item)
		}
		s.saveDBS(data)
	}
}

func (s *JmdkinaAdd) saveDBS(data []utils.M) error {
	jklog.L().Debugf("Save to dbs of length %d\n", len(data))
	for _, v := range data {
		GlobalDBS().Add("proj", "images", v)
	}
	return nil
}
