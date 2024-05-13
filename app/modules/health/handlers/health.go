package modules

import (
	"net/http"
	"sports-competition/app/helpers"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.CreateResponse(http.StatusOK, "I'm good.", "Really really good."))
}
