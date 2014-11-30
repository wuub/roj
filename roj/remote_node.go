package roj

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/armon/consul-api"
)

type RemoteNode struct {
	service *consulapi.ServiceEntry
	sysinfo *SysInfo
}

func (r *RemoteNode) Name() string {
	return r.service.Node.Node
}

func (r *RemoteNode) ApiRoot() string {
	return fmt.Sprintf("http://%s:%d", r.service.Node.Address, r.service.Service.Port)
}

func (r *RemoteNode) Status() string {
	status := map[string]int{"critical": 0, "warning": 0, "passing": 0, "unknown": 0}
	for _, c := range r.service.Checks {
		status[c.Status] += 1
	}

	if status["critical"] > 0 {
		return "critical"
	} else if status["warning"] > 0 {
		return "warning"
	} else if status["unknown"] > 0 {
		return "unknown"
	}

	return "passing"
}

func (r *RemoteNode) SysInfo() *SysInfo {
	if r.sysinfo == nil {
		r.sysinfo = new(SysInfo)
		resp, err := http.Get(r.ApiRoot() + "/v1/node/sysinfo")
		if err != nil {
			return r.sysinfo
		}
		err = json.NewDecoder(resp.Body).Decode(&r.sysinfo)
		if err != nil {
			return r.sysinfo
		}
	}
	return r.sysinfo
}

func (r *RemoteNode) String() string {
	sysinfo := r.SysInfo()
	return fmt.Sprintf("%s\t%s\t%s\t%dMB\t%.2f%%\t%.2f/%.2f/%.2f",
		r.Name(), r.ApiRoot(), r.Status(),
		sysinfo.VirtualMemory.Total/(1000*1024), sysinfo.VirtualMemory.UsedPercent,
		sysinfo.LoadAvg.Load1, sysinfo.LoadAvg.Load5, sysinfo.LoadAvg.Load15,
	)
}
