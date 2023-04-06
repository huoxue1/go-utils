package xorm_log

import (
	"github.com/huoxue1/go-utils/base/log"
	"github.com/huoxue1/go-utils/help"
	"github.com/sirupsen/logrus"
	log1 "xorm.io/xorm/log"
)

type XormLog struct {
	logger  *log.Logger
	showSql bool
	level   string
	levels  []logrus.Level
}

func (x *XormLog) BeforeSQL(context log1.LogContext) {

}

func (x *XormLog) AfterSQL(context log1.LogContext) {
	if x.showSql && help.In(logrus.DebugLevel, x.levels) {
		x.logger.WithField("model", "xorm").Debugf("sql: %v,args:%v,result:%v,time:%vs\n", context.SQL, context.Args, context.Result, context.ExecuteTime.Seconds())
	}
}

func (x *XormLog) Debugf(format string, v ...interface{}) {
	if help.In(logrus.DebugLevel, x.levels) {
		x.logger.WithField("model", "xorm").Debugf(format, v...)
	}

}

func (x *XormLog) Errorf(format string, v ...interface{}) {
	if help.In(logrus.ErrorLevel, x.levels) {
		x.logger.WithField("model", "xorm").Debugf(format, v...)

	}
}

func (x *XormLog) Infof(format string, v ...interface{}) {
	if help.In(logrus.InfoLevel, x.levels) {
		x.logger.WithField("model", "xorm").Infof(format, v...)

	}
}

func (x *XormLog) Warnf(format string, v ...interface{}) {
	if help.In(logrus.WarnLevel, x.levels) {
		x.logger.WithField("model", "xorm").Warnf(format, v...)

	}
}

func (x *XormLog) Level() log1.LogLevel {
	switch x.level {
	case "debug":
		return log1.LOG_DEBUG
	case "info":
		return log1.LOG_INFO
	case "warn":
		return log1.LOG_WARNING
	default:
		return log1.LOG_UNKNOWN
	}
}

func (x *XormLog) SetLevel(l log1.LogLevel) {
	switch l {
	case log1.LOG_INFO:
		x.level = "info"
	case log1.LOG_DEBUG:
		x.level = "debug"
	case log1.LOG_WARNING:
		x.level = "warn"
	default:
		x.level = "debug"

	}
}

func (x *XormLog) ShowSQL(show ...bool) {
	if len(show) > 0 {
		x.showSql = show[0]
	}

}

func (x *XormLog) IsShowSQL() bool {
	return x.showSql
}

func GetXormLogger(logger *log.Logger, level string, showSql bool) log1.ContextLogger {
	x := &XormLog{
		logger:  logger,
		showSql: showSql,
		level:   level,
		levels:  help.GetLevel(level),
	}
	return x
}
