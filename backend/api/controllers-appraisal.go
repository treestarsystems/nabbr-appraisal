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
	responseData := NewAppraisalChart()
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"Appraisal chart template generated successfully",
		[]utils.NabbrAppraisalChart{responseData},
	)
	c.JSON(http.StatusOK, apiResponse)
}

func PostPutAppraisalChartJSON(c *gin.Context) {
	appraisalId := c.Param("appraisalId")
	var appraisalChart utils.NabbrAppraisalChart
	if err := c.BindJSON(&appraisalChart); err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusBadRequest,
			fmt.Sprintf("Invalid input: %v", err),
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
		return
	}

	// Save the appraisalChart to the database
	returnedAppraisalId, err := load.LoadDbDataAppraisals(appraisalChart, appraisalId)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to save appraisal chart (%v)", err),
			[]string{},
		)
		c.JSON(http.StatusInternalServerError, apiResponse)
		return
	}

	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"Appraisal chart saved successfully",
		[]string{returnedAppraisalId},
	)
	c.JSON(http.StatusOK, apiResponse)
}

func GetAppraisalChartAllJSON(c *gin.Context) {
	// Get all appraisal charts
	appraisalCharts := retrieve.RetrieveDbDataAll()

	if len(appraisalCharts) == 0 {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusNotFound,
			"No charts available",
			[]string{},
		)
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"Appraisal charts retrieved successfully",
		appraisalCharts,
	)
	c.JSON(http.StatusOK, apiResponse)
}

func GetAppraisalChartByIdJSON(c *gin.Context) {
	appraisalId := c.Param("appraisalId")
	// Get the appraisal chart by appraisalId
	appraisalChart, err := retrieve.RetrieveDbDataBySearchTerm(appraisalId)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusNotFound,
			fmt.Sprintf("Failed to get appraisal chart by appraisalId (%v)", err),
			[]string{},
		)
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}
	if len(appraisalChart) == 0 {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusNotFound,
			"This chart is not available",
			[]string{},
		)
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}
	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"Appraisal chart retrieved successfully",
		[]utils.NabbrAppraisalChart{appraisalChart[0]},
	)
	c.JSON(http.StatusOK, apiResponse)
}

func DeleteAppraisalChartJSON(c *gin.Context) {
	appraisalId := c.Param("appraisalId")
	// To be honest this next section should not really happen.
	if appraisalId == "" || appraisalId == " " {
		apiResponse := utils.NewAPIResponse(
			"failure",
			http.StatusBadRequest,
			"Missing Appraisal ID",
			[]string{},
		)
		c.JSON(http.StatusBadRequest, apiResponse)
	}

	// Delete the appraisalChart from the database
	err := load.LoadDbDataAppraisalDelete(appraisalId)
	if err != nil {
		apiResponse := utils.NewAPIResponse(
			"failuire",
			http.StatusNotFound,
			fmt.Sprintf("Failed to delete appraisal chart (%v)", err),
			[]string{},
		)
		c.JSON(http.StatusNotFound, apiResponse)
		return
	}

	apiResponse := utils.NewAPIResponse(
		"success",
		http.StatusOK,
		"Appraisal chart deleted successfully",
		[]string{},
	)
	c.JSON(http.StatusOK, apiResponse)
}
