package jkcenter

import (
	"io"
	"jk/jklog"
	"jk/jkprotocol"
	"net"
	"strconv"
)

type CenterControl struct {
	lis   net.Listener
	proto *jkprotocol.JKProtocolWrap
	lists []CommandItem
}

type CommandItem struct {
	conn net.Conn
	resp string
}

// nettype, 0 tcp, 1 udp
func InitCenter(laddr string, lport int, nettype int) (*CenterControl, error) {
	cc := &CenterControl{}

	lis, err := net.Listen("tcp", laddr+":"+strconv.Itoa(lport))
	if err != nil {
		return nil, err
	}
	jklog.L().Debugf("Now list [%s]\n", laddr + ":" + strconv.Itoa(lport))
	cc.lis = lis
	cc.proto, err = jkprotocol.NewJKProtocolWrap(jkprotocol.JK_PROTOCOL_VERSION_5)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (cc *CenterControl) Recv() error {
	for {
		conn, err := cc.lis.Accept()
		if err != nil {
			return err
		}
		jklog.L().Debugf("accept one connection from [%s]\n", conn.RemoteAddr().String())
		go func() {
			for {
				// Recv data
				rdata := make([]byte, 2 << 15)
				n, err := conn.Read(rdata)
				if err == io.EOF {
					jklog.L().Errorf("has EOF")
					break
				}

				// Parse
				str, err := cc.proto.Parse(string(rdata[:n]))
				if err != nil {
					jklog.L().Errorln("Parse error ", err)
					continue
				}
				jklog.L().Debugln("lllll ", str)

				// transfer to other depends on cmd

				item := CommandItem{}
				item.conn = conn
				item.resp = str
				cc.lists = append(cc.lists, item)
				jklog.L().Infof("Got command of [ %d ]\n", cc.proto.CmdType)

				// Exit
				if cc.proto.CmdType == jkprotocol.JK_PROTOCOL_C_LEAVE {
					break
				}
			} // recv
		}()
	} // accept

	return nil
}
