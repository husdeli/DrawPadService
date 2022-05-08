package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/husdeli/DrawPadService.git/internal/config"
	"github.com/husdeli/DrawPadService.git/internal/room"
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
	room.RegisterHandlers(router)

	log.Fatal(router.Run(fmt.Sprintf("%s:%d", conf.Host, conf.Port)))
}
