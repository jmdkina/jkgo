package main

import (
    "flag"
    "time"
    "os"
    "jk/jklog"
)

var (
    unixpath = flag.String("path", "/tmp/av.unix", "unix path")
    logfile = flag.String("logfile", "/tmp/jkavdu.log", "jkavdu logfile")
    logconsole = flag.Bool("logconsole", false, "Log console print instead file")
)

func main() {
    flag.Parse()

    if !*logconsole {
        jklog.InitLog(*logfile)
    }

    os.Remove(*unixpath)
    rd,_ := NewRecvData(*unixpath)

    f, _ := os.OpenFile("/tmp/v.h264", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

    rd.Accept()
    for {
        rdata, _ := rd.Read()
        jklog.L().Debugln("Get data of len ", len(rdata))
        f.Write(rdata)
        time.Sleep(time.Millisecond * 40)
    }
    defer f.Close()
    rd.Release()
}

