package mq

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/freitzzz/monica/internal/logging"
	"github.com/freitzzz/monica/internal/schema"
	"github.com/freitzzz/monica/internal/state"
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
	adv := schema.NodeInfo{
		ID:           state.Id(),
		Hostname:     state.GetHostname(),
		Type:         state.GetType(),
		Hardware:     state.GetHardware(),
		Distribution: state.GetDistribution(),
	}

	return SendRouteMessage(s, publishNodeInformationRoute, adv)
}

func publishStats(s *zmq4.Socket) (bool, error) {
	stats := schema.NodeUsage{
		ID:     state.Id(),
		CPU:    state.GetCPUUsage(),
		RAM:    state.GetRAMUsage(),
		Disk:   state.GetDiskUsage(),
		Uptime: state.GetUptime(),
	}

	return SendRouteMessage(s, publishNodeStatsRoute, stats)
}
