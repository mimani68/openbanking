package flow

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

//
//	example:
//
// 	TID := "jf93n22"
//	paymentFlow := flow.FlowGenerator(TID, p.log)
//	paymentFlow.
//		If("first-validation", firstValidation, "need-wait").
//		Do("calculate-sth-1", func() {}).
//		Do("calculate-sth-2", func() {}).
//		IfElse("second-validation", func() bool {
//			return true
//		}, "calculate-sth-3", "will-send-package")
//	paymentFlow.End()
//
//
func FlowGenerator(id string, log log.Ilogger) *FlowElement {
	log.Debug("Flow defining", nil)
	return &FlowElement{
		TID:             id,
		log:             log,
		firstStepPassed: false,
	}
}
