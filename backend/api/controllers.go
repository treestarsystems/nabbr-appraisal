package api

import (
	"fmt"
	"nabbr-appraisal/load"
	"nabbr-appraisal/retrieve"
	"nabbr-appraisal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAppraisalChartTemplateJSON(c *gin.Context) {
	responseData := GenerateAppraisalEmptyChart()
	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"httpStatus": http.StatusOK,
		"message":    "Appraisal chart template generated successfully",
		"payload":    []utils.NabbrAppraisalChart{responseData},
	})
}

func PostPutAppraisalChartJSON(c *gin.Context) {
	appraisalId := c.Param("appraisalId")
	var appraisalChart utils.NabbrAppraisalChart

	if err := c.BindJSON(&appraisalChart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":     "failuire",
			"httpStatus": http.StatusBadRequest,
			"message":    fmt.Sprintf("Invalid input: %v", err),
			"payload":    []utils.NabbrAppraisalChart{},
		})
		return
	}

	// Save the appraisalChart to the database
	returnedAppraisalId, err := load.LoadDbData(appraisalChart, appraisalId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     "failuire",
			"httpStatus": http.StatusInternalServerError,
			"message":    fmt.Sprintf("Failed to save appraisal chart: %v", err),
			"payload":    []utils.NabbrAppraisalChart{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"httpStatus": http.StatusOK,
		"message":    "Appraisal chart saved successfully",
		"payload":    []string{returnedAppraisalId},
	})
}

func GetAppraisalChartAllJSON(c *gin.Context) {
	// Get all appraisal charts
	appraisalCharts := retrieve.RetrieveDbDataAll()

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"httpStatus": http.StatusOK,
		"message":    "Appraisal charts retrieved successfully",
		"payload":    appraisalCharts,
	})
}

func GetAppraisalChartByIdJSON(c *gin.Context) {
	appraisalId := c.Param("appraisalId")

	// Get the appraisal chart by appraisalId
	appraisalChart, err := retrieve.RetrieveDbDataBySearchTerm(appraisalId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":     "failuire",
			"httpStatus": http.StatusNotFound,
			"message":    fmt.Sprintf("Failed to get appraisal chart by appraisalId: %v", err),
			"payload":    []utils.NabbrAppraisalChart{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"httpStatus": http.StatusOK,
		"message":    "Appraisal chart retrieved successfully",
		"payload":    []utils.NabbrAppraisalChart{appraisalChart[0]},
	})
}
