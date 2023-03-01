package flow

import (
	"github.com/mimani68/fintech-core/pkg/log"
)

type FlowElement struct {
	log log.Ilogger

	TID             string
	nextStep        string
	firstStepPassed bool
	currentStep     string
}

func (f *FlowElement) If(alias string, condition func() bool, failure string) *FlowElement {
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.currentStep == alias || f.nextStep == alias {
		f.currentStep = alias
		state := condition()
		if !state {
			f.nextStep = failure
		}
		f.log.Debug("If condition", map[string]string{
			"TID":          f.TID,
			"next-step":    f.nextStep,
			"current-step": f.currentStep,
		})
	}
	return f
}

func (f *FlowElement) IfElse(alias string, condition func() bool, success string, failure string) *FlowElement {
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.currentStep == alias || f.nextStep == alias {
		f.currentStep = alias
		state := condition()
		if state {
			f.nextStep = failure
		} else {
			f.nextStep = success
		}
		f.log.Debug("IfElse condition", map[string]string{
			"TID":          f.TID,
			"next-step":    f.nextStep,
			"current-step": f.currentStep,
		})
	}
	return f
}

func (f *FlowElement) Do(alias string, condition func()) *FlowElement {
	if !f.firstStepPassed {
		f.firstStepPassed = true
		f.currentStep = alias
	}
	if f.currentStep == alias || f.nextStep == alias || f.nextStep == "" {
		f.currentStep = alias
		condition()
		f.log.Debug("Do condition", map[string]string{
			"TID":          f.TID,
			"next-step":    f.nextStep,
			"current-step": f.currentStep,
		})
	}
	return f
}

func FlowGenerator(id string, log log.Ilogger) *FlowElement {
	log.Debug("Flow defining", nil)
	return &FlowElement{
		TID:             id,
		log:             log,
		firstStepPassed: false,
	}
}
