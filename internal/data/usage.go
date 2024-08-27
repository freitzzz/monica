package data

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

type UsageRepository interface {
	CPU() (float64, error)
	RAM() (float64, error)
}

type GopsUtilUsageRepository struct {
}

func (r GopsUtilUsageRepository) CPU() (float64, error) {
	stats, err := cpu.Percent(time.Duration(0), false)

	if err == nil {
		return stats[0], err
	}

	return 0, err
}

func (r GopsUtilUsageRepository) RAM() (float64, error) {
	vm, err := mem.VirtualMemory()

	if err == nil {
		return vm.UsedPercent, err
	}

	return 0, err
}
