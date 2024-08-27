package core

import "github.com/freitzzz/monica/internal/data"

type vault struct {
	OSRepository    data.OSRepository
	UsageRepository data.UsageRepository
}

var global vault

func init() {
	global = vault{
		OSRepository:    data.GOUtilsOSRepository{},
		UsageRepository: data.GopsUtilUsageRepository{},
	}
}

func Vault() vault {
	return global
}
