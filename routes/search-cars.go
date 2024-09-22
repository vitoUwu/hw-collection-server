package routes

import (
	"hwc/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Search cars
// @Produce json
// @Param q query string true "Query"
// @Success 200 {object} []db.SearchResult
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /search [get]
func SearchCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("q")

		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
			return
		}

		cars, err := db.Db.Search(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, cars)
	}
}
