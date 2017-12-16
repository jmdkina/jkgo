// +build linux

//
// Author: jmdvirus@roamter.com
//
// System information
//
package jksys

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"jk/jklog"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

// Someday this name will get from file with different location
const (
	TotalName = "Total"
	FreeName  = "Free"
)

func NewSystemInfo() *KFSystemInfo {
	si := &KFSystemInfo{}
	si.KFCPUInfo()
	si.KFDiskInfo()
	si.KFMemInfo()
	si.KFSysBase()
	return si
}

func (si *KFSystemInfo) KFSysBase() error {
	d, err := ioutil.ReadFile("/etc/issue")
	if err != nil {
		return err
	}
	si.OSName = string(d)
	return nil
}

func (si *KFSystemInfo) KFMemInfo() (uint64, uint64) {
	var sys syscall.Sysinfo_t
	err := syscall.Sysinfo(&sys)
	if err != nil {
		return 0, 0
	}
	si.TotalRam = sys.Totalram
	si.FreeRam = sys.Freeram
	return sys.Totalram, sys.Freeram
}

func (si *KFSystemInfo) MemToString() string {
	return fmt.Sprintf("%s:%d,%s:%d", TotalName, si.TotalRam, FreeName, si.FreeRam)
}

func (si *KFSystemInfo) MemToStringM() string {
	total := si.TotalRam >> 20
	free := si.FreeRam >> 20
	return fmt.Sprintf("%s:%dM,%s:%dM", TotalName, total, FreeName, free)
}

// Get Disk info
// Return Total,Free
func (si *KFSystemInfo) KFDiskInfo() (uint64, uint64) {
	var sys syscall.Statfs_t
	err := syscall.Statfs("/", &sys)
	if err != nil {
		return 0, 0
	}
	si.TotalDisk = sys.Blocks * uint64(sys.Bsize)
	si.FreeDisk = sys.Bfree * uint64(sys.Bsize)
	return si.TotalDisk, si.FreeDisk
}

func (si *KFSystemInfo) KFDiskString() string {
	return fmt.Sprintf("%s:%d,%s:%d", TotalName, si.TotalDisk, FreeName, si.FreeDisk)
}

func (si *KFSystemInfo) KFDiskStringM() string {
	total := si.TotalDisk >> 30
	free := si.FreeDisk >> 30
	return fmt.Sprintf("%s:%dG,%s:%dG", TotalName, total, FreeName, free)
}

// Return cpu percentage
func (si *KFSystemInfo) KFCPUInfo() float64 {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		jklog.L().Errorln("error excute ps: ", err)
		jklog.Lfile().Errorln("error excute ps: ", err)
		return 0
	}
	processes := make([]*ProcessCPU, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		// log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			jklog.L().Errorln("parse error: ", err)
			jklog.Lfile().Errorln("parse error: ", err)
			return 0
		}
		processes = append(processes, &ProcessCPU{pid, cpu})
	}
	si.procCPU = processes
	for _, p := range processes {
		si.CPUUsage += p.cpu
	}

	return si.CPUUsage
}

func (si *KFSystemInfo) KFCPUToString() string {
	return fmt.Sprintf("%.2f%%", si.CPUUsage)
}
