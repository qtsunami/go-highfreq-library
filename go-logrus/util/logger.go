package util

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Logger
}

const (
	DebugLevel = logrus.DebugLevel
	InfoLevel  = logrus.InfoLevel
	WarnLevel  = logrus.WarnLevel
	ErrorLevel = logrus.ErrorLevel
	PanicLevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
)

func NewLogger() *Logger {
	return &Logger{
		logrus.New(),
	}
}

func (l Logger) LogHandler(m map[string]interface{}) *logrus.Entry {
	return l.WithFields(m)
}
