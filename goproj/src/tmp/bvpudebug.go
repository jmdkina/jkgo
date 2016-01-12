package main

import (
	"flag"
	"fmt"
	"io"
	"jk/jklog"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// connect server with tcp and send command with special format
// only support these command:
// pipe[xxxx] xxxx is the command will send to center control.
//

var (
	serverAddr = flag.String("serverAddr", "", "which server to connect")
	serverPort = flag.Int("serverPort", 23433, "which port to connect")
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

func receiverAndOutValue(conn net.Conn) {
	jklog.L().Infoln("Receiver data from [", conn.LocalAddr().String(), "]")

	for {
		recvData := make([]byte, 1024)
		n, err := conn.Read(recvData)
		if io.EOF == err {
			break
		}
		if err != nil {
			jklog.L().Errorln("read error : ", err)
			break
		}
		jklog.L().Infof("recv data [%s], len [%d]\n", recvData, n)
		break
	}
}

func receiverAndOutResponse(listen, cmd string) {
	jklog.L().Infoln("Receiver data from [", listen, "] of file ", cmd)

	lListen, err := net.Listen("tcp", listen)
	if err != nil {
		return
	}
	defer lListen.Close()

	c, err := lListen.Accept()
	if err != nil {
		jklog.L().Errorln("Accept error : ", err)
		return
	}

	for {
		recvdata := make([]byte, 1024)

		// go func() {
		n, err := c.Read(recvdata)
		if io.EOF == err {
			time.Sleep(10 * time.Microsecond)
			break
		}
		if err != nil {
			jklog.L().Errorln("read error : ", err)
			break
		}
		jklog.L().Infof("recv data [%s], len [%d]\n", recvdata, n)
	}
}

func receiverAndSaveToFile(listen, file string) {
	jklog.L().Infoln("Receiver data from [", listen, "] of file ", file)

	lListen, err := net.Listen("tcp", listen)
	if err != nil {
		return
	}
	defer lListen.Close()

	c, err := lListen.Accept()
	if err != nil {
		jklog.L().Errorln("Accept error : ", err)
		return
	}
	f, ferr := os.Create(file)
	if ferr != nil {

	}
	defer f.Close()

	for {
		jklog.L().Infoln("Start to recever data ...")
		recvdata := make([]byte, 1024)

		// go func() {
		n, err := c.Read(recvdata)
		if io.EOF == err {
			time.Sleep(10 * time.Microsecond)
			break
		}
		if err != nil {
			jklog.L().Errorln("read error : ", err)
			break
		}

		if ferr == nil {
			f.WriteString(string(recvdata[0:n]))
		} else {
			jklog.L().Infoln("get data [", n, "] from out ", string(recvdata))
		}
		// }()
	}
}

func parseOutName(data string) string {
	out := data

	outs := strings.Split(data, "[")
	if len(outs) > 1 {
		out = outs[1][0 : len(outs[1])-1]
	}
	i := strings.LastIndex(out, "/")
	if i > 0 {
		out = out[i+1 : len(out)]
	}

	return out
}

func main() {
	flag.Parse()

	jklog.L().Infoln("Start to connect to ", *serverAddr, ", port ", *serverPort)

	connAddr := *serverAddr + ":" + strconv.Itoa(*serverPort)
	conn, err := net.Dial("tcp", connAddr)
	if err != nil {
		jklog.L().Errorln("fail to dail ", connAddr)
		return
	}

	for {
		fmt.Println("Give one command to deal with")
		data := scanLine()

		// data := "pipe[0 xx:misc_ctrl 0 debug aa level=2]"
		if string(data) == "exit" {
			break
		}

		n, err := conn.Write([]byte(data))
		if err != nil {
			jklog.L().Errorln("Write error ")
			continue
		}
		jklog.L().Infoln("Send success ", n, " of data ", string(data))

		if strings.HasPrefix(data, "file") {
			receiverAndSaveToFile(conn.LocalAddr().String(), parseOutName(data))
		} else if strings.HasPrefix(data, "cmd") {
			receiverAndOutValue(conn)
			// receiverAndOutResponse(conn.LocalAddr().String(), parseOutName(data))
		}

		time.Sleep(400 * time.Microsecond)
	}

	jklog.L().Infoln("Exit ...")
	return
}
