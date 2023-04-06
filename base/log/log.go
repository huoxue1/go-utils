package log

import (
	"github.com/huoxue1/go-utils/base/log/hook/file"
	"github.com/huoxue1/go-utils/base/log/hook/std"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger

	xormLevel   string
	xormShowSql bool
}

func New() *Logger {
	logger := logrus.New()
	logger.SetOutput(&nullWriter{})
	return &Logger{Logger: logger}
}

func NewStd() *Logger {
	logger := logrus.New()
	logger.SetOutput(&nullWriter{})
	logger.AddHook(std.NewStdHook())
	return &Logger{Logger: logger}
}

func Default() *Logger {
	logger := logrus.New()
	logger.SetOutput(&nullWriter{})
	hook := std.NewStdHook()
	logger.AddHook(hook)
	fileHook, _ := file.NewFileHook()
	logger.AddHook(fileHook)
	return &Logger{Logger: logger}
}

type nullWriter struct {
}

func (n2 *nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
