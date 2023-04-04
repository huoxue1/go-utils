package log

import (
	"github.com/kataras/iris/v12/context"
	"time"
)

func GetIsisLogHandler(logger *Logger) context.Handler {
	return func(ctx *context.Context) {
		start := time.Now()

		ctx.Next()

		logger.Infof("[IRIS_LOG] request_id: (%v), remote_addr: (%v), [%v] %v time_cost: %dms",
			ctx.GetID(), ctx.Request().RemoteAddr, ctx.Request().Method, ctx.Request().URL.String(), time.Since(start).Nanoseconds()/1e6)
	}
}
