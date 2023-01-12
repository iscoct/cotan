package logger

var Standard Logger

func init() {
	Standard = NewLoggerWithDefaultOutputs()
}
