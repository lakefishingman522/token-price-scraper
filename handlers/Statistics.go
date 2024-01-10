package handlers

import (
	"net/http"

	"errors"

	"github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
	"github.com/CascadiaFoundation/CascadiaTokenScrapper/utils"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Statistics struct {
	TimeStamp uint64  `json:"timestamp"`
	P360      float64 `json:"p360"`
	P180      float64 `json:"p180"`
	P90       float64 `json:"p90"`
	P30       float64 `json:"p30"`
	P14       float64 `json:"p14"`
	P7        float64 `json:"p7"`
	P1        float64 `json:"p1"`
}

// @router   	  	/getStatistics    [get]
// @description	  	List queries
// @accept      	application/json
// @produce      	application/json
// @tags          	statistics
// @summary			List queries
// @id				queries_history
// @success	  		200 {object}  models.Query
// @failure	  		400 {object}  utils.HTTPError
func (h handler) Statistics(ctx *gin.Context) {
	var stats models.TokenStatisticsModel

	// Attempt to fetch the latest record
	result := h.DB.Order("updated_at desc").First(&stats)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Record not found, fetch data from other function
			stats = utils.GetPriceStatistics()
			// Insert new record into DB with data from function
			h.DB.Create(&stats)
			// fetch the latest record again
			result = h.DB.Order("updated_at desc").First(&stats)
		} else {
			// Handle other errors
			ctx.JSON(http.StatusBadRequest, utils.HTTPError{Message: result.Error.Error()})
		}
	}

	ret := Statistics{
		TimeStamp: stats.TimeStamp,
		P360:      stats.P360,
		P180:      stats.P180,
		P90:       stats.P90,
		P30:       stats.P30,
		P14:       stats.P14,
		P7:        stats.P7,
		P1:        stts.P1,
	}

	ctx.JSON(http.StatusOK, stats)
}
