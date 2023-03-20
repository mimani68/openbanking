package flow

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

type FlowElement struct {
	log log.Ilogger

	TID             string
	firstStepPassed bool
	currentStep     string
	jumptToStep     string
	isEnded         bool
}

func (f *FlowElement) If(alias string, condition func() bool, failure string) *FlowElement {
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.jumptToStep == alias || f.jumptToStep == "" {
		f.currentStep = alias
		if f.jumptToStep == alias {
			f.jumptToStep = ""
		}
		state := condition()
		if !state {
			f.jumptToStep = failure
		}
		f.log.Debug("If condition", map[string]string{
			"type":         "flow",
			"TID":          f.TID,
			"next-step":    f.jumptToStep,
			"current-step": f.currentStep,
		})
		f.storeState()
	}
	return f
}

func (f *FlowElement) IfElse(alias string, condition func() bool, success string, failure string) *FlowElement {
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.jumptToStep == alias || f.jumptToStep == "" {
		f.currentStep = alias
		if f.jumptToStep == alias {
			f.jumptToStep = ""
		}
		state := condition()
		if !state {
			f.jumptToStep = failure
		} else {
			f.jumptToStep = success
		}
		f.log.Debug("IfElse condition", map[string]string{
			"type":         "flow",
			"TID":          f.TID,
			"next-step":    f.jumptToStep,
			"current-step": f.currentStep,
		})
		f.storeState()
	}
	return f
}

// Any kind of operation which dosen't need valid and clear output
// you can assume this functionality as pipline
func (f *FlowElement) Do(alias string, cb func(), exceptSteps ...string) *FlowElement {
	for _, item := range exceptSteps {
		if f.currentStep == item {
			return f
		}
	}
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.jumptToStep == "" || f.jumptToStep == alias {
		f.currentStep = alias
		if f.jumptToStep == alias {
			f.jumptToStep = ""
		}
		cb()
		f.log.Debug("Do callback", map[string]string{
			"type":         "flow",
			"TID":          f.TID,
			"next-step":    f.jumptToStep,
			"current-step": f.currentStep,
		})
		f.storeState()
	}
	return f
}

func (f *FlowElement) End() {
	f.isEnded = true
	f.storeState()
}
