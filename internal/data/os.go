package data

import (
	"os"
	"runtime"
)

type OSRepository interface {
	Hostname() (string, error)
	Type() (string, error)
	Hardware() (string, error)
}

type GOUtilsOSRepository struct {
}

func (r GOUtilsOSRepository) Hostname() (string, error) {
	return os.Hostname()
}

func (r GOUtilsOSRepository) Type() (string, error) {
	return runtime.GOOS, nil
}

func (r GOUtilsOSRepository) Distribution() (string, error) {
	return os.Hostname()
}

func (r GOUtilsOSRepository) Hardware() (string, error) {
	if _, err := os.Stat("/sys/firmware/devicetree/base/model"); err == nil {
		return "raspberry-pi", nil
	}

	return "server", nil
}
