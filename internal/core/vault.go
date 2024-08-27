package core

import "github.com/freitzzz/monica/internal/data"

var global vault

type vault struct {
	OSRepository    data.OSRepository
	UsageRepository data.UsageRepository
}

func init() {
	global = vault{
		OSRepository:    data.GOUtilsOSRepository{},
		UsageRepository: data.GopsUtilUsageRepository{},
	}
}

func Vault() vault {
	return global
}
