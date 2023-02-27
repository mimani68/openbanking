package main

import (
	. "github.com/mimani68/fintech-core/pkg/log"
)

func main() {
	logger := Log()
	logger.Info("App is about to start", nil)
	logger.Info("Running", map[string]string{
		"developer-mode": "True",
	})
	logger.Debug("Current version for steady connectin in all distribution is LTS 11", map[string]string{
		"version": "1",
	})
	logger.Error("Port definition", map[string]string{
		"expected": "3000",
		"current":  "NULL",
	})
}
