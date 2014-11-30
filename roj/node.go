package roj

import (
	"os"

	"github.com/shirou/gopsutil"
)

type Node struct {
	Name string
}

func NewNode() *Node {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &Node{Name: name}
}

type SysInfo struct {
	LoadAvg       gopsutil.LoadAvgStat
	CPUInfo       []gopsutil.CPUInfoStat
	HostInfo      gopsutil.HostInfoStat
	SwapMemory    gopsutil.SwapMemoryStat
	VirtualMemory gopsutil.VirtualMemoryStat
}

func (n *Node) SysInfo() *SysInfo {

	cpuInfo, err := gopsutil.CPUInfo()
	if err != nil {
		panic(err)
	}
	loadAvg, err := gopsutil.LoadAvg()
	if err != nil {
		panic(err)
	}
	hostInfo, err := gopsutil.HostInfo()
	if err != nil {
		panic(err)
	}
	swapMemory, err := gopsutil.SwapMemory()
	if err != nil {
		panic(err)
	}
	virtualMemory, err := gopsutil.VirtualMemory()
	if err != nil {
		panic(err)
	}

	return &SysInfo{
		HostInfo:      *hostInfo,
		LoadAvg:       *loadAvg,
		CPUInfo:       cpuInfo,
		SwapMemory:    *swapMemory,
		VirtualMemory: *virtualMemory,
	}
}
