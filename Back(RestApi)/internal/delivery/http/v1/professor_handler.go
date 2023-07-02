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

func (h *Handler) initProfessorsRoutes(api *gin.RouterGroup) {
	api.POST("/professors/add", h.addProfessor)
	api.DELETE("/professors/delete", h.deleteProfessor)
	api.PATCH("/professors/update", h.updateProfessor)
	api.GET("/professors/list", h.listProfessor)
}

func (h *Handler) addProfessor(c *gin.Context) {
	var professor models.Professor
	if err := c.BindJSON(&professor); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}
	if professor.Email == "" || professor.Name == "" {
		newErrorResponse(c, http.StatusBadRequest, "error: empty fields")
		return
	}

	err := h.professorService.AddProfessor(c.Request.Context(), &professor)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, professor)
}

func (h *Handler) deleteProfessor(c *gin.Context) {
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
	err := h.professorService.DeleteProfessor(c.Request.Context(), input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}

func (h *Handler) updateProfessor(c *gin.Context) {
	var input models.Professor
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body\n"+err.Error())
		return
	}
	err := h.professorService.UpdateProfessor(c.Request.Context(), &input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
}

func (h *Handler) listProfessor(c *gin.Context) {
	var input struct {
		Name       string
		Department string
		Degree     string
		Page       int
		PageSize   int
		filters.Filters
	}

	v := validator.New()

	qs := c.Request.URL.Query()
	fmt.Println(c.Request.URL)
	input.Name = helpers.ReadString(qs, "name", "")
	input.Department = helpers.ReadString(qs, "department", "")
	input.Degree = helpers.ReadString(qs, "degree", "")

	input.Filters.Page = helpers.ReadInt(qs, "page", 1, v)
	input.Filters.PageSize = helpers.ReadInt(qs, "page_size", 20, v)

	input.Filters.SortSafeList = []string{"Name", "-Name", "Email", "-Email", "Id", "-Id"}

	input.Filters.Sort = helpers.ReadString(qs, "sort", "id")

	professors, _, err := h.professorService.ListProfessors(c.Request.Context(), input.Name, input.Department, input.Degree, &input.Filters)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, professors)
}
