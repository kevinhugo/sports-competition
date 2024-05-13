package services

import (
	"net/http"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	sportRepository "sports-competition/app/modules/sport/repositories"
	userRepository "sports-competition/app/modules/user/repositories"
	userResources "sports-competition/app/modules/user/resources"
)

type UserService struct {
	userRepo  userRepository.UserRepository
	sportRepo sportRepository.SportRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepo:  userRepository.NewUserRepository(),
		sportRepo: sportRepository.NewSportRepository(),
	}
}

func (h *UserService) Login(loginData *userResources.UserLogin) helpers.Response {
	var checkResult error = h.loginValidate(loginData)
	if checkResult != nil {
		return helpers.CreateBadRequestResponse(checkResult.Error())
	}
	var hashedPassword string = helpers.SHA512(loginData.Password)

	userData, err := h.userRepo.GetUserByUsername(
		loginData.Username,
	)
	if err != nil {
		logger.Error("Something went wrong while trying to get user data, see logs below.")
		logger.Error(err)
		return helpers.CreateGeneralErrorResponse("Failed to get user data.")
	}

	if userData.ID == 0 {
		userData, err = h.userRepo.SaveUser(map[string]interface{}{
			"username": loginData.Username,
			"password": hashedPassword,
		})
	} else {
		if userData.Password != hashedPassword {
			return helpers.CreateUnauthorizedResponse("Invalid username or password.")
		}
	}

	token, err := helpers.GenerateAccessToken(userData.UserID, userData.Username)
	if err != nil {
		return helpers.CreateGeneralErrorResponse("Failed to generate token.")
	}

	return helpers.CreateResponse(http.StatusOK, "success", "Login success.", userResources.UserLoginResponse{
		AccessToken: token,
	})
}

func (h *UserService) GetUserIdentity(tokenData *helpers.AccessToken, firstName string) helpers.Response {
	var checkResult error = h.getUserIdentityValidate(firstName)
	if checkResult != nil {
		return helpers.CreateBadRequestResponse(checkResult.Error())
	}

	userData, err := h.userRepo.GetUserByUserID(tokenData.ID)
	if err != nil {
		logger.Error("Something went wrong while trying to get user data, see logs below.")
		logger.Error(err)
		return helpers.CreateGeneralErrorResponse("Failed to get user data.")
	}
	if userData.ID == 0 {
		return helpers.CreateBadRequestResponse("User data not found.")
	}

	existingSportData, err := h.sportRepo.GetByUserID(tokenData.ID)
	if err != nil {
		logger.Error("Something went wrong while trying to get sport data, see logs below.")
		logger.Error(err)
		return helpers.CreateGeneralErrorResponse("Failed to get sport data.")
	}

	return helpers.CreateResponse(http.StatusOK, "success", "Get user identity has been successfuly done.", userResources.GetUserIdentityResponse{
		Username:    userData.Username,
		FirstName:   firstName,
		Proficiency: existingSportData.Proficiency,
	})
}
