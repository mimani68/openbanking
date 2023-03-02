package flow

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

func FlowGenerator(id string, log log.Ilogger) *FlowElement {
	log.Debug("Flow defining", nil)
	return &FlowElement{
		TID:             id,
		log:             log,
		firstStepPassed: false,
	}
}
