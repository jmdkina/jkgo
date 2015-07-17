package jkserver

import (
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

const (
	jk_listen_port = 10031
	jk_send_port   = 10041
)

const (
	jk_send_data_len = 312 // Force transfer len
)

func JK_StartServer() error {
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(jk_listen_port))
	if err != nil {
		return err
	}

	waitresp, err := net.ListenUDP("udp", service)

	if err != nil {
		return err
	}

	go func() {
		for {
			buf := make([]byte, jk_send_data_len)
			n, _, err := waitresp.ReadFromUDP(buf)
			if err != nil {
				jklog.L().Errorln("read from udp ", err)
				break
			}
			if n == 0 {
				time.Sleep(500 * time.Microsecond)
			} else {
				go func() {
					header, str := jk_receive_command(buf[0:n])
					jklog.L().Infoln("---remote addr : ", waitresp.RemoteAddr())
					jk_send_data_udp("255.255.255.255", jk_send_port, str, header)
				}()
			}
		}
		jklog.L().Errorln("shouldn't down here")
	}()
	jklog.L().Infoln("JKServer has stared in ", strconv.Itoa(jk_listen_port))

	return nil
}

func jk_send_data_udp(ip string, port int, data []byte, header *jkCommandHeader) (int, error) {
	sendsvc, err := net.ResolveUDPAddr("udp4", ip+":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("send addr resove err==>", err)
		return 0, err
	}
	sendto, err := net.DialUDP("udp", nil, sendsvc)
	if err != nil {
		jklog.L().Errorln("dial udp addr ==>", err)
		return 0, err
	}
	lsend := 0
	lremain := len(data)
	/// every time send jk_send_data_len(512) data
	for {
		if lremain <= 0 {
			break
		}
		if lremain >= jk_send_data_len {
			header.set_job_continue(JK_SEND_CONTINUE)
			header.set_data_length(jk_send_data_len)
			send_data := []byte(header.ToString() + "\r\n")
			send_data = append(send_data, data[lsend:lsend+jk_send_data_len]...)
			sendto.Write(send_data)
			lremain -= jk_send_data_len
			lsend += jk_send_data_len
		} else {
			header.set_job_continue(JK_SEND_LAST)
			header.set_data_length(lremain)
			send_data := []byte(header.ToString() + "\r\n")
			send_data = append(send_data, data[lsend:lremain+lsend]...)
			sendto.Write(send_data)
			lremain = 0
			lsend += lremain
		}
	}
	sendto.Close()

	return lsend, nil
}
