package jkserver

import (
	// "bytes"
	// "encoding/binary"
	"jk/jklog"
	"strconv"
	"strings"
)

type jkCommandHeader struct {
	id           string
	majorVersion int
	minorVersion int
	sequence     int
	job          string
	command      string
	length       int
	eachcommand  map[int]string
	eachjob      map[int]string
	data         []byte
}

/*
 * JK\r\n
 * Version:0.1\r\n
 * Sequence:2\r\n
 * Job:Get/Set/Report,Request/Response,Last/Continue\r\n
 * Command:BVWork,Ethernet,List\r\n
 * Length:18\r\n
 */

/*
 * bit
 * |--------|--------|--------|---------|
 * |        JK                          |
 * | 0    1 |0 0 0 0  |
 * |  1     |    1   |      2           |
 * |      id                            |
 * | len             |                  |
 * 4bytes: JK
 *
 * 5bits: major version(0-63)
 * 3bits: minor version(1-11)
 * 2bits: type(0 Get, 1 Set, 2 report)
 * 1bits: 0 send, 1 response
 * 1bits: 0 the last, 1 will continue
 * 4bits: reserved, set to 0
 *
 * 1bytes: major type(command type)
 * 1bytes: minor type(command type)
 * 2bytes: min type(command type)
 *
 * 4bytes: id (from 1 to 2^32-1 and return to 1 if over)
 * 2bytes: len

 	header := make([]byte, 16)

	header[0] = 'J'
	header[1] = 'K'
	header[4] |= (MAJOR_VER << 4) // major version num (4-7)
	header[4] |= (MINOR_VER << 0) // minor version num

	header[5] |= conti << 4 // This is continue, don't set if it is last
	header[5] |= resp << 5  // This is response, don't set if it is send
	header[5] |= cType << 6 // This is Set, don't set if it is get
*/

const (
	jkid               = "JK"
	major_ver          = 0
	minor_ver          = 1
	jk_header_version  = "Version"
	jk_header_job      = "Job"
	jk_header_command  = "Command"
	jk_header_length   = "Length"
	jk_header_sequence = "Sequence"
)

const (
	JK_TYPE_GET = iota
	JK_TYPE_SET
	JK_TYPE_REPORT

	JK_SEND_DOING    = 10
	JK_SEND_RESPONSE = iota
	JK_SEND_FAIL

	JK_SEND_LAST     = 20
	JK_SEND_CONTINUE = iota
)

const (
	JK_BVWORK = iota

	JK_BVETHERNET = 100

	JK_DEVICELIST = 10000
)

const (
	jk_type_get_name      = "Get"
	jk_type_set_name      = "Set"
	jk_type_report_name   = "Report"
	jk_send_doing_name    = "Request"
	jk_send_response_name = "Response"
	jk_send_continue_name = "Continue"
	jk_send_last_name     = "Last"
	jk_send_fail_name     = "Fail"

	jk_cmd_bvwork     = "JK_BVWork"
	jk_cmd_bvethernet = "JK_BVEthernet"
	jk_cmd_devicelist = "JK_DeviceList"
)

var jk_command map[int]string
var jk_job_name map[int]string

func init() {
	jk_command = map[int]string{
		JK_BVWORK:     jk_cmd_bvwork,
		JK_BVETHERNET: jk_cmd_bvethernet,
		JK_DEVICELIST: jk_cmd_devicelist,
	}
	jk_job_name = map[int]string{
		JK_TYPE_GET:      jk_type_get_name,
		JK_TYPE_SET:      jk_type_set_name,
		JK_TYPE_REPORT:   jk_type_report_name,
		JK_SEND_DOING:    jk_send_doing_name,
		JK_SEND_RESPONSE: jk_send_response_name,
		JK_SEND_LAST:     jk_send_last_name,
		JK_SEND_CONTINUE: jk_send_continue_name,
		JK_SEND_FAIL:     jk_send_fail_name,
	}
}

// Give me a name, give you a type int
// like jk_job_name(revert)
func (header *jkCommandHeader) name_type(name string) int {
	switch name {
	case jk_type_get_name:
		return JK_TYPE_GET
	case jk_type_set_name:
		return JK_TYPE_SET
	case jk_send_doing_name:
		return JK_SEND_DOING
	case jk_send_response_name:
		return JK_SEND_RESPONSE
	case jk_send_last_name:
		return JK_SEND_LAST
	case jk_send_continue_name:
		return JK_SEND_CONTINUE
	case jk_cmd_bvwork:
		return JK_BVWORK
	case jk_cmd_bvethernet:
		return JK_BVETHERNET
	case jk_cmd_devicelist:
		return JK_DEVICELIST
	case jk_send_fail_name:
		return JK_SEND_FAIL
	default:
		return -1
	}
	return -1
}

var seq = 0

func NewJKCommandHeader() *jkCommandHeader {
	jk_cmd_header := jkCommandHeader{
		id: jkid,
	}
	jk_cmd_header.eachjob = make(map[int]string)
	jk_cmd_header.eachcommand = make(map[int]string)
	return &jk_cmd_header
}

func (header *jkCommandHeader) GenJKCommandHeaderOneResponse(fCmd, sCmd, tCmd int) {
	header.set_job_info(JK_TYPE_GET, JK_SEND_RESPONSE, JK_SEND_LAST)
	header.set_command_info(fCmd, sCmd, tCmd)
	header.set_data_length(0)
}

