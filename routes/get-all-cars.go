package routes

import (
	"hwc/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all cars
// @Produce json
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} []db.Car
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /cars [get]
func GetAllCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := c.Query("limit")
		offset := c.Query("offset")

		if limit == "" {
			limit = "100"
		}

		if offset == "" {
			offset = "0"
		}

		intLimit, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "limit must be an integer"})
			return
		}

		intOffset, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "offset must be an integer"})
			return
		}

		cars, err := db.Db.GetCars(intLimit, intOffset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cars)
	}
}
