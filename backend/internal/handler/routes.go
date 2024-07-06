package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		external := api.Group("/external")
		{
			sectors := external.Group("/sectors")
			{
				sectors.GET("/", h.GetAllSectors)
				sectors.GET("/rec/", h.GetRecomendation)
			}
		}
	}

	return router
}
