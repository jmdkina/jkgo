// +build darwin

//
// Author: jmdvirus@roamter.com
//
// System information
//
package jksys

import (
	"bytes"
	"fmt"
	"jk/jklog"
	"os/exec"
	"strconv"
	"strings"
)

// Get CPU info
type ProcessCPU struct {
	pid int
	cpu float64
}

type KFSystemInfo struct {
	TotalRam uint64
	FreeRam  uint64

	TotalDisk uint64
	FreeDisk  uint64

	CPUUsage float64
	procCPU  []*ProcessCPU
}

// Someday this name will get from file with different location
const (
	// TotalName = "Total"
	TotalName = "总计"
	// FreeName  = "Free"
	FreeName = "剩余"
)

func (si *KFSystemInfo) KFMemInfo() (uint64, uint64) {
	return 0, 0
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
	return 0, 0
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
