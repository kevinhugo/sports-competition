package handlers

import (
	"net/http"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	"sports-competition/app/modules/sport/resources"
	"sports-competition/app/modules/sport/services"

	"github.com/gin-gonic/gin"
)

type SportHandler struct {
	sportService services.SportService
}

func NewSportHandler() *SportHandler {
	return &SportHandler{
		sportService: *services.NewSportService(),
	}
}

// ShowAccount godoc
// @Summary      Begin Competition
// @Description  Start the competititon
// @Tags         sport
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        "Request Data" body resources.SportBeginCompetition true "Request Data"
// @Success      200  {object}  helpers.Response
// @Failure      404  {string}  "404 page not found"
// @Failure      400  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /sports-competition/v1/sport/begin-competition [post]
func (h *SportHandler) BeginCompetition(c *gin.Context) {
	tokenDataFromheader, _ := c.Get("tokenData")
	var tokenData helpers.AccessToken
	helpers.JsonToStruct(&tokenDataFromheader, &tokenData)

	var beginCompoetititonData resources.SportBeginCompetition
	err := c.ShouldBindJSON(&beginCompoetititonData)
	if err != nil {
		logger.Error("Error while binding competition data.")
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.CreateBadRequestResponse("Invalid data."))
		return
	}

	competitionResult := h.sportService.BeginCompetition(&tokenData, &beginCompoetititonData)
	c.JSON(200, competitionResult)
}
