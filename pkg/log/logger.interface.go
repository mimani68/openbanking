package log

type Ilogger interface {
	Info(message string, meta interface{})
	Debug(message string, meta interface{})
	Error(message string, meta interface{})
}
