package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthStatus struct {
	Status string `json:"status"`
}

type Status struct {
	Version string `json:"version"`
	Date    int64  `json:"date"`
}

// @router   	  	/    [get]
// @description	  	Show current status
// @accept      	application/json
// @produce      	application/json
// @tags          	status
// @summary			Show current status
// @id				query_status
// @success	  		200 {object}  handlers.Status
func (h handler) GetStatus(ctx *gin.Context) {
	currentDate := time.Now().Unix()
	// Read hostname, password, dbname and username from environment variables
	version := os.Getenv("VERSION")

	var response = Status{
		Version: version,
		Date:    currentDate,
	}

	ctx.IndentedJSON(http.StatusOK, response)
}

// @router   	  	/health    [get]
// @description	  	Show health status
// @accept      	application/json
// @produce      	application/json
// @tags          	health
// @summary			Show health status
// @id				query_health
// @success	  		200 {object}  handlers.HealthStatus
func (h handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HealthStatus{
		Status: "OK",
	})
}
