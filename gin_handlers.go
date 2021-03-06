package pprof_help

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/pprof"
	rpprof "runtime/pprof"

	"runtime"
)

type Range struct {
	Name  string
	Start int
	End   int
}

func durationExceedsWriteTimeout(ctx context.Context, seconds float64) bool {
	srv, ok := ctx.Value(http.ServerContextKey).(*http.Server)
	return ok && srv.WriteTimeout != 0 && seconds >= srv.WriteTimeout.Seconds()
}

//func splitTrace(res ptrace.ParseResult) []Range {
//	params := &traceParams{
//		parsed:  res,
//		endTime: math.MaxInt64,
//	}
//	s, c := splittingTraceConsumer(100 << 20) // 100M
//	if err := generateTrace(params, c); err != nil {
//		dief("%v\n", err)
//	}
//	return s.Ranges
//}

func SetDebugHandlers(engine *gin.Engine) {
	engine.GET("/debug/pprof/", func(ctx *gin.Context) {
		pprof.Index(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/cmdline", func(ctx *gin.Context) {
		pprof.Cmdline(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/symbol", func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	})
	engine.POST("/debug/pprof/symbol", func(ctx *gin.Context) {
		pprof.Symbol(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/profile", func(ctx *gin.Context) {
		pprof.Profile(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/heap", func(ctx *gin.Context) {
		pprof.Handler("heap").ServeHTTP(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/goroutine", func(ctx *gin.Context) {
		pprof.Handler("goroutine").ServeHTTP(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/block", func(ctx *gin.Context) {
		pprof.Handler("block").ServeHTTP(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/threadcreate", func(ctx *gin.Context) {
		pprof.Handler("threadcreate").ServeHTTP(ctx.Writer, ctx.Request)
	})
	engine.GET("/debug/pprof/allocs", func(ctx *gin.Context) {
		pprof.Handler("allocs").ServeHTTP(ctx.Writer, ctx.Request)
	})
	engine.GET("debug/pprof/memory", func(ctx *gin.Context) {
		var buff bytes.Buffer

		runtime.GC() // get up-to-date statistics
		if err := rpprof.WriteHeapProfile(&buff); err != nil {
			panic(err)
		}
		ctx.Writer.Header().Set("X-Content-Type-Options", "nosniff")
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Disposition", `attachment; filename="prof.mem"`)
		if _, err := ctx.Writer.Write(buff.Bytes()); err != nil {
			panic(err)
		}
	})
	engine.GET("/debug/pprof/trace", func(ctx *gin.Context) {
		pprof.Trace(ctx.Writer, ctx.Request)
		//sec, err := strconv.ParseFloat(ctx.PostForm("seconds"), 64)
		//if sec <= 0 || err != nil {
		//	sec = 1
		//}
		//if durationExceedsWriteTimeout(ctx, sec) {
		//	ctx.JSON(400, fmt.Errorf("profile duration exceeds server's WriteTimeout"))
		//}
		//buff := ctx.Writer
		//if err := trace.Start(buff); err != nil {
		//	// trace.Start failed, so no writes yet.
		//	ctx.JSON(400, fmt.Errorf("Could not enable tracing: %s ", err))
		//	return
		//}
		//time.Sleep(time.Duration(sec * float64(time.Second)))
		//trace.Stop()

		//res, err := ptrace.Parse(&buff, "")
		//if err != nil {
		//	ctx.JSON(400, fmt.Errorf("failed to parse trace: %v ", err))
		//}

	})
}
