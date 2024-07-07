package handler

import (
	"app/backend/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

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

// GetRecomendation handles the request for sector recommendations.
// It binds the request body to a RecomendationInput struct,
// calls the service to get the recommendations, and returns the result as a JSON response.
// If there is an error, it returns a 400 or 500 status code with the corresponding error message.
func (h *Handler) GetRecomendation(c *gin.Context) {
	// Bind the request body to a RecomendationInput struct.
	var input service.RecomendationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// If there was an error, return a 400 status code with the error message.
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call the service to get the sector recommendations.
	arr, err := h.service.GetRecomendation(input)
	if err != nil {
		// If there was an error, return a 500 status code with the error message.
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Return a 200 status code with the recommendations as a JSON response.
	c.JSON(http.StatusOK, arr)
}
