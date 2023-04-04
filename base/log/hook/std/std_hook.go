package std

import (
	"fmt"
	nested "github.com/Lyrics-you/sail-logrus-formatter/sailor"
	"github.com/huoxue1/go-utils/help"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Hook struct {
	levels       []logrus.Level
	formatter    logrus.Formatter
	customWriter io.Writer
}

type HookOption func(hook *Hook)

func WithCustomWriter(writer io.Writer) HookOption {
	return func(hook *Hook) {
		hook.customWriter = io.MultiWriter(hook.customWriter, writer)
	}
}

func WithFormatter(formatter logrus.Formatter) HookOption {
	return func(hook *Hook) {
		hook.formatter = formatter
	}
}

func WithLevel(level string) HookOption {
	return func(hook *Hook) {
		lvl, err := logrus.ParseLevel(level)
		if err != nil {
			fmt.Println("ParseLevel error: " + err.Error())
			lvl = logrus.DebugLevel
		}
		switch lvl {
		case logrus.TraceLevel:
			hook.levels = []logrus.Level{logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
		case logrus.DebugLevel:
			hook.levels = []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
		case logrus.InfoLevel:
			hook.levels = []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
		case logrus.WarnLevel:
			hook.levels = []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
		case logrus.ErrorLevel:
			hook.levels = []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
		case logrus.FatalLevel:
			hook.levels = []logrus.Level{logrus.FatalLevel, logrus.PanicLevel}
		case logrus.PanicLevel:
			hook.levels = []logrus.Level{logrus.PanicLevel}
		}
	}
}

func NewStdHook(options ...HookOption) *Hook {
	h := &Hook{
		levels:       []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
		customWriter: os.Stdout,
		formatter: &nested.Formatter{
			FieldsOrder:           nil,
			TimeStampFormat:       "2006-01-02 15:04:05",
			CharStampFormat:       "",
			HideKeys:              false,
			Position:              true,
			Colors:                true,
			FieldsColors:          true,
			FieldsSpace:           true,
			ShowFullLevel:         false,
			LowerCaseLevel:        true,
			TrimMessages:          true,
			CallerFirst:           false,
			CustomCallerFormatter: nil,
		},
	}
	for _, option := range options {
		option(h)
	}

	return h
}

func (s *Hook) Levels() []logrus.Level {
	return s.levels
}

func (s *Hook) Fire(entry *logrus.Entry) error {
	if help.In(entry.Level, s.levels) {
		data, err := s.formatter.Format(entry)
		if err != nil {
			return err
		}
		_, err = s.customWriter.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}
