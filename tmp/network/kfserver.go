package main

import (
	"jk/jklog"
	sv "jk/jknewserver"
)

type KFServer struct {
	handle *sv.JKNewServer
	proc   sv.JKServerProcess
}

func (s *KFServer) dealResponse(proc sv.JKServerProcess, item *sv.JKServerProcessItem) {
	for {
		jklog.L().Debugln("wait response of read result.")
		ret := <-item.ReadDone

		if ret {
			jklog.L().Debugln("response of deal ", item.RemoteAddr)
			jklog.L().Debugln("data is : ", string(item.Data))
		} else {
			jklog.L().Debugln("read response failed ", item.RemoteAddr)
			break
		}
	}
}

func (s *KFServer) startServer() bool {
	ret := s.handle.Start(&s.proc)
	if !ret {
		jklog.L().Errorln("failed start server.")
		return false
	}

	for {
		jklog.L().Debugln("wait accept. ")
		c := s.handle.Accept(&s.proc)
		if c == nil {
			jklog.L().Errorln("accept failed.")
			return false
		}

		item := &sv.JKServerProcessItem{}
		item.Conn = c
		item.ReadDone = make(chan bool)
		go s.dealResponse(s.proc, item)
		go func() bool {
			for {
				_, err := s.handle.Read(&s.proc, item)
				if err != nil {
					item.ReadDone <- false
					return false
				}
			}
			return true
		}()
	}
	jklog.L().Errorln("Program return for failed start.")

	return true
}

func main() {
	s := &KFServer{}
	s.startServer()
}
