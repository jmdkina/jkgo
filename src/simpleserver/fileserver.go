package simpleserver

import (
	"net/http"
	"strings"
	"jk/jklog"
)

type FileServerInfo struct {
	Path      string
	Status    int
}

var file_server_map map[string]*FileServerInfo

func init() {
	file_server_map = make(map[string]*FileServerInfo, 10)
}

func GetFileServers() map[string]*FileServerInfo {
	return file_server_map
}

// Golang fileserver need:
// J:\\d
// You need add handler of the name last name "d"
// handler to the path before "d" , J:
// That will valid
func AddFileServer(path string) error {
	fs_t := file_server_map[path]
	if fs_t != nil {
		return nil
	}
	fs := &FileServerInfo{
		Path: path,
		Status: 1,
	}
	n := strings.LastIndex(path, "\\")
	prefix := path[:n]
	pathname := path[n+1:]
	http.Handle("/" + pathname + "/", http.FileServer(http.Dir(prefix)))
	jklog.L().Debugf("Add [%s] of [%s]\n", pathname, path)
	file_server_map[path] = fs
	return nil
}

func DelFileServer(path string) error {

	return nil
}
