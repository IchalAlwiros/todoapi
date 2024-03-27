package routes

import (
	"github.com/gin-gonic/gin"
	"goapi.railway.app/cmd/api/controllers"
)

func RouterSetup()  *gin.Engine{
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
			"data" : "Welcome TODO Go",
		})
	})


	route.GET("/todo", controllers.GetAllTodo)
	route.GET("/todo/:id", controllers.GetOneTodo)
	route.POST("/todo", controllers.CreateTodo)
	route.PUT("/todo/:id", controllers.UpdateTodo)
	route.PATCH("/todo/:id", controllers.PatchTodo)
	route.DELETE("/todo/:id", controllers.RomoveTodo)

	return route
}