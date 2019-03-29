//
// Author: jmdvirus@roamter.com
//
// System operation
//
package jksys

import (
	"bytes"
	"errors"
	"jk/jklog"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	//	"fmt"
)

var Localaddr string
var LocalMac string

// Get CPU info
type ProcessCPU struct {
	pid int
	cpu float64
}

type KFSystemInfo struct {
	Kernel   string
	Core     string
	Platform string
	Hostname string
	CPUs     int

	OSName   string
	IPAddr   string
	Mac      string
	TotalRam uint64
	FreeRam  uint64

	TotalDisk uint64
	FreeDisk  uint64

	CPUUsage float64
}

func KFLocalMac() (string, error) {
	if len(LocalMac) > 0 {
		return LocalMac, nil
	}
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, inter := range interfaces {
		if (inter.Flags & net.FlagLoopback) == net.FlagLoopback {
			continue
		}
		macb := inter.HardwareAddr.String()
		mac := strings.Replace(macb, ":", "", -1)
		LocalMac = mac
		return mac, nil
	}
	return "", errors.New("No mac")
}

func KFIPAddressName(inter string) (string, error) {
	if len(Localaddr) > 0 {
		return Localaddr, nil
	}
	ifi, err := net.InterfaceByName(inter)
	if err != nil {
		return "", err
	}
	addrs, err := ifi.Addrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				Localaddr = ipnet.IP.String()
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("not find any address")
}

// Return ip address
func KFIPAddress() string {
	if len(Localaddr) > 0 {
		return Localaddr
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		jklog.L().Errorln("Can't get ip address : ", err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// jklog.L().Infoln("Get ip address: ", ipnet.IP.String())
				Localaddr = ipnet.IP.String()
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Return mac address with upper chars of the first not lo one.
func KFMacAddress() string {
	addrs, err := net.Interfaces()
	if err != nil {
		jklog.L().Error("Can't get mac address: ", err)
		return ""
	}
	for _, inter := range addrs {
		if strings.Contains(inter.Name, "lo") {
			continue
		}
		str := inter.HardwareAddr.String()
		realstr := strings.Replace(str, ":", "", -1)
		realstr = strings.ToUpper(realstr)
		// jklog.L().Infoln(realstr)
		return realstr
	}
	return ""
}

/**
 * Execute command with args
 */
func system(command string, args []string) bool {

	cmd := exec.Command(command, args[0])
	if len(args) == 2 {
		cmd = exec.Command(command, args[0], args[1])
	}
	jklog.L().Debugln(cmd.Args)
	// cmd.Env = os.Environ()
	// var out bytes.Buffer
	cmd.Stderr = os.Stderr
	// cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		jklog.L().Errorln("execute command failed. ", err)
		return false
	}
	return true
}

func systemargs(command string, args string) bool {
	cmd := exec.Command(command, strings.Split(args, " ")...)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		jklog.L().Errorln("execute command failed. ", err)
		return false
	}
	return true
}

func systemdata(command string, args string) string {
	cmd := exec.Command(command, strings.Split(args, " ")...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		jklog.L().Errorln("execute failed: ", err)
		return ""
	}
	return out.String()
}

// copy back of shadowsocks.json and restart shadowsocks.
func KFServerConfig(str string) error {
	// session := sh.NewSession()

	// first backup the file.
	prefix := "/etc"
	confFile := prefix + "/shadowsocks.json"
	// session.ShowCMD = true
	args := []string{prefix + "/shadowsocks.json", prefix + "/shadowsocks.json.back"}
	system("cp", args)
	// session.Command("cp", prefix+"/shadowsocks.json", prefix+"/shadowsocks.json.back").Run()

	f, err := os.OpenFile(confFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		jklog.L().Errorln("failed to open file: ", err)
		return err
	}
	defer f.Close()

	// jklog.L().Debugln("will write string: ", string(str))
	_, err = f.WriteString(str)
	if err != nil {
		jklog.L().Errorln("Write to shadowsocks failed.")
		return err
	} else {
		// session.Command("/etc/init.d/shadowsocks", "restart").Run()
		args = []string{"restart"}
		system("/etc/init.d/shadowsocks", args)
	}
	jklog.L().Infoln("System call done.")

	return nil
}

func KFReadAllData(filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		jklog.L().Errorln("open file ", filename, " failed.", err)
		return nil
	}
	defer f.Close()

	data := make([]byte, 4096<<2)
	n, err := f.Read(data)
	if err != nil {
		jklog.L().Errorln("read from file ", filename, " failed. ", err)
		return nil
	}
	jklog.L().Debugln("read data out len: ", n)

	return data[0:n]
}

