package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goapi.railway.app/cmd/api/database"
	"goapi.railway.app/cmd/api/models"
)


  


func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil{
		ResponseBody(c, http.StatusBadRequest,err.Error())
		return
	}

	database.DB.Create(&todo)
	ResponseBody(c, http.StatusOK, todo)
}


func GetAllTodo(c *gin.Context) {
	var todo []models.Todo
	
	database.DB.Scopes(Paginate(c)).Find(&todo)
	ResponseBody(c, http.StatusOK, todo)
}



func GetOneTodo(c *gin.Context) {
	var todo models.Todo

	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil{
		ResponseBody(c, http.StatusBadRequest, "Record not found!")
		return
	}

	ResponseBody(c, http.StatusOK, todo)
}


func UpdateTodo(c *gin.Context) {
	var todo models.Todo

	// VALIDATION TO CHECK FOUND ID DATA
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil{
		ResponseBody(c, http.StatusBadRequest, "Record not found!")
		return
	}

	// INPUT VALIDATION
	if err := c.ShouldBindJSON(&todo); err != nil{
		ResponseBody(c, http.StatusBadRequest,err.Error())
		return
	}

	database.DB.Save(&todo)	
	ResponseBody(c, http.StatusOK, todo)
}


func PatchTodo(c *gin.Context) {
	var todo models.Todo

	// VALIDATION TO CHECK FOUND ID DATA
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil{
		ResponseBody(c, http.StatusBadRequest, "Record not found!")
		return
	}
	
	var updateTodo models.Todo
	if err := c.BindJSON(&updateTodo); err != nil {
		ResponseBody(c, http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Model(&todo).Updates(updateTodo)
	ResponseBody(c, http.StatusOK, todo)
}


func RomoveTodo(c *gin.Context) {
	var todo models.Todo

	// VALIDATION TO CHECK FOUND ID DATA
	if err := database.DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil{
		ResponseBody(c, http.StatusBadRequest, "Record not found!")
		return
	}

	database.DB.Delete(&todo)
	ResponseBody(c, http.StatusOK, todo)
}