package router

import (
	"github.com/gin-gonic/gin"
	"example/go-study-proj/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	tasksGroup := router.Group("tasks") 
	{
		tasksController := new(controllers.TaskController)
		tasksGroup.GET("/", tasksController.GetMany)
	}

	return router
}
