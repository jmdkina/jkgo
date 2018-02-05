package main

import (
    "net"
    "jk/jklog"
)

type RecvData struct {
    Path     string
    rdata    []byte
    conn     *net.UnixListener
    c        *net.UnixConn
}

func NewRecvData(path string) (*RecvData, error) {
    rd := &RecvData{
        Path: path,
        rdata: make([]byte, 40960),
    }
    u, err := net.ResolveUnixAddr("unix", path)
    if err != nil {
        jklog.L().Errorln("Resolve error : ", err)
        return nil, err
    }
    rd.conn, err = net.ListenUnix("unix", u)
    if err != nil {
        jklog.L().Errorln("listen unix error : ", err)
        return nil, err
    }
    return rd, nil
}

func (rd *RecvData) Accept() error {
    jklog.L().Debugf("accept from %s\n", rd.Path)
    c, err := rd.conn.AcceptUnix()
    if err != nil {
        jklog.L().Errorln("accept unix error: ", err)
        return err
    }
    rd.c = c
    jklog.L().Infof("accept one %s\n", rd.c.RemoteAddr().String())
    return nil
}

func (rd *RecvData) Read() ([]byte, error) {
    n, _ := rd.c.Read(rd.rdata)
    return rd.rdata[:n], nil
}

func (rd *RecvData) Release() {
    rd.conn.Close()
}
