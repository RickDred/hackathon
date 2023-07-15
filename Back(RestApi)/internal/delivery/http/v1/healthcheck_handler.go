package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initHealthCheckRoutes(api *gin.RouterGroup) {
	api.GET("/healthcheck", h.healthcheckHandler)
}

func (h *Handler) healthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"version": "1.0",
		},
	})
	//env := map[string]interface{}{
	//	"status": "available",
	//	"system_info": map[string]string{
	//		"version": "1.0",
	//	},
	//}

	//helpers.WriteJSON(w, http.StatusOK, env, nil)
}
