package services

import (
	"errors"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	userResources "sports-competition/app/modules/user/resources"

	"regexp"
)

func (h *UserService) loginValidate(loginData *userResources.UserLogin) error {
	if loginData.Username == "" {
		return errors.New("username cannot be empty.")
	}
	regexCheck, err := regexp.MatchString(helpers.REGEX_ALPHA_NUMBER, loginData.Username)
	if err != nil {
		logger.Error("Something went wrong, see logs below.")
		logger.Error(err)
		return errors.New("Data cannot be processed.")
	}
	if regexCheck == false {
		return errors.New("Incorrect username format.")
	}

	if loginData.Password == "" {
		return errors.New("password cannot be empty.")
	}
	regexCheck, err = regexp.MatchString(helpers.REGEX_GENERAL, loginData.Password)
	if err != nil {
		logger.Error("Something went wrong, see logs below.")
		logger.Error(err)
		return errors.New("Data cannot be processed.")
	}
	if regexCheck == false {
		return errors.New("Incorrect password format.")
	}

	return nil
}

func (h *UserService) getUserIdentityValidate(firstName string) error {
	if firstName == "" {
		return errors.New("first_name cannot be empty.")
	}
	regexCheck, err := regexp.MatchString(helpers.REGEX_ALPHA, firstName)
	if err != nil {
		logger.Error("Something went wrong, see logs below.")
		logger.Error(err)
		return errors.New("Data cannot be processed.")
	}
	if regexCheck == false {
		return errors.New("Incorrect first_name format.")
	}

	return nil
}
