package api

import (
	"github.com/gin-gonic/gin"
)

func RoutesAppraisal(router *gin.Engine) {
	appraisal := router.Group("/api/v1/appraisal")
	{
		appraisal.GET("/template/chart", GetAppraisalChartTemplateJSON)
		appraisal.POST("/chart", PostAppraisalChartJSON)
	}
}
