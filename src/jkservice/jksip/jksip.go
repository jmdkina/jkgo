package jksip

import (
	"strings"
	"fmt"
)

type SipControl struct {
	Command string
	CallID  string
	Branch  string
	ResOK   bool
}

const (
	REGISTER_NAME = "REGISTER"
	INVITE_NAME = "INVITE"
	MESSAGE_NAME = "MESSAGE"
	VIA_NAME = "Via"
	CALLID_NAME = "Call-ID"
)

// parse

func NewSipCtonrol(data string) (*SipControl, error) {
	sc := SipControl{}
	if strings.HasPrefix(data, REGISTER_NAME) {
		sc.Command = REGISTER_NAME
	} else if strings.HasPrefix(data, INVITE_NAME) {
		sc.Command = INVITE_NAME
	} else if strings.HasPrefix(data, MESSAGE_NAME) {
		sc.Command = MESSAGE_NAME
	}
	strs := strings.Split(data, "\r\n")
	for _, str := range strs {
		if strings.HasPrefix(str, VIA_NAME) {
			// parse branch
			indx := strings.Index(str, "branch=")
			if indx > 0 {
				branch := str[indx+7:]
				sc.Branch = branch
			}
		}
		// find Call-ID
		if strings.HasPrefix(str, CALLID_NAME) {
			callids := strings.Split(str, ":")
			sc.CallID = callids[1]
		}

		// if resonseok
		if strings.Contains(str, "200 OK") {
			sc.ResOK = true
		}

		if sc.ResOK {
			// pares out command
			if strings.Contains(str, "CSeq") {
				if strings.Contains(str, INVITE_NAME) {
					sc.Command = INVITE_NAME
				} else if strings.Contains(str, REGISTER_NAME) {
					sc.Command = REGISTER_NAME
				}
			}
		}
	}
	return &sc, nil
}

// Generate response OK string.
func (sc *SipControl) GenerateOK() (string, error) {
	branch := fmt.Sprintf("Via: SIP/2.0/UDP 192.168.64.35:5062;rport=5062;branch=%s", sc.Branch)
	callid := fmt.Sprintf("Call-ID: %s", sc.CallID)

	ress := []string{"SIP/2.0 200 OK",
	branch,
	"From: <sip:34020000001320000050@3402000000>;tag=330156484",
	"To: <sip:34020000001320000050@3402000000>;tag=288315908",
	callid,
	"CSeq: 1 REGISTER",
	"User-Agent: Arges SIP-AS/1.0",
	"Date: 2016-10-26T09:40:29.050",
	"Expires: 86400",
	"Content-Length: 0", "\r\n"}
	return strings.Join(ress, "\r\n"), nil
}

func (sc *SipControl) GenerateInvite(port int) (string, error) {
	sdpstrs := []string{
		"v=0",
		"o=34020000002000000001 0 0 IN IP4 192.168.64.123",
		"s=Play",
		"c=IN IP4 192.168.64.123",
		"t=0 0",
		"m=video 40004 RTP/AVP 96",
		"a=recvonly",
		"a=rtpmap:96 PS/90000",
		"a=streamprofile:0",
		"y=0394939821",
		"",
	}
	sdpstr := strings.Join(sdpstrs, "\r\n")
	lengthstr := fmt.Sprintf("Content-Length:%d", len(sdpstr))
	ress := []string {
		"INVITE sip:34020000001320000001@192.168.64.35:5062 SIP/2.0",
		"Via: SIP/2.0/UDP 192.168.64.123:5060;rport;branch=z9hG4bK1323148440",
		"From: <sip:34020000002000000001@192.168.64.123:5060>;tag=1377356103",
		"To: <sip:34020000001320000001@192.168.64.35:5062>",
		"Call-ID: 20161025154933496295500@192.168.64.123",
		"CSeq: 20 INVITE",
		"Contact: <sip:34020000002000000001@192.168.64.123:5060>",
		"Max-Forwards: 70",
		"User-Agent: Arges SIP-AS/1.0",
		"Expires: 120",
		"Subject: 34020000001320000001:00000001,34020000002000000001:00000001",
		"Content-Type: application/sdp",
		lengthstr,
        "",
		sdpstr,
	}
	return strings.Join(ress, "\r\n"), nil
}

func (sc *SipControl) GenerateACK() (string, error) {
	ress := [] string {
		"ACK sip:34020000001320000050@192.168.64.35:5062 SIP/2.0",
		"Via: SIP/2.0/UDP 192.168.64.123:5060;rport;branch=z9hG4bK763181511",
		"From: <sip:34020000002000000001@192.168.64.123:5060>;tag=1377356103",
		"To: <sip:34020000001320000001@192.168.64.35:5062>;tag=383369377",
		"Call-ID: 20161025154933496295500@192.168.64.123",
		"CSeq: 20 ACK",
		"Contact: <sip:34020000002000000001@192.168.64.123:5060>",
		"Max-Forwards: 70",
		"User-Agent: Arges SIP-AS/1.0",
		"Content-Length: 0", "\r\n"}
	return strings.Join(ress, "\r\n"), nil
}

func (sc *SipControl) GenerateKeepAlive() (string, error ) {
	branch := fmt.Sprintf("Via: SIP/2.0/UDP 192.168.64.35:5062;rport=5062;branch=%s", sc.Branch)
	callid := fmt.Sprintf("Call-ID: %s", sc.CallID)
	//"Via: SIP/2.0/UDP 192.168.64.35:5062;rport=5062;branch=z9hG4bK1049979858",
	//"Call-ID: 1721120310",
	ress := [] string {
		"SIP/2.0 200 OK",
		branch,
		"From: <sip:34020000001320000050@3402000000>;tag=1521090264",
		"To: <sip:34020000001110000001@3402000000>;tag=1214878693",
		callid,
		"CSeq: 20 MESSAGE",
		"User-Agent: Arges SIP-AS/1.0",
		"Content-Length: 0",
	}
	return strings.Join(ress, "\r\n"), nil
}
