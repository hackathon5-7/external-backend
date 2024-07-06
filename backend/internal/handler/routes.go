package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes the routes for the handler.
//
// It returns a *gin.Engine representing the root router.
func (h *Handler) InitRoutes() *gin.Engine {
	// Create a new gin router.
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Create a new API group.
	api := router.Group("/api")
	{
		// Create a new external group.
		external := api.Group("/external")
		{
			// Create a new sectors group.
			sectors := external.Group("/sectors")
			{
				// Add a GET route to retrieve all sectors.
				sectors.GET("/", h.GetAllSectors)

				// Add a POST route to get sector recommendations.
				sectors.POST("/recom/", h.GetRecomendation)
			}
		}
	}

	// Return the root router.
	return router
}
