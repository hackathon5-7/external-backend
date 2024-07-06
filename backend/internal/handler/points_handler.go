package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSectors handles the GET /sectors endpoint.
// It retrieves all sectors from the service and returns them as JSON.
// If there is an error, it returns a 500 status code with the error message.
// Otherwise, it returns a 200 status code with the sectors.
func (h *Handler) GetAllPoints(c *gin.Context) {
	// Retrieve all sectors from the service.
	sectors, err := h.service.GetAllPoints()

	// If there was an error, return a 500 status code with the error message.
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Return a 200 status code with the sectors as JSON.
	c.JSON(http.StatusOK, sectors)
}

func (h *Handler) GetRecomendation(c *gin.Context) {

}