// Reset firewall with @ports
// Set Router mode
// sysctl net.ipv4.ip_forward=1
// iptables -A FORWARD -d 133.130.97.109 -i eth1 -p tcp -m tcp --dport 8989 -j ACCEPT
// iptables -A FORWARD -d 133.130.97.109 -i eth1 -p udp -m udp --dport 8989 -j ACCEPT
// iptables -t nat -A PREROUTING -p tcp -m tcp --dport 8989 -j DNAT --to-destination 133.130.97.109:8989
// iptables -t nat -A PREROUTING -p udp -m udp --dport 8989 -j DNAT --to-destination 133.130.97.109:8989
// iptables -t nat -A POSTROUTING -o eth1 -j MASQUERADE
//
// dstip: forward to
// srcport: data from the port
// dstport: forward to port
// interf: check data for the interface (eth0,eth1)
//
func KFSetFireWallRouter(ports []string, dstip string, srcport, dstport int, interf string) bool {
	args := []string{"-F"}
	// clean firewall
	system("/sbin/iptables", args)
	jklog.Lfile().Infoln("clear iptables ...")

	nargs := []string{"net.ipv4.ip_forward=1"}
	system("sysctl", nargs)
	jklog.Lfile().Infoln("sysclt ip_forward to 1...")
	for _, v := range ports {
		strargs := "-A FORWARD -i " + interf + " -p tcp -m tcp --dport " + v + " -j ACCEPT"
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-A FORWARD -i " + interf + " -p udp -m udp --dport " + v + " -j ACCEPT"
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-A FORWARD -i " + interf + " -p tcp -m tcp --sport " + v + " -j ACCEPT"
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-A FORWARD -i " + interf + " -p udp -m udp --sport " + v + " -j ACCEPT"
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-t nat -A PREROUTING -p tcp -m tcp --dport " + v + " -j DNAT --to-destination " + dstip + ":" + v
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-t nat -A PREROUTING -p udp -m udp --dport " + v + " -j DNAT --to-destination " + dstip + ":" + v
		jklog.Lfile().Debugf("port: %s, args: [%s]\n", v, strargs)
		systemargs("iptables", strargs)

		strargs = "-t nat -A POSTROUTING -p tcp -m tcp --dport " + v + " -j MASQUERADE"
		systemargs("iptables", strargs)

		strargs = "-t nat -A POSTROUTING -p udp -m udp --dport " + v + " -j MASQUERADE"
		systemargs("iptables", strargs)

		strargs = "-t nat -A POSTROUTING -p tcp -m tcp --sport " + v + " -j MASQUERADE"
		systemargs("iptables", strargs)

		strargs = "-t nat -A POSTROUTING -p udp -m udp --sport " + v + " -j MASQUERADE"
		systemargs("iptables", strargs)
	}

	if len(ports) > 0 {

		cmd := exec.Command("iptables-save")
		var out bytes.Buffer
		cmd.Stderr = os.Stderr
		cmd.Stdout = &out
		err := cmd.Run()
		ret := true
		if err != nil {
			jklog.L().Errorln("error: ", err)
			ret = false
		}
		// jklog.L().Debugln("stdout: ", out.String())
		filesave := "/etc/sysconfig/iptables"
		f, err := os.OpenFile(filesave, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			jklog.L().Errorln("failed open file: ", err)
			ret = false
		}

		if len(out.String()) > 0 {
			n, err := f.WriteString(out.String())
			if err != nil {
				jklog.L().Errorln("write failed : ", err)
				ret = false
			}
			jklog.L().Debugln("write to file length: ", n)
			jklog.Lfile().Debugln("Write iptables config: ", out.String()[len(out.String())-50:len(out.String())])
		}
		defer f.Close()

		if ret {
			args = []string{"restart"}
			system("/etc/init.d/iptables", args)
		}
	}
	return true
}

