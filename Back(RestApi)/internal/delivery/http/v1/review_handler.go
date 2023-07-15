package v1

import (
	"hackathon/internal/models"
	"hackathon/pkg/filters"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initreviewsRoutes(api *gin.RouterGroup) {
	api.POST("/reviews/add", h.addReview)
	api.DELETE("/reviews/delete", h.DeleteReview)
	api.PATCH("/reviews/update", h.updateReview)
	api.GET("/reviews/list", h.listReviewsByProfessorId)
}

func (h *Handler) addReview(c *gin.Context) {
	var review models.Review
	if err := c.BindJSON(&review); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if review.Feedback == "" || review.Overall == nil {
		newErrorResponse(c, http.StatusBadRequest, "error: empty fields")
		return
	}

	_, err := h.reviewService.AddNewReview(c.Request.Context(), &review)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, review)
}

func (h *Handler) DeleteReview(c *gin.Context) {
	var input struct {
		ProfessorId string `json:"ProfessorId"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if input.ProfessorId == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty professorid")
		return
	}
	err := h.reviewService.DeleteReview(c.Request.Context(), input.ProfessorId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}

func (h *Handler) updateReview(c *gin.Context) {
	var input models.Review
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body\n"+err.Error())
		return
	}
	_, err := h.reviewService.UpdateReview(c.Request.Context(), &input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
}

func (h *Handler) listReviewsByProfessorId(c *gin.Context) {
	var input struct {
		ProfessorId string
		Filters     filters.Filters
	}

	reviews, _, err := h.reviewService.ListReviewsByProfessor(c.Request.Context(), input.ProfessorId, input.Filters)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}
