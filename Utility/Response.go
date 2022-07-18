package Utility

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func JSON(result *gin.Context, status int, payload interface{}) {
	var res ResponseData

	res.Status = status
	res.Data = payload

	result.JSON(200, res)
}
