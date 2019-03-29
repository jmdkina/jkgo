package main

import (
	"jk/jklog"
	"net"
	"strconv"
	"time"
	"jk/jksip"
	"os"
)

var longdata = "I mean you are write, because if is not ok, it is must not like this, of course, not like that also, and this is just a test, if you don't beleive, I have no idea also, but you must beleive me, I'm write."


var start_av bool

// This file is use for test for udp
// one is listen udp, each receive from udp send one data out for test
// the one is send data with udp

func connect_out(addr string, port int) {
	jklog.L().Infoln("connect out with udp of ", addr, ":", port)
	addrto, err := net.ResolveUDPAddr("udp4", addr+":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("resolve udp addr failed")
		return
	}

	sendto, err := net.DialUDP("udp", nil, addrto)
	if err != nil {
		jklog.L().Errorln("dial udp failed ", err)
		return
	}
	defer sendto.Close()

	// outdata := "Here is for test"
	n, _ := sendto.Write([]byte(longdata))
	jklog.L().Infoln("send out data of len ", n)
}

func listenLocal(port int) {
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("resolve udp fail: ", err)
		return
	}

	waitresp, err := net.ListenUDP("udp", service)
	// defer waitresp.Close()

	jklog.L().Infoln("Start to listen ", port)
	has_invite := false
	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := waitresp.ReadFromUDP(buf[0:])
			if err != nil {
				jklog.L().Errorln("accept err :", err)
				break
			}
			if n == 0 {
				time.Sleep(500 * time.Millisecond)
			} else {
				//jklog.L().Infoln("received data ", n)
				str := string(buf[0:n])
				sc, _ := jksip.NewSipCtonrol(str)

				can_invite := true
				if sc.Command == "REGISTER" {
					jklog.L().Debugln("data: \n", str)
					jklog.L().Debugln("sip control value: ", sc.CallID, ", ", sc.Branch)
					resok, _ := sc.GenerateOK()

					jklog.L().Debugln(resok)
					_, err := waitresp.WriteToUDP([]byte(resok), addr)
					if err != nil {
						jklog.L().Errorln("write response fail: ", err)
						continue
					}
					jklog.L().Infoln("write response len ", n, " to ")
					can_invite = true
				} else {
					//jklog.L().Errorln("Other response , don't deal now.")
				}

				if sc.Command == "MESSAGE" {
					resmsg, _ := sc.GenerateKeepAlive()
					_, err := waitresp.WriteToUDP([]byte(resmsg), addr)
					if err != nil {
						jklog.L().Errorln("message response failed ", err)
					}
					continue
				}
				if can_invite && !has_invite {
					// Start to send invite
					resinvite, _ := sc.GenerateInvite(40004)
					_, err = waitresp.WriteToUDP([]byte(resinvite), addr)
					if err != nil {
						jklog.L().Errorln("send invite failed", err)
						continue
					}
					jklog.L().Infoln("has send invite")
					start_av = true
					//break


					has_invite = true
				}

				if sc.ResOK && sc.Command == "INVITE" {
					resack, _ := sc.GenerateACK()
					_, err = waitresp.WriteToUDP([]byte(resack), addr)
					if err != nil {
						jklog.L().Errorln("send ack falied ", err)
						continue
					}
					jklog.L().Debugln("Response ACK ok.")
				}
			}
		}
	}()
}

func waitAVData(port int) {
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("resolve udp fail: ", err)
		return
	}

	waitresp, err := net.ListenUDP("udp", service)
	// defer waitresp.Close()

	f, err := os.OpenFile("/tmp/28181.dat", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		jklog.L().Warnln("open file failed.", err)
	}

	jklog.L().Infoln("Start to listen for av data ", port)
	go func() {
		//ret := <- start_av
		for {
			if start_av {
				jklog.L().Infoln("OK, start to recevie video audio data.")
				for {

					buf := make([]byte, 1024)
					n, _, err := waitresp.ReadFromUDP(buf[0:])
					if err != nil {
						jklog.L().Errorln("accept err :", err)
						break
					}
					jklog.L().Infoln("recevied data of ", n)
                    _, err = f.Write(buf[0:n])
					if err != nil {
						jklog.L().Errorln("write failed ", err)
					}
				}
			}
			time.Sleep(time.Microsecond*500)
		}
		//jklog.L().Errorln("give a wrong value")
	}()
}

func main() {
	if_send_data_out := false

	listenLocal(5060)

	waitAVData(40004)

	time.Sleep(5000 * time.Millisecond)
	// if_send_data_out = true
	if if_send_data_out {
		connect_out("192.168.6.152", 10041)
	}
	for {
		time.Sleep(200 * time.Millisecond)
	}
}
