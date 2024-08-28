package state

import (
	"fmt"

	"github.com/freitzzz/monica/internal/core"
	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/schema"
)

var vault = core.Vault()

var (
	cacheDisk *core.CacheValue[schema.DiskUsage]
)

func GetCPUUsage() schema.CPUUsage {
	repo := vault.UsageRepository
	cpu, err := repo.CPU()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get CPU stats: %v", err))

		return schema.CPUUsage{}
	}

	return cpu
}

func GetRAMUsage() schema.RAMUsage {
	repo := vault.UsageRepository
	ram, err := repo.RAM()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get RAM stats: %v", err))

		return schema.RAMUsage{}
	}

	return ram
}

func GetDiskUsage() schema.DiskUsage {
	if cacheDisk == nil {
		cacheDisk = core.Cached(getDiskUsage())
	} else {
		cacheDisk.LookupOrRecache(getDiskUsage)
	}

	return cacheDisk.Value
}

func GetUptime() uint64 {
	repo := vault.UsageRepository
	uptime, err := repo.Uptime()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get uptime stats: %v", err))

		return 0
	}

	return uptime
}

func getDiskUsage() schema.DiskUsage {
	repo := vault.UsageRepository
	disk, err := repo.Disk()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get disk stats: %v", err))

		return schema.DiskUsage{}
	}

	return disk
}
