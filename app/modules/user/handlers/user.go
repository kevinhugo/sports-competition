package handlers

import (
	"net/http"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	"sports-competition/app/modules/user/resources"
	"sports-competition/app/modules/user/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		UserService: *services.NewUserService(),
	}
}

// @Summary      User Login
// @Description  Login User ( first login with unique username is considered as register )
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        "Request Data" body resources.UserLogin true "Request Data"
// @Success      200  {object}  helpers.Response
// @Failure      404  {string}  "404 page not found"
// @Failure      400  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /sports-competition/v1/user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var userLoginData resources.UserLogin
	err := c.ShouldBindJSON(&userLoginData)
	if err != nil {
		logger.Error("Error while binding user login data.")
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.CreateBadRequestResponse("Invalid data."))
		return
	}

	loginResponse := h.UserService.Login(&userLoginData)
	c.JSON(200, loginResponse)
}

// @Summary      Get user Identity
// @Description  Get Current User Identity and stats with firstname as param
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer token"
// @Param        first_name   path      string  true  "First Name"
// @Success      200  {object}  helpers.Response
// @Failure      404  {string}  "404 page not found"
// @Failure      400  {object}  helpers.Response
// @Failure      500  {object}  helpers.Response
// @Router       /sports-competition/v1/user/identity/{first_name} [get]
func (h *UserHandler) GetUserIdentity(c *gin.Context) {
	tokenDataFromheader, _ := c.Get("tokenData")
	var tokenData helpers.AccessToken
	helpers.JsonToStruct(&tokenDataFromheader, &tokenData)

	var firstName string = c.Param("first_name")

	competitionResult := h.UserService.GetUserIdentity(&tokenData, firstName)
	c.JSON(200, competitionResult)
}
