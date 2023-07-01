package v1

import (
	"github.com/gin-gonic/gin"
	"hackathon/internal/models"
	"net/http"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	api.POST("/register", h.registerStudent)
	api.POST("/auth", h.loginStudent)
}

func (h *Handler) registerStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if student.Email == "" || student.Name == "" || student.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "error: empty fields")
		return
	}

	err := h.studentService.RegisterStudent(c.Request.Context(), &student)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *Handler) loginStudent(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	if input.Email == "" || input.Password == "" {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	student, err := h.studentService.AuthenticateStudent(c.Request.Context(), input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, student)
}
