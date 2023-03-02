package flow

import "github.com/mimani68/fintech-core/pkg/log"

func SqliteStateStore(log log.Ilogger) {
	state := map[string]string{
		"adaptor": "sqlite",
	}
	log.Debug("Store state in sqlite", state)
}
