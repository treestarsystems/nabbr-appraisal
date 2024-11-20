package api

import (
	"github.com/gin-gonic/gin"
)

func GetAppraisalChartTemplateJSON(c *gin.Context) {
	responseData := GenerateAppraisalEmptyChart()
	c.JSON(200, responseData)
}
