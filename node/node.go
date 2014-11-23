package node

import (
	"os"

	"github.com/shirou/gopsutil"
)

type Node struct {
	Name string
}

func New() *Node {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return &Node{Name: name}
}

type SysInfo struct {
	LoadAvg       *gopsutil.LoadAvgStat
	CPUInfo       []gopsutil.CPUInfoStat
	HostInfo      *gopsutil.HostInfoStat
	SwapMemory    *gopsutil.SwapMemoryStat
	VirtualMemory *gopsutil.VirtualMemoryStat
}

func (n *Node) SysInfo() *SysInfo {

	cpuInfo, _ := gopsutil.CPUInfo()
	loadAvg, _ := gopsutil.LoadAvg()
	hostInfo, _ := gopsutil.HostInfo()

	swapMemory, _ := gopsutil.SwapMemory()
	virtualMemory, _ := gopsutil.VirtualMemory()

	return &SysInfo{
		HostInfo:      hostInfo,
		LoadAvg:       loadAvg,
		CPUInfo:       cpuInfo,
		SwapMemory:    swapMemory,
		VirtualMemory: virtualMemory,
	}
}
