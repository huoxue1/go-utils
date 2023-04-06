package help

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func GetLevel(lvl string) []logrus.Level {
	l, err := logrus.ParseLevel(lvl)
	if err != nil {
		fmt.Println("ParseLevel error: " + err.Error())
		l = logrus.DebugLevel
	}
	switch l {
	case logrus.TraceLevel:
		return []logrus.Level{logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	case logrus.DebugLevel:
		return []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	case logrus.InfoLevel:
		return []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	case logrus.WarnLevel:
		return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	case logrus.ErrorLevel:
		return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	case logrus.FatalLevel:
		return []logrus.Level{logrus.FatalLevel, logrus.PanicLevel}
	case logrus.PanicLevel:
		return []logrus.Level{logrus.PanicLevel}
	default:
		return []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
	}
}
