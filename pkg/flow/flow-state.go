package flow

import "fmt"

func (f *FlowElement) storeState() {
	stateStore := "redis"
	state := map[string]string{
		"TID":          f.TID,
		"next-step":    f.jumptToStep,
		"current-step": f.currentStep,
		"ended":        fmt.Sprintf("%v", f.isEnded),
	}
	f.log.Debug("[x] StateStore", state)
	switch stateStore {
	case "redis":
		RedisStateStore(f.log)
	case "sqlite":
		SqliteStateStore(f.log)
	}
}
