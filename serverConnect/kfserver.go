package main

import (
	"jk/jklog"
	sv "jk/jknewserver"
	. "jk/jkprotocol"
	// "os"
	. "jk/jkcommon"
	// "time"
)

const ()

type KFServer struct {
	handle *sv.JKNewServer
	proc   sv.JKServerProcess
}

func (s *KFServer) dealResponseCmd(data string, item *sv.JKServerProcessItem) {
	p := ParseJKProtocol(data)

	jklog.Lfile().Debugln("start parse jkprotocol down.")
	if p != nil {
		jklog.Lfile().Debugln("command: ", p.Command(), "subcommand: ", p.SubCommand())
		if p.Command() == JK_PROTOCOL_CMD_REGISTER {
			jklog.Lfile().Debugln("This is Register command.")
			retstr := p.GenerateResponseOK()
			item.Conn.Write([]byte(retstr))
		} else if p.Command() == JK_PROTOCOL_CMD_CONTROL && p.SubCommand() == JK_PROTOCOL_CMD_SAVEFILE {
			jklog.Lfile().Debugln("This is save file command")
			filename, fdata := p.ParseFilenameData()
			ret := JKSaveFileData(p.ID(), filename, fdata)
			jklog.L().Debugln("filename: ", filename, ", data length: ", len(fdata))
			if ret {
				retstr := p.GenerateResponseOK()
				item.Conn.Write([]byte(retstr))
			} else {
				retstr := p.GenerateResponseFail()
				item.Conn.Write([]byte(retstr))
			}
		}
	} else {
		jklog.L().Debugln("invalid data: ", data)
	}
	jklog.Lfile().Debugln("jkprotocol parse down.")
}

func (s *KFServer) dealResponse(proc sv.JKServerProcess, item *sv.JKServerProcessItem) {
	for {
		jklog.Lfile().Infoln("wait response of read result.", item.Id)
		ret := <-item.ReadDone

		/*
			now := time.Now().Unix()
			if item.TimeLast == 0 {
				item.TimeLast = now
			}
			if now-item.TimeLast > 5 {
				// 5 seconds
				continue
			}
			item.TimeLast = now
		*/
		if ret {
			jklog.Lfile().Debugln(item.Id, " : response of deal ", item.RemoteAddr)
			// jklog.L().Debugln("data is : ", string(item.Data))
			s.dealResponseCmd(string(item.Data), item)
		} else {
			jklog.Lfile().Errorln(item.Id, " : read response failed ", item.RemoteAddr)
			break
		}
	}
}

func (s *KFServer) startServer() bool {
	ret := s.handle.Start(&s.proc)
	if !ret {
		jklog.Lfile().Errorln("failed start server.")
		return false
	}

	for {
		jklog.L().Infoln("wait accept. ")
		jklog.Lfile().Infoln("wait accept. ")
		c := s.handle.Accept(&s.proc)
		if c == nil {
			jklog.Lfile().Errorln("accept failed.")
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
	jklog.Lfile().Errorln("Program return for failed start.")

	return true
}

func main() {

	jklog.InitLog("/tmp/kfserver.log")

	s := &KFServer{}
	s.startServer()
}
