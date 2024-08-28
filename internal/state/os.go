package state

import (
	"fmt"

	"github.com/freitzzz/monica/internal/logging"
)

func GetHostname() string {
	repo := vault.OSRepository
	hostname, err := repo.Hostname()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get hostname: %v", err))

		return "-"
	}

	return hostname
}

func GetType() string {
	repo := vault.OSRepository
	ttype, err := repo.Type()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get OS type: %v", err))

		return "-"
	}

	return ttype
}

func GetDistribution() string {
	return "-"
}

func GetHardware() string {
	repo := vault.OSRepository
	hardware, err := repo.Hardware()

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("failed to get hardware information: %v", err))

		return "-"
	}

	return hardware
}
