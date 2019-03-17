package simpleserver

import (
	"io"
	"io/ioutil"
	"jk/jklog"
	"jk/jksys"
	"net/http"
	"os"
)

type NotFound struct {
	Base
}

func NewNotFound(path string) *NotFound {
	n := &NotFound{}
	n.SetPath(path)
	http.HandleFunc("/", n.ServeHttp)
	return n
}

func (b *NotFound) ServeHttp(w http.ResponseWriter, r *http.Request) {
	indexpage := GlobalBaseConfig().Index
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/"+indexpage, http.StatusFound)
	}

	sp := SimpleParse{}
	filename := b.path + "/404.html"
	jklog.L().Debugf("Not found html [%s]\n", r.URL.Path)

	if _, err := os.Stat(filename); err != nil && !os.IsExist(err) {
		sp.ParseString(w, "The page you request has go to Mars, manager deploy error, please contact to manager", "")
		return
	}

	sp.Parse(w, nil, filename)
}

type Index struct {
	Base
}

type IndexInfo struct {
	Sysinfo jksys.KFSystemInfo
}

func NewIndex(path string) *Index {
	i := &Index{}
	i.SetPath(path)
	i.SetFunc("/index", i)
	return i
}

func (b *Index) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := b.path + "/index.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	ii := IndexInfo{
		Sysinfo: *jksys.NewSystemInfo(),
	}

	err := sp.Parse(w, ii, filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (b *Index) openfile(filename string) string {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		jklog.L().Errorf("open file [%s] error %v\n", filename, err)
		return ""
	}
	return string(d)
}

func (b *Index) writefile(filename string, content string) error {
	err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
	return err
}

func (b *Index) Post(w http.ResponseWriter, r *http.Request) {
	method := r.FormValue("method")
	switch method {
	case "open":
		filename := r.FormValue("filename")
		data := b.openfile(filename)
		b.WriteSerialData(w, data, 200)
		break
	case "modify":
		filename := r.FormValue("filename")
		content := r.FormValue("content")
		code := r.FormValue("code")
		if code != "jmdcode" {
			b.WriteSerialData(w, "", 403)
			break
		}
		err := b.writefile(filename, content)
		b.WriteSerialData(w, err, 200)
		break
	}
}

type DirServer struct {
	Base
}

func NewDirServer(path string) *DirServer {
	ds := &DirServer{}
	ds.SetPath(path)
	http.HandleFunc("/dir", ds.ServeHttp)
	return ds
}

type DirServerInfo struct {
	FSS []FileServerInfo
}

func (b *DirServer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	if r.Method == "GET" {
		filename := b.path + "/dir.html"
		jklog.L().Debugf("filepath [%s]\n", filename)
		dsi := DirServerInfo{}
		fsss := GetFileServers()
		for _, v := range fsss {
			dsi.FSS = append(dsi.FSS, *v)
		}
		err := sp.Parse(w, dsi, filename)
		if err != nil {
			jklog.L().Errorln("Parse error ", err)
			return
		}
	} else if r.Method == "POST" {
		cmd := r.FormValue("jk")
		jklog.L().Debugf("addfileserver cmd: %s\n", cmd)
		switch cmd {
		case "addfileserver":
			path := r.FormValue("path")
			if len(path) > 0 {
				jklog.L().Debugf("add path [%s]\n", path)
				AddFileServer(path)
			}
			break
		case "adduploadpath":
			path := r.FormValue("path")
			if len(path) > 0 {
				jklog.L().Debugf("add upload path [%s]\n", path)
				SetFileUploadPath(path)
			}
			break
		}
	}
}

type UploadServer struct {
	Base
}

type UploadServerInfo struct {
	UploadInfo FileUploadInfo
}

func NewUploadServer(path string) *UploadServer {
	us := &UploadServer{}
	us.SetPath(path)
	http.HandleFunc("/upload", us.ServeHttp)
	return us
}

func (b *UploadServer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sp := SimpleParse{}
		filename := b.path + "/upload.html"
		jklog.L().Debugf("upload path [%s]\n", filename)
		usi := UploadServerInfo{
			UploadInfo: GetFileUploadPath(),
		}
		if len(usi.UploadInfo.Path) <= 0 {
			usi.UploadInfo.Enable = 0
		}
		err := sp.Parse(w, usi, filename)
		if err != nil {
			jklog.L().Errorln("Parse error ", err)
			return
		}
	} else {
		uploadpath := GetFileUploadPath().Path
		if len(uploadpath) <= 0 {
			jklog.L().Errorln("Have not set upload file ,exit")
			return
		}
		r.ParseMultipartForm(2 << 24)
		jklog.L().Debugln("upload file")
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			jklog.L().Errorln(err)
			return
		}
		defer file.Close()
		jklog.L().Debugf("[%s][%s]\n", r.FormValue("subId"), r.FormValue("uploadfile"))
		jklog.L().Debugf("%v, save to file [%s]\n", handler.Header,
			GetFileUploadPath().Path+"/"+handler.Filename)
		f, err := os.OpenFile(GetFileUploadPath().Path+"/"+handler.Filename,
			os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			jklog.L().Errorln(err)
			return
		}
		defer f.Close()
		n, err := io.Copy(f, file)
		if err != nil {
			jklog.L().Errorln("copy error: ", err)
		} else {
			jklog.L().Debugf("copy done len [%d]", n)
		}
		http.Redirect(w, r, "/upload", http.StatusFound)
	}
}
