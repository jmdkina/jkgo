/*
 * Author: jmdvirus@roamter.com
 * System command test
 */

package jksys

import (
	"jk/jklog"
	"strings"
	"testing"
)

func TestIsProgramRunning(t *testing.T) {
	ret := SysProgramRunning("thunder")
	if ret {
	} else {
		t.Fatal("ShouldRunning")
	}
}

func TestKFServerConfig(t *testing.T) {

}

func TestKFLocalMac(t *testing.T) {
	mac, err := KFLocalMac()
	if err != nil {
		t.Fatal("err: ", err)
	}
	t.Fatal("str: ", mac)
	strings.ToUpper(mac)
}

func TestKFIPAddress(t *testing.T) {
	str, err := KFIPAddressName("en0")
	if err != nil {
		t.Logf("value is %v", err)
	}
	t.Logf("out string: %s", str)
	KFIPAddress()
}

func TestKFMacAddress(t *testing.T) {
	KFMacAddress()
}

func TestKFReadAllData(t *testing.T) {
	data := KFReadAllData("/etc/shadowsocks.json")
	jklog.L().Infoln("data: ", string(data))
}

/*
func TestKFSetFireWall(t *testing.T) {
	ports := []string{"10036", "10088", "10019"}
	ret := KFSetFireWall(ports)
	if ret != true {
		t.Fatal("error ", ret)
	}
}

*/

func TestKFParseTrafficFromString(t *testing.T) {
	data := "      1043   103888 ACCEPT     tcp  --  *      *       0.0.0.0/0            0.0.0.0/0           tcp dpt:10063" + "\n" +
		"       227    21011 ACCEPT     udp  --  *      *       0.0.0.0/0            0.0.0.0/0           udp dpt:10063" + "\n" +
		"       884   505401 ACCEPT     tcp  --  *      *       0.0.0.0/0            0.0.0.0/0           tcp spt:10063" + "\n" +
		"       227    33429 ACCEPT     udp  --  *      *       0.0.0.0/0            0.0.0.0/0           udp spt:10063"
	v1, v2 := KFParseTrafficFromString(data, "10063")
	if v2 != 103888+21012 || v1 != 505401+33429 {
		t.Fatalf("%d, %d, failed.", v1, v2)
	}
}

func TestKFParseTrafficRouterFromString(t *testing.T) {
	data := "   10952   686644 ACCEPT     tcp  --  eth1   *       0.0.0.0/0            159.203.31.82       tcp dpt:10001" + "\n" +
		"   3      231 ACCEPT     udp  --  eth1   *       0.0.0.0/0            159.203.31.82       udp dpt:10001"
	v1, v2 := KFParseTrafficFromString(data, "10001")
	t.Fatalf("%d , %d\n", v1, v2)
}

// Sysinfo Test
func TestKFMemInfo(t *testing.T) {
	mi := &KFSystemInfo{}
	total, free := mi.KFMemInfo()
	jklog.L().Infoln("Mem total: ", total, ", free: ", free)

	str := mi.MemToString()
	jklog.L().Infoln(str)

	str = mi.MemToStringM()
	jklog.L().Infoln(str)
}

func TestKFDiskInfo(t *testing.T) {
	mi := &KFSystemInfo{}
	total, free := mi.KFDiskInfo()
	jklog.L().Infoln("DiskInfo: ", total, ", free: ", free)

	str := mi.KFDiskString()
	jklog.L().Infoln(str)

	str = mi.KFDiskStringM()
	jklog.L().Infoln(str)
}

func TestKFCPUUsage(t *testing.T) {
	mi := &KFSystemInfo{}
	cpu := mi.KFCPUInfo()

	jklog.L().Infoln("CPUUsage: ", cpu)

	str := mi.KFCPUToString()
	jklog.L().Infoln(str)
}
