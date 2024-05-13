package services

import (
	"net/http"
	"sports-competition/app/helpers"
	"sports-competition/app/logger"
	sportRepository "sports-competition/app/modules/sport/repositories"
	sportResources "sports-competition/app/modules/sport/resources"
)

type SportService struct {
	sportRepo sportRepository.SportRepository
}

func NewSportService() *SportService {
	return &SportService{
		sportRepo: sportRepository.NewSportRepository(),
	}
}

func (h *SportService) BeginCompetition(tokenData *helpers.AccessToken, competitionData *sportResources.SportBeginCompetition) helpers.Response {
	var checkResult error = h.beginCompetitionValidate(competitionData)
	if checkResult != nil {
		return helpers.CreateBadRequestResponse(checkResult.Error())
	}

	var maxProficiency uint = h.CalculateMaxProficiency(competitionData)

	existingSportData, err := h.sportRepo.GetByUserID(tokenData.ID)
	if err != nil {
		logger.Error("Something went wrong while trying to get sport data, see logs below.")
		logger.Error(err)
		return helpers.CreateGeneralErrorResponse("Failed to get sport data.")
	}
	if existingSportData.ID == 0 {
		_, err = h.sportRepo.AddSport(tokenData.ID, maxProficiency)
		if err != nil {
			logger.Error("Something went wrong while trying to add sport data, see logs below.")
			logger.Error(err)
			return helpers.CreateGeneralErrorResponse("Failed to add sport data.")
		}
	} else {
		if maxProficiency != existingSportData.Proficiency {
			err = h.sportRepo.UpdateSport(existingSportData, map[string]any{
				"proficiency": maxProficiency,
			})
			if err != nil {
				logger.Error("Something went wrong while trying to update sport data, see logs below.")
				logger.Error(err)
				return helpers.CreateGeneralErrorResponse("Failed to update sport data.")
			}
		}
	}

	return helpers.CreateResponse(http.StatusOK, "success", "Competition has been successfuly done.", sportResources.SportBeginCompetitionResult{
		MaxProficiency: maxProficiency,
	})
}

func (h *SportService) CalculateMaxProficiency(competitionData *sportResources.SportBeginCompetition) uint {
	var currentProficiency uint = competitionData.BaseSkill[1]
	var opponentList []sportResources.SportBeginCompetitionEachOpponentData = competitionData.GetDescendingOpponentList()

	for _, each := range opponentList {
		if each.Proficiency <= currentProficiency {
			currentProficiency += each.Exp
		} else {
			break
		}
	}
	return currentProficiency
}
