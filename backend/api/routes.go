package api

import (
	"github.com/gin-gonic/gin"
)

func RoutesAppraisal(router *gin.Engine) {
	appraisal := router.Group("/api/v1/appraisal")
	{
		appraisal.GET("/template/chart", GetAppraisalChartTemplateJSON)
		appraisal.GET("/chart", GetAppraisalChartAllJSON)
		appraisal.GET("/chart/:appraisalId", GetAppraisalChartByIdJSON)
		appraisal.POST("/chart", PostPutAppraisalChartJSON)
		appraisal.PUT("/chart/:appraisalId", PostPutAppraisalChartJSON)
	}
}
