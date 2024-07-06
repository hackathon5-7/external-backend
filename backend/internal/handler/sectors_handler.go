package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type RecomendationInput struct {
	Filters Filters `json:"filters" binding:"required"`
}

type Filters struct {
	AgeFrom  int         `json:"ageFrom" binding:"required"`
	AgeTo    int         `json:"ageTo" binding:"required"`
	Gender   string      `json:"gender" binding:"required"`
	Income   IncomeInput `json:"income" binding:"required"`
	Quantity int         `json:"quantity" binding:"required"`
}

type IncomeInput struct {
	A bool `json:"a" binding:"required"`
	B bool `json:"b" binding:"required"`
	C bool `json:"c" binding:"required"`
}

// GetAllSectors retrieves all sectors from the service and returns them as JSON.
// If there is an error, it returns a 500 status code with the error message.
// Otherwise, it returns a 200 status code with the sectors.
func (h *Handler) GetAllSectors(c *gin.Context) {
	// Retrieve all sectors from the service.
	sectors, err := h.service.GetSectorsArray(os.Getenv("SECTORS_PATH_FRONT"))

	// If there was an error, return a 500 status code with the error message.
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Return a 200 status code with the sectors as JSON.
	c.JSON(http.StatusOK, sectors)
}

func (h *Handler) GetRecomendation(c *gin.Context) {
	var input RecomendationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	arr, err := h.service.GetRecomendation()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, arr)
}
