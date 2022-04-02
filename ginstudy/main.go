package main

import (
	"ginstudy/logger"
	"ginstudy/router"
	"ginstudy/service"
	"math/rand"
	"time"

	// "github.com/gin-contrib/sessions"
	// "github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

//中间件
func auditLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path

		ctx.Next()

		logger.Infof("%s | %3d | %15s | %-7s | %s\n%s", start.Format("2006/01/02 - 15:04:05"),
			ctx.Writer.Status(),
			ctx.ClientIP(),
			ctx.Request.Method,
			path,
			ctx.Errors.ByType(gin.ErrorTypePrivate).String())
	}
}

func main() {

	service.Init()

	rand.Seed(time.Now().UnixNano())
	engine := gin.Default()

	// store := memstore.NewStore([]byte("secret"))
	// engine.Use(sessions.Sessions("mysession", store))

	engine.Use(auditLogger())
	engine.Use(func(ctx *gin.Context) {

		ctx.Next()

	})
	engine.LoadHTMLGlob("templates/*")
	engine.Static("/static", "./static")
	router.Route(engine)
	_ = engine.Run() //listen and serve on 8080
}
