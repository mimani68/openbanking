package flow

import "github.com/mimani68/fintech-core/pkg/log"

func RedisStateStore(log log.Ilogger) {
	state := map[string]string{
		"adaptor": "redis",
	}
	log.Debug("Store state in redis", state)
}
