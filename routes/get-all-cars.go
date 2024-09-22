package routes

import (
	"hwc/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all cars
// @Produce json
// @Success 200 {object} []db.Car
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /cars [get]
func GetAllCars() gin.HandlerFunc {
	return func(c *gin.Context) {
		cars, err := db.Db.GetCars()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cars)
	}
}
