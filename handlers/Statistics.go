package handlers

import (
	"net/http"

	"errors"

	"github.com/CascadiaFoundation/CascadiaTokenScrapper/models"
	"github.com/CascadiaFoundation/CascadiaTokenScrapper/utils"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

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

	ctx.JSON(http.StatusOK, stats)
}
