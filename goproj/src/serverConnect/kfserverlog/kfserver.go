package main

import (
	"jk/jklog"
	sv "jk/jknewserver"
	. "jk/jkprotocol"
	// "os"
	. "jk/jkcommon"
	 "time"
	"flag"
	daemon "github.com/tyranron/daemonigo"
)

const ()

type KFServer struct {
	handle     *sv.JKNewServer
	proc       sv.JKServerProcess
	onlineCnts int64 // for how many item online concurrent
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
			ret := false
			if len(filename) > 0 && len(fdata) > 0 {
				ret = JKSaveFileData(p.ID(), filename, fdata)
			}
			jklog.L().Debugln("filename: ", filename, ", data length: ", len(fdata))
			if ret {
				retstr := p.GenerateResponseOK()
				item.Conn.Write([]byte(retstr))
			} else {
				retstr := p.GenerateResponseFail()
				item.Conn.Write([]byte(retstr))
			}
		} else {
			jklog.Lfile().Errorln("Command not support")
			retstr := p.GenerateResponseFail()
			item.Conn.Write([]byte(retstr))
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
		if s.onlineCnts > 500 {
			time.Sleep(time.Millisecond*1000)
			continue
		}
		c := s.handle.Accept(&s.proc)
		if c == nil {
			jklog.Lfile().Errorln("accept failed.")
			return false
		}

		s.onlineCnts = s.onlineCnts+1

		item := &sv.JKServerProcessItem{}
		item.Conn = c
		item.ReadDone = make(chan bool)

		jklog.Lfile().Debugf("a new item connect in [%s] now counts [%d].\n", c.RemoteAddr().String(), s.onlineCnts)
		go s.dealResponse(s.proc, item)
		go func() bool {
			for {
				_, err := s.handle.Read(&s.proc, item)
				if err != nil {
					s.onlineCnts = s.onlineCnts -1
					jklog.Lfile().Warnf("closed for read error [%s] now online counts [%d].\n", item.Conn.RemoteAddr().String(), s.onlineCnts)
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

var (
	signal     = flag.String("action", "start", "{quit|stop|reload}")
	background = flag.Bool("d", false, "Background run")
)

func main() {
	flag.Parse()

	jklog.InitLog("/tmp/kfserver.log")
	jklog.L().Infoln("Program start ...")

	s := &KFServer{}

	// Daemonizing echo server application.
	if *background {
		jklog.L().Infoln("background run now.")
		switch isDaemon, err := daemon.Daemonize(*signal); {
		case !isDaemon:
			return
		case err != nil:
			jklog.L().Errorln("daemon start failed : ", err.Error())
		}
	}

	s.startServer()
}
