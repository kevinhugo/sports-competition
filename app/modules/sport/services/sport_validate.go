package services

import (
	"errors"
	"fmt"
	sportResources "sports-competition/app/modules/sport/resources"
	"strconv"
)

func (h *SportService) beginCompetitionValidate(competitionData *sportResources.SportBeginCompetition) error {
	if len(competitionData.BaseSkill) != 2 {
		return errors.New("base_skill is expected to be an array of unsigned integer and must have length of 2.")
	}
	if competitionData.BaseSkill[0] < 1 {
		return errors.New("base_skill array index of 0 value cannot be less than 1.")
	}
	if competitionData.BaseSkill[1] > 100 {
		return errors.New("base_skill array index of 1 value cannot be greater than 100.")
	}

	if int(competitionData.BaseSkill[0]) != len(competitionData.OpponentProficiencies) {
		return errors.New("opponent_proficiencies array length and must be the same as the value of base_skill array index of 0.")
	}

	if int(competitionData.BaseSkill[0]) != len(competitionData.OpponentExps) {
		return errors.New("opponent_exps array length and must be the same as the value of base_skill array index of 0.")
	}

	for index, each := range competitionData.OpponentProficiencies {
		if each < 1 {
			return errors.New(fmt.Sprintf("opponent_proficiencies array of index %s value cannot be less than 1", strconv.Itoa(index)))
		}
		if each > 1000 {
			return errors.New(fmt.Sprintf("opponent_proficiencies array of index %s value cannot be greater than 1000", strconv.Itoa(index)))
		}
	}

	for index, each := range competitionData.OpponentExps {
		if each < 1 {
			return errors.New(fmt.Sprintf("opponent_exps array of index %s value cannot be less than 1", strconv.Itoa(index)))
		}
		if each > 1000 {
			return errors.New(fmt.Sprintf("opponent_exps array of index %s value cannot be greater than 1000", strconv.Itoa(index)))
		}
	}

	return nil
}
