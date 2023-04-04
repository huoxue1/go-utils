package file

import (
	"fmt"
	hooks "github.com/huoxue1/go-utils/base/log/hook"
	"github.com/huoxue1/go-utils/help"
	rotates "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"path"
	"time"
)

type HookOption func(hook *Hook)

type Hook struct {
	levels    []logrus.Level
	formatter logrus.Formatter
	writer    io.Writer
	Dir       string
}

func (h *Hook) Levels() []logrus.Level {
	return h.levels
}

func (h *Hook) Fire(entry *logrus.Entry) error {
	if help.In(entry.Level, h.levels) {
		dup := entry.Dup()
		if entry.HasCaller() && entry.Caller != nil {
			dup.Data["caller"] = fmt.Sprintf(
				"%s:%d %s",
				entry.Caller.File,
				entry.Caller.Line,
				entry.Caller.Function)
		} else {
			dup.Data["caller"] = hooks.FindCaller(5)
		}

		data, err := h.formatter.Format(dup)
		if err != nil {
			return err
		}
		_, err = h.writer.Write(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func WithFormatter(formatter logrus.Formatter) HookOption {
	return func(hook *Hook) {
		hook.formatter = formatter
	}
}
func WithCustomWriter(writer io.Writer) HookOption {
	return func(hook *Hook) {
		if hook.writer != nil {
			hook.writer = io.MultiWriter(hook.writer, writer)
		} else {
			hook.writer = writer
		}

	}
}

func WithDir(dir string) HookOption {
	return func(hook *Hook) {
		hook.Dir = dir
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

func NewFileHook(option ...HookOption) (*Hook, error) {
	h := &Hook{
		levels: []logrus.Level{logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
		formatter: &logrus.JSONFormatter{
			TimestampFormat:   time.DateTime,
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		},
		Dir: "./logs",
	}
	for _, hookOption := range option {
		hookOption(h)
	}
	w, err := rotates.New(path.Join(h.Dir, "%Y-%m-%d.log"), rotates.WithRotationTime(time.Hour*24), rotates.WithMaxAge(time.Hour*24*7), rotates.WithLinkName(path.Join(h.Dir, "new.log")))
	if err != nil {
		return nil, err
	}
	if h.writer != nil {
		h.writer = io.MultiWriter(w, h.writer)
	} else {
		h.writer = w
	}

	return h, nil
}
