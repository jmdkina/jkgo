package main

import (
	"flag"
	"fmt"
	"io"
	"jk/jklog"
	"net"
	"strconv"
	"strings"
	"time"
)

func scanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
}

var bexit = false

type RTKFDebug struct {
	usedPorts []int
}

type RTKFDebugItem struct {
	conn net.Conn
	data chan string
}

func usage() {
	fmt.Println()
	fmt.Println("  cmd[ps w]        -- execute command and response.")
	fmt.Println("  file[/etc/model] -- get file info")
	fmt.Println("  cmd[rtclose]     -- let client exit.")
	fmt.Println("  exitself         -- close self.")
	fmt.Println()
}

func listenLocalTcp(port int) {
	addrto, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("error reolve tcp addr ", err)
		return
	}

	listen, err := net.ListenTCP("tcp", addrto)
	if err != nil {
		jklog.L().Errorln("list local with tcp error ", err)
		return
	}

	jklog.L().Infoln("Start to listen with tcp of ", port)
	go func() {
		for {
			jklog.L().Infoln("start to accept from remote ...")
			from, err := listen.Accept()
			if err != nil {
				jklog.L().Errorln("accept error ", err)
				continue
			}
			for {
				fmt.Print("Connect Client " + from.RemoteAddr().String() + " Cmd : > ")
				senddata := scanLine()
				if strings.Contains(senddata, "help") {
					usage()
					continue
				}
				if len(senddata) <= 5 || senddata == "\r" {
					continue
				}
				from.Write([]byte(senddata))

				if strings.Contains(senddata, "exitself") {
					bexit = true
				}

				remainLen := -1
				totalLen := 0
				for {

					buf := make([]byte, 2<<12)
					n, err := from.Read(buf)
					if err == io.EOF {
						break
					}
					if err != nil {
						jklog.L().Infoln("read error ", err)
						break
					}
					if totalLen > 0 {
						remainLen = remainLen + n
					}

					value := string(buf[0:n])
					if strings.HasPrefix(value, "length") {
						val := strings.Split(value, ":")
						if len(val) > 1 {
							l, err := strconv.Atoi(val[1])
							if err == nil {
								totalLen = l
								remainLen = 0
							}
						}
					}

					jklog.L().Infoln("read out data ", n, " of \n", string(buf[0:n]))
					if totalLen == 0 {
						break
					}
					if remainLen >= totalLen {
						break
					}
				}
			}
		}
		jklog.L().Errorln("Can't be here")
	}()

}

var (
	serverPort = flag.Int("serverPort", 23433, "which port to connect")
)

func RTKFDebugConsult(port int) {
	addrto, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("error reolve tcp addr ", err)
		return
	}

	listen, err := net.ListenTCP("tcp", addrto)
	if err != nil {
		jklog.L().Errorln("list local with tcp error ", err)
		return
	}

	jklog.L().Infoln("Start to consult with tcp of ", port)
	go func() {
		for {
			jklog.L().Infoln("start to accept consult from remote ...")
			from, err := listen.Accept()
			if err != nil {
				jklog.L().Errorln("accept error ", err)
				continue
			}
			jklog.L().Infoln("Recieved a connected client.. start to read data...")
			for {
				buf := make([]byte, 2<<12)
				n, err := from.Read(buf)
				if err == io.EOF {
					break
				}
				if err != nil {
					jklog.L().Infoln("read error ", err)
					break
				}
				jklog.L().Debugln("Request result: ", buf[0:n])
				val := string(buf[0:n])
				if strings.HasPrefix(val, "requestport") {
					jklog.L().Infoln("Request success, give response port: ")
					from.Write([]byte("port:23434"))
					go listenLocalTcp(23434)
				}
			}
		}
	}()

}

func main() {

	flag.Parse()

	// RTKFDebugConsult(*serverPort)
	listenLocalTcp(*serverPort)

	for {
		if bexit == true {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
}
