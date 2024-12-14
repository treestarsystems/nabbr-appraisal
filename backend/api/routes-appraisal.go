package api

import (
	"nabbr-appraisal/utils"

	"github.com/gin-gonic/gin"
)

func RoutesAppraisal(router *gin.Engine) {
	appraisal := router.Group("/api/v1/appraisal")
	{
		// Set up authentication middleware
		appraisal.Use(utils.Authentication)
		// All routes that require authentication should be placed after this line
		appraisal.GET("/chart", GetAppraisalChartAllJSON)
		appraisal.GET("/chart/template", GetAppraisalChartTemplateJSON)
		appraisal.GET("/chart/:appraisalId", GetAppraisalChartByIdJSON)
		appraisal.POST("/chart", PostPutAppraisalChartJSON)
		appraisal.PUT("/chart/:appraisalId", PostPutAppraisalChartJSON)
		appraisal.DELETE("/chart/:appraisalId", DeleteAppraisalChartJSON)
	}
}
