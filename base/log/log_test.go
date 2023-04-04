package log

import (
	"github.com/huoxue1/go-utils/base/log/hook/file"
	"github.com/huoxue1/go-utils/base/log/hook/std"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestStdHook(t *testing.T) {
	logger := logrus.New()
	l := &Logger{logger}
	logrus.Infoln()
	l.AddHook(std.NewStdHook())
	l.WithField("app", "123").Infof("test")
}

func TestFileHook(t *testing.T) {
	logger := logrus.New()
	l := &Logger{logger}
	f, _ := os.OpenFile("test.log", os.O_RDWR|os.O_APPEND, 0666)
	hook, _ := file.NewFileHook(file.WithCustomWriter(f))
	l.AddHook(hook)

	l.AddHook(std.NewStdHook(std.WithCustomWriter(f)))

	l.WithField("app", "123").Infof("test")
}

func TestGlobalLogger(t *testing.T) {
	AddHook(std.NewStdHook())
	Infoln("123")
}
