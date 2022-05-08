package room

import (
	"github.com/gin-gonic/gin"
	"github.com/husdeli/DrawPadService.git/pkg/logger"
)

const (
	RoomsUrl        = "rooms"
	ConcreteRoomUrl = "rooms/:id"
)

type Resource struct {
	logger  *logger.Logger
	service RoomsService
}

func NewResource() *Resource {
	return &Resource{
		logger:  logger.GetLogger(),
		service: NewService(NewRepository()),
	}
}

func RegisterHandlers(router *gin.Engine) {
	var res = NewResource()

	router.GET(ConcreteRoomUrl, res.getRoom)
	router.POST(RoomsUrl, res.createRoom)
}

func (res *Resource) getRoom(ctx *gin.Context) {
	roomId := ctx.Params.ByName("id")

	if roomId == "" {
		res.logger.Warnf("Invalid roomId %s", roomId)
		ctx.AbortWithStatus(400)
	}

	room, err := res.service.GetRoom(roomId)

	if err != nil {
		res.logger.Errorf("Room is not found %v", err)
		ctx.AbortWithStatus(404)
	}

	ctx.JSON(200, room)
}

func (res *Resource) createRoom(ctx *gin.Context) {
	room, err := res.service.CreateRoom()

	if err != nil {
		res.logger.Errorf("Create room failed. Error: %v", err)
		ctx.AbortWithStatus(500)
	}

	ctx.JSON(201, room)
}
