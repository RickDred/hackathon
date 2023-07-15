package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hackathon/internal/models"
	"hackathon/pkg/filters"
	"hackathon/pkg/helpers"
	"hackathon/pkg/validator"
	"net/http"
)

func (h *Handler) initStudentsRoutes(api *gin.RouterGroup) {
	api.POST("/students/register", h.registerStudent)
	api.POST("/students/auth", h.loginStudent)
	api.DELETE("/students/delete", h.deleteStudent)
	api.PATCH("/students/update", h.updateStudent)
	api.GET("/students/list", h.listStudent)
}

func (h *Handler) listStudent(c *gin.Context) {
	var input struct {
		Name     string
		Email    string
		Page     int
		PageSize int
		filters.Filters
	}

	v := validator.New()

	qs := c.Request.URL.Query()
	fmt.Println(c.Request.URL)
	input.Name = helpers.ReadString(qs, "name", "")
	input.Email = helpers.ReadString(qs, "email", "")

	input.Filters.Page = helpers.ReadInt(qs, "page", 1, v)
	input.Filters.PageSize = helpers.ReadInt(qs, "page_size", 20, v)

	input.Filters.SortSafeList = []string{"Name", "-Name", "Email", "-Email", "Id", "-Id"}

	input.Filters.Sort = helpers.ReadString(qs, "sort", "id")

	students, _, err := h.studentService.ListStudents(c.Request.Context(), input.Name, input.Email, &input.Filters)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, students)
}

func (h *Handler) updateStudent(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body\n"+err.Error())
		return
	}
	err := h.studentService.UpdateStudent(c.Request.Context(), &models.Student{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
}

func (h *Handler) deleteStudent(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if input.Email == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty email")
		return
	}
	err := h.studentService.DeleteStudent(c.Request.Context(), input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
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
