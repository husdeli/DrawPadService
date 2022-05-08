package user

import (
	"github.com/gin-gonic/gin"
	"github.com/husdeli/DrawPadService.git/pkg/logger"
)

const (
	UsersUrl        = "users"
	ConcreteUserUrl = "users/:id"
)

func RegisterHandlers(router *gin.Engine, service UserService, logger *logger.Logger) {
	var res = resource{service, logger}

	router.GET(UsersUrl, res.getUsers)
	router.GET(ConcreteUserUrl, res.getUserById)
	router.POST(UsersUrl, res.createUser)
	router.PUT(ConcreteUserUrl, res.updateUser)
	router.PATCH(ConcreteUserUrl, res.patchUser)
	router.DELETE(ConcreteUserUrl, res.deleteUser)
}

type resource struct {
	service UserService
	logger  *logger.Logger
}

func (resource *resource) getUserById(context *gin.Context) {
	var user, err = resource.service.GetUserById(context.Params.ByName("id"))
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(404)
	}
	context.JSON(200, user)
}

func (resource *resource) getUsers(context *gin.Context) {
	var users, err = resource.service.GetUsers()
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(500)
	}
	context.JSON(200, users)
}

func (resource *resource) createUser(context *gin.Context) {
	var user, err = resource.service.CreateUser()
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(500)
	}
	context.JSON(200, user)
}

func (resource *resource) updateUser(context *gin.Context) {
	var user, err = resource.service.UpdateUser(context.Params.ByName("id"))
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(500)
	}
	context.JSON(200, user)
}

func (resource *resource) patchUser(context *gin.Context) {
	var user, err = resource.service.PatchUser(context.Params.ByName("id"))
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(500)
	}
	context.JSON(200, user)
}

func (resource *resource) deleteUser(context *gin.Context) {
	var err = resource.service.DeleteUser(context.Params.ByName("id"))
	if err != nil {
		resource.logger.Error(err)
		context.AbortWithStatus(500)
	}
	context.Writer.WriteHeader(204)
}