// Reset firewall with @ports
func KFSetFireWall(ports []string) bool {
	args := []string{"-F"}
	// clean firewall
	system("/sbin/iptables", args)
	for _, v := range ports {
		strargs := "-A INPUT -p tcp --dport " + v + " -j ACCEPT"
		systemargs("iptables", strargs)

		strargs = "-A OUTPUT -p tcp --sport " + v + " -j ACCEPT"
		systemargs("iptables", strargs)

		strargs = "-A INPUT -p udp --dport " + v + " -j ACCEPT"
		systemargs("iptables", strargs)

		strargs = "-A OUTPUT -p udp --sport " + v + " -j ACCEPT"
		systemargs("iptables", strargs)
	}
	if len(ports) > 0 {

		cmd := exec.Command("iptables-save")
		var out bytes.Buffer
		cmd.Stderr = os.Stderr
		cmd.Stdout = &out
		err := cmd.Run()
		ret := true
		if err != nil {
			jklog.L().Errorln("error: ", err)
			ret = false
		}
		// jklog.L().Debugln("stdout: ", out.String())
		filesave := "/etc/sysconfig/iptables"
		f, err := os.OpenFile(filesave, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
		if err != nil {
			jklog.L().Errorln("failed open file: ", err)
			ret = false
		}

		if len(out.String()) > 0 {
			n, err := f.WriteString(out.String())
			if err != nil {
				jklog.L().Errorln("write failed : ", err)
				ret = false
			}
			jklog.L().Debugln("write to file length: ", n)
			jklog.Lfile().Debugln("Write iptables config: ", out.String()[len(out.String())-50:len(out.String())])
		}
		defer f.Close()

		if ret {
			args = []string{"restart"}
			system("/etc/init.d/iptables", args)
		}
	}
	return true
}

func KFStopFirewall() {
	args := []string{"stop"}
	system("/etc/init.d/iptables", args)
}

// first return is input traffic, second return is output traffic
func KFParseTrafficFromString(trafficstr, port string) (int64, int64) {
	// jklog.Lfile().Debugf("The traffic: [[   %s   ]\n", trafficstr)
	lines := strings.Split(trafficstr, "\n")

	var traffic []string
	for _, k := range lines {
		if strings.Contains(k, port) {
			k = strings.TrimSpace(k)
			k = strings.Trim(k, " ")
			items := strings.Split(k, " ")
			i := 0
			for _, d := range items {
				if len(d) > 0 {
					if i == 1 {
						traffic = append(traffic, d)
						break
					}
					i = i + 1
					continue
				}
				// jklog.L().Infoln("the alvue ? ", d)
			}

			// traffic = append(traffic, items[1])

			// jklog.L().Infoln("I need this line: ", k)
		}
	}

	var inputTraffic int64
	var outputTraffic int64
	if len(traffic) >= 4 {

		v1, _ := strconv.Atoi(traffic[0])
		v2, _ := strconv.Atoi(traffic[1])
		v3, _ := strconv.Atoi(traffic[2])
		v4, _ := strconv.Atoi(traffic[3])
		inputTraffic = int64(v3 + v4)
		outputTraffic = int64(v1 + v2)
		// jklog.L().Debugf("%d %d %d %d", v1, v2, v3, v4)
		// jklog.L().Debugf("%s: %d, %d", port, inputTraffic, outputTraffic)
	} else if len(traffic) == 2 {
		v1, _ := strconv.Atoi(traffic[0])
		v2, _ := strconv.Atoi(traffic[1])
		inputTraffic = int64(v1)
		outputTraffic = int64(v2)
	}

	return inputTraffic, outputTraffic
}

func KFPortTrafficRouter(port string) (int64, int64) {
	str := "-v -n -x -L FORWARD"
	data := systemdata("iptables", str)
	// jklog.L().Debugln("out data : ", data)
	return KFParseTrafficFromString(data, port)
}

// Caculate the traffic.
func KFPortTraffic(port string) (int64, int64) {
	str := "-v -n -x -L -t filter"
	data := systemdata("iptables", str)
	// jklog.L().Debugln("out data : ", data)
	return KFParseTrafficFromString(data, port)
}
