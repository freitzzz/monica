package data

import (
	"time"

	"github.com/freitzzz/monica/internal/schema"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type UsageRepository interface {
	CPU() (schema.CPUUsage, error)
	RAM() (schema.RAMUsage, error)
	Disk() (schema.DiskUsage, error)
	Uptime() (uint64, error)
}

type GopsUtilUsageRepository struct {
}

func (r GopsUtilUsageRepository) CPU() (schema.CPUUsage, error) {
	var usage schema.CPUUsage
	stats, err := cpu.Percent(time.Duration(0), false)

	if err == nil {
		usage.Used = stats[0]
	}

	return usage, err
}

func (r GopsUtilUsageRepository) RAM() (schema.RAMUsage, error) {
	var usage schema.RAMUsage
	vm, err := mem.VirtualMemory()

	if err == nil {
		usage.Used = vm.UsedPercent
		usage.Total = vm.Total
	}

	return usage, err
}

func (r GopsUtilUsageRepository) Disk() (schema.DiskUsage, error) {
	var usage schema.DiskUsage
	dsk, err := disk.Usage("/")

	if err == nil {
		usage.Used = dsk.UsedPercent
		usage.Total = dsk.Total
	}

	return usage, nil
}

func (r GopsUtilUsageRepository) Uptime() (uint64, error) {
	return host.Uptime()
}
