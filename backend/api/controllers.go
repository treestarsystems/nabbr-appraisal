package api

import (
	"fmt"
	"nabbr-appraisal/load"
	"nabbr-appraisal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAppraisalChartTemplateJSON(c *gin.Context) {
	responseData := GenerateAppraisalEmptyChart()
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Appraisal chart template generated successfully",
		"payload": []utils.NabbrAppraisalChart{responseData},
	})
}

func PostAppraisalChartJSON(c *gin.Context) {
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
	if err := load.LoadDbData(appraisalChart); err != nil {
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
		"payload":    []utils.NabbrAppraisalChart{},
	})
}
