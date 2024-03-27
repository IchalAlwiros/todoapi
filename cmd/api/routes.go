package main

import (
	"github.com/gin-gonic/gin"
	"goapi.railway.app/cmd/api/controllers"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()

	// Define the available routes
	// router.GET("/v1/healthcheck", app.healthcheckHandler)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
			"data" : "Welcome TODO Go",
		})
	})


	router.GET("/todo", controllers.GetAllTodo)
	router.GET("/todo/:id", controllers.GetOneTodo)
	router.POST("/todo", controllers.CreateTodo)
	router.PUT("/todo/:id", controllers.UpdateTodo)
	router.PATCH("/todo/:id", controllers.PatchTodo)
	router.DELETE("/todo/:id", controllers.RomoveTodo)

	return router
}
