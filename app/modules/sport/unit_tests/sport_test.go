package unit_tests

import (
	"fmt"
	sportsResources "sports-competition/app/modules/sport/resources"
	sportServices "sports-competition/app/modules/sport/services"
	"strconv"
	"testing"
)

func TestSports(t *testing.T) {
	TestSportCalculateMaxProficiency(t)
}

func TestSportCalculateMaxProficiency(t *testing.T) {
	t.Logf("Testing CalculateMaxProficiency() function.")
	var sportsService *sportServices.SportService = sportServices.NewSportService()
	var maxProficiency = sportsService.CalculateMaxProficiency(&sportsResources.SportBeginCompetition{
		BaseSkill:             []uint{5, 9},
		OpponentProficiencies: []uint{2, 3, 6, 7, 8},
		OpponentExps:          []uint{3, 4, 2, 2, 3},
	})
	if maxProficiency != 23 {
		t.Fatalf(fmt.Sprintf("CalculateMaxProficiency() expected to return uint value of 23 but returned %s instead", strconv.Itoa(int(maxProficiency))))
	}
}
