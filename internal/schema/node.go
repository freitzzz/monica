package schema

// Autogenerated using https://transform.tools/json-to-go
type Node struct {
	ID string `json:"id"`
	OS struct {
		Hostname     string `json:"hostname"`
		Type         string `json:"type"`
		Distribution string `json:"distribution"`
		Hardware     string `json:"hardware"`
	} `json:"os"`
	Usage struct {
		CPU    CPUUsage  `json:"cpu"`
		RAM    RAMUsage  `json:"ram"`
		Disk   DiskUsage `json:"disk"`
		Uptime uint64    `json:"uptime"`
	} `json:"usage"`
}

type NodeInfo struct {
	ID           string
	Hostname     string
	Type         string
	Distribution string
	Hardware     string
}

type NodeUsage struct {
	ID     string
	Uptime uint64
	CPU    CPUUsage
	RAM    RAMUsage
	Disk   DiskUsage
}

type CPUUsage struct {
	Used float64
}

type RAMUsage struct {
	Used  float64
	Total uint64
}

type DiskUsage struct {
	Used  float64
	Total uint64
}

func ToNode(info NodeInfo, usage NodeUsage) Node {
	n := Node{
		ID: info.ID,
	}

	n.OS.Distribution = info.Distribution
	n.OS.Hardware = info.Hardware
	n.OS.Hostname = info.Hostname
	n.OS.Type = info.Type

	n.Usage.CPU = usage.CPU
	n.Usage.RAM = usage.RAM
	n.Usage.Disk = usage.Disk
	n.Usage.Uptime = usage.Uptime

	return n
}
