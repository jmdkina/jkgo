package jkbase

import (
	"errors"
	"io"
	"io/ioutil"
	"jk/jklog"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

var NetTypeBase = 1

// go 实现类似 C++ 的虚拟继承，是一种假像，至少使用上是
// 通过一种特殊的方法，将子类的对像以接品的形式交给父类
// 父类再调用相同的方法，其实就是调用子类实现的方法，而非
// 父类的。但是这种方法也比使用回调稍微好一些，也不是很好理解
// 不如c++的虚函数来的简单，明了，是对语言的变相使用
type JKSuperBase interface {
	HandleMsg(conn net.Conn, data string) error
}

type JKNetBase struct {
	address string
	port    int
	nettype int

	// As client dial to
	conn     net.Conn
	recvdata string

	// As server listen
	listener  net.Listener
	items     map[string](JKNetBaseItem)
	HandleMsg func(conn net.Conn, data string) error
}

type JKNetBaseItem struct {
	conn       net.Conn
	remoteaddr string
	data       string
}

func (nb *JKNetBase) GetAddress() (string, int) {
	return nb.address, nb.port
}

func (nb *JKNetBase) GetNettype() int {
	return nb.nettype
}

// Set addr and port with nettype
// But listen or dial depends on what your operation next
func (nb *JKNetBase) New(addr string, port int, nettype int) error {
	nb.address = addr
	nb.port = port
	nb.nettype = nettype
	nb.items = make(map[string]JKNetBaseItem)
	return nil
}

func (nb *JKNetBase) Listen() error {
	if nb.conn != nil {
		return errors.New("Has used for dial")
	}
	var err error
	nb.listener, err = net.Listen("tcp", nb.address+":"+strconv.Itoa(nb.port))
	if err != nil {
		return err
	}
	return nil
}

func (nb *JKNetBase) Dial() error {
	if nb.listener != nil {
		return errors.New("Has used for listen")
	}
	conn, err := net.Dial("tcp", nb.address+":"+strconv.Itoa(nb.port))
	if err != nil {
		return err
	}
	nb.conn = conn
	return nil
}

func (nb *JKNetBase) Send(data string) int {
	if nb.conn == nil {
		return -2
	}
	n, err := nb.conn.Write([]byte(data))
	if err != nil {
		return -1
	}
	return n
}

func (nb *JKNetBase) DoRecvCycle() error {
	for {
		jklog.L().Debugln("Start to accept ...")
		conn, err := nb.listener.Accept()
		if err != nil {
			jklog.L().Errorln("accept failed ", err)
			return err
		}
		jklog.L().Infof("accept one client %s\n", conn.RemoteAddr().String())
		go func() {
			for {
				rdata := make([]byte, 2<<15)
				n, err := conn.Read(rdata)
				if err == io.EOF {
					jklog.L().Errorln("IO EOF exit")
					break
				}
				if err != nil {
					jklog.L().Errorln("Read error ", err)
					break
				}

				jklog.L().Debugf("Recv data of length %d, from %s\n", n, conn.RemoteAddr().String())
				item := JKNetBaseItem{}
				item.conn = conn
				item.remoteaddr = conn.RemoteAddr().String()
				item.data = string(rdata[0:n])
				nb.items[item.remoteaddr] = item
				nb.HandleMsg(conn, item.data)
			}
		}()
	}
	return nil
}

func (nb *JKNetBase) RecvClient() ([]byte, error) {
	if nb.conn == nil {
		return nil, errors.New("not init")
	}
	rdata := make([]byte, 2<<15)
	n, err := nb.conn.Read(rdata)
	if err != nil {
		return nil, err
	}
	nb.recvdata = string(rdata[0:n])
	return rdata[0:n], nil
}

func CMHttpGet(rurl string, params url.Values) ([]byte, error) {
	var Url *url.URL
	Url, err := url.Parse(rurl)
	if err != nil {
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func CMHttpPost(rurl string, params url.Values) ([]byte, error) {
	resp, err := http.PostForm(rurl, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