func (header *jkCommandHeader) GenJKCommandHeaderOneSend(fCmd, sCmd, tCmd int) {
	header.set_job_info(JK_TYPE_GET, JK_SEND_DOING, JK_SEND_LAST)
	header.set_command_info(fCmd, sCmd, tCmd)
	header.set_data_length(0)
}

func (header *jkCommandHeader) Join(data string) string {
	header.length = len(data)
	header.data = []byte(data)
	str := header.ToString()
	str += "\r\n" + data
	return str
}

// Parse success mean the header is valid
func jk_parse_command_header(str string) *jkCommandHeader {
	sp := strings.Split(str, "\r\n")
	if len(sp) < 6 || sp[0] != jkid {
		jklog.L().Errorln("wrong args id")
		return nil
	}
	var err error
	jkcommandheader := NewJKCommandHeader()

	// parse version
	if strings.HasPrefix(sp[1], jk_header_version) {
		wsp := strings.Split(sp[1], ":")
		if len(wsp) != 2 {
			jklog.L().Errorln("parse args error")
			return nil
		}
		nsp := strings.Split(wsp[1], ".")
		if len(nsp) != 2 {
			jklog.L().Errorln("parse version error")
			return nil
		}
		jkcommandheader.majorVersion, err = strconv.Atoi(nsp[0])
		if err != nil {
			return nil
		}
		jkcommandheader.minorVersion, err = strconv.Atoi(nsp[1])
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	// parse sequence
	if strings.HasPrefix(sp[2], jk_header_sequence) {
		wsp := strings.Split(sp[2], ":")
		if len(wsp) != 2 {
			return nil
		}
		jkcommandheader.sequence, err = strconv.Atoi(wsp[1])
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	// parse jobs
	wsp := strings.Split(sp[3], ":")
	if len(wsp) != 2 || wsp[0] != jk_header_job {
		jklog.L().Errorln("parse job err ")
		return nil
	}
	nsp := strings.Split(wsp[1], ",")
	if len(nsp) != 3 {
		jklog.L().Errorln("parse job err")
		return nil
	}
	for i, k := range nsp {
		jkcommandheader.eachjob[i] = k
	}

	// parse command
	wsp = strings.Split(sp[4], ":")
	if len(wsp) != 2 || wsp[0] != jk_header_command {
		jklog.L().Errorln("parse command err")
		return nil
	}
	nsp = strings.Split(wsp[1], ",")
	if len(nsp) != 3 {
		jklog.L().Errorln("parse command err")
		return nil
	}
	for i, k := range nsp {
		jkcommandheader.eachcommand[i] = k
	}

	// parse length
	wsp = strings.Split(sp[5], ":")
	if len(wsp) != 2 || wsp[0] != jk_header_length {
		jklog.L().Errorln("parse length err")
		return nil
	}
	jkcommandheader.length, err = strconv.Atoi(wsp[1])
	if err != nil {
		jklog.L().Errorln("parse length err", err)
		return nil
	}

	if len(sp) > 6 {
		jkcommandheader.data = make([]byte, jkcommandheader.length)
		jkcommandheader.data = []byte(sp[6])
	}

	return jkcommandheader
}

/**
 * cType: command type (see JK_TYPE_*)
 * resp: 0 send, 1 response
 * conti: 0 the last package, 1 will continue
 */
func (header *jkCommandHeader) set_common_info() {
	header.majorVersion = major_ver
	header.minorVersion = minor_ver
	header.sequence = seq
}

func (head *jkCommandHeader) set_job_info(cType int, resp int, conti int) {
	head.set_common_info()
	var header string
	header += jk_job_name[cType]

	if head.eachjob == nil {
		head.eachjob = make(map[int]string)
	}
	head.eachjob[0] = jk_job_name[cType]

	header += ","
	header += jk_job_name[resp]
	head.eachjob[1] = jk_job_name[resp]

	header += ","
	header += jk_job_name[conti]
	head.eachjob[2] = jk_job_name[conti]

	head.job = header
}

func (header *jkCommandHeader) set_command_info(fCmd, sCmd, tCmd int) {
	if header.eachcommand == nil {
		header.eachcommand = make(map[int]string)
	}
	header.eachcommand[0] = jk_command[fCmd]
	header.eachcommand[1] = jk_command[sCmd]
	header.eachcommand[2] = jk_command[tCmd]
	header.command = jk_command[fCmd] + "," + jk_command[sCmd] + "," + jk_command[tCmd]
}

func (header *jkCommandHeader) set_data_length(lendata int) {
	header.length = lendata
}

func (header *jkCommandHeader) set_job_continue(cont int) {
	header.eachjob[2] = jk_job_name[cont]
}

func (header *jkCommandHeader) ToString() string {
	header_str := header.id + "\r\n" +
		"Version:" + strconv.Itoa(header.majorVersion) + "." +
		strconv.Itoa(header.minorVersion) + "\r\n" +
		"Sequence:" + strconv.Itoa(header.sequence) + "\r\n" +
		"Job:" + header.eachjob[0] + "," +
		header.eachjob[1] + "," + header.eachjob[2] + "\r\n" +
		"Command:" + header.eachcommand[0] + "," +
		header.eachcommand[1] + "," + header.eachcommand[2] + "\r\n" +
		"Length:" + strconv.Itoa(header.length)

	return header_str
}
