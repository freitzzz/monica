package mq

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/freitzzz/monica/internal/core"
	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/schema"
	"github.com/pebbe/zmq4"
)

const (
	delayLookupStatsMSEnvKey = "delay_lookup_stats_ms"
	lookupDelayFallbackMS    = 1000
)

var (
	delayLookupStatsMS = os.Getenv(serverHostEnvKey)
	lookupDelay        time.Duration
)

func init() {
	if delay, err := strconv.ParseInt(delayLookupStatsMS, 10, 64); err == nil {
		lookupDelay = time.Duration(delay) * time.Millisecond
	} else {
		lookupDelay = time.Duration(lookupDelayFallbackMS) * time.Millisecond

		logging.Aspirador.Error(
			fmt.Sprintf(
				"failed to convert %s from env key (%s). falling back to %d",
				delayLookupStatsMS, delayLookupStatsMSEnvKey, lookupDelayFallbackMS,
			),
		)
	}
}

func RegisterPub(s *zmq4.Socket) {
	var err error
	registered := false

	for {
		if !registered {
			registered, err = publishNode(s)
		} else {
			_, err = publishStats(s)
		}

		if err != nil {
			logging.Aspirador.Error(fmt.Sprintf("failed publishing message %v", err))
		}

		time.Sleep(lookupDelay)
	}
}

func publishNode(s *zmq4.Socket) (bool, error) {
	vault := core.Vault()
	repo := vault.OSRepository
	rs, err := core.Collect(repo.Hostname, repo.Type, repo.Hardware)

	if err != nil {
		return false, err
	}

	adv := schema.Advertisement{
		ID:       "minion-01",
		Hostname: rs[0],
		Type:     rs[1],
		Hardware: rs[2],
	}

	return SendRouteMessage(s, nodeAdvertisement, adv)
}

func publishStats(s *zmq4.Socket) (bool, error) {
	vault := core.Vault()
	repo := vault.UsageRepository
	rs, err := core.Collect(repo.CPU, repo.RAM)

	if err != nil {
		return false, err
	}

	stats := schema.Stats{
		ID:  "minion-01",
		CPU: rs[0],
		RAM: rs[1],
	}

	return SendRouteMessage(s, nodeStats, stats)
}
