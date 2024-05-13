package resources

type SportBeginCompetition struct {
	BaseSkill             []uint `json:"base_skill"`
	OpponentProficiencies []uint `json:"opponent_proficiencies"`
	OpponentExps          []uint `json:"opponent_exps"`
}

func (data *SportBeginCompetition) GetDescendingOpponentList() []SportBeginCompetitionEachOpponentData {
	var opponentList []SportBeginCompetitionEachOpponentData
	for i := 0; i < len(data.OpponentProficiencies); i++ {
		opponentList = append(opponentList, SportBeginCompetitionEachOpponentData{
			Proficiency: data.OpponentProficiencies[i],
			Exp:         data.OpponentExps[i],
		})
	}
	for i := 0; i < len(opponentList); i++ {
		for j := i + 1; j < len(opponentList); j++ {
			if opponentList[j].Proficiency < opponentList[i].Proficiency {
				opponentList[i], opponentList[j] = opponentList[j], opponentList[i]
			}
		}
	}

	return opponentList
}

type SportBeginCompetitionEachOpponentData struct {
	Proficiency uint `json:"skill_level"`
	Exp         uint `json:"exp"`
}

type SportBeginCompetitionResult struct {
	MaxProficiency uint `json:"max_proficiency"`
}
