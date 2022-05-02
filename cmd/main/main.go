package main

import (
	"github.com/gin-gonic/gin"
	"github.com/husdeli/DrawPadService.git/internal/config"
	"github.com/husdeli/DrawPadService.git/internal/user"
	"github.com/husdeli/DrawPadService.git/pkg/logger"
)

func indexHandler(context *gin.Context) {
	context.Writer.Write([]byte("Hello world"))
}

func main() {
	var log = logger.GetLogger()
	var router = gin.Default()
	var conf = config.GetConfig()
	router.GET("/", indexHandler)

	user.RegisterHandlers(
		router,
		user.NewService(user.NewRepository()),
		log,
	)

	log.Fatal(router.Run(conf.Host + ":" + conf.Port))
}
