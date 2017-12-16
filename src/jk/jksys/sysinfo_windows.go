// +build windows

//
// Author: jmdvirus@roamter.com
//
// System information
//
package jksys

import (
	"fmt"
	ssi "github.com/matishsiao/goInfo"
)

// Someday this name will get from file with different location
const (
	TotalName = "Total"
	FreeName  = "Free"
)

func NewSystemInfo() *KFSystemInfo {
	si := &KFSystemInfo{}
	info := *ssi.GetInfo()
	si.CPUs = info.CPUs
	si.Kernel = info.Kernel
	si.Core = info.Core
	si.OSName = info.OS
	si.Platform = info.Platform
	si.Hostname = info.Hostname
	si.GetAddrInfo()
	si.KFCPUInfo()
	si.KFDiskInfo()
	si.KFMemInfo()
	si.KFSysBase()
	return si
}

func (si *KFSystemInfo) GetAddrInfo() error {
	si.IPAddr = KFIPAddress()
	si.Mac, _ = KFLocalMac()
	return nil
}

func (si *KFSystemInfo) KFSysBase() error {
	si.OSName = "windows"
	return nil
}

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
	return 20
}

func (si *KFSystemInfo) KFCPUToString() string {
	return fmt.Sprintf("%.2f%%", si.CPUUsage)
}
