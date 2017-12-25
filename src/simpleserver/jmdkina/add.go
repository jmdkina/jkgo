package simpleserver

import (
	"golanger.com/utils"
	"jk/jklog"
	. "simpleserver"
	// "jkdbs"
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

func (s *JmdkinaAdd) encodeImages(basepath, path string) error {
	en := img.JKEncoderImage{
		ScanPath: basepath + "/" + path,
		SavePos:  basepath + "/q10/" + path,
		Scale:    24,
		Quality:  40,
	}
	jklog.L().Debugf("Encode images 1 of %s to %s\n", en.ScanPath, en.SavePos)
	en.JK_convertWithPath()
	jklog.L().Debugf("Encode images 1 done\n")
	en2 := img.JKEncoderImage{
		ScanPath: basepath + "/" + path,
		SavePos:  basepath + "/q60/" + path,
		Scale:    40,
		Quality:  60,
	}
	jklog.L().Debugf("Encode images 2 of %s to %s\n", en2.ScanPath, en2.SavePos)
	en2.JK_convertWithPath()
	jklog.L().Debugf("Encode images 2 done\n")
	return nil
}

func (s *JmdkinaAdd) saveImages(r *http.Request) {
	path := r.FormValue("path")
	tag := r.FormValue("tag")
	author := r.FormValue("author")
	content := r.FormValue("content")
	r.ParseMultipartForm(32 << 20)

	data := []utils.M{}

	bconfig := GlobalBaseConfig()
	basepath := bconfig.PicsPath + "/" + path
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
		go func() {
			s.encodeImages(bconfig.PicsPath, path)
		}()
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
