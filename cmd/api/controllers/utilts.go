package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := c.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}


		pageLimitStr := c.DefaultQuery("page_limit", "10")
		pageLimit, _ := strconv.Atoi(pageLimitStr)
		switch {
		case pageLimit > 100:
			pageLimit = 100
		case pageLimit <= 0:
			pageLimit = 10
		}

		offset := (page - 1) * pageLimit
		return db.Offset(offset).Limit(pageLimit)
	}
}



func ResponseBody(c *gin.Context, statusCode any, data any) {
	switch statusCode{
	case 200:
		c.JSON(http.StatusOK, gin.H{"success" : true,"data": data})
	case 400:
		c.JSON(http.StatusBadRequest, gin.H{"success" : false,"message": data})
	}
}