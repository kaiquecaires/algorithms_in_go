package tournamentwinner

const HOME_TEAM_WON = 1

type TeamWithHighestScore struct {
	Team  string
	Score int
}

func TournamentWinner(competitions [][]string, results []int) string {
	scores := map[string]int{}
	teamWithHighestScore := &TeamWithHighestScore{}

	for i, competition := range competitions {
		winnerIdx := results[i] ^ 1

		if _, ok := scores[competition[winnerIdx]]; ok {
			scores[competition[winnerIdx]] += 3
		} else {
			scores[competition[winnerIdx]] = 3
		}

		if teamWithHighestScore.Score < scores[competition[winnerIdx]] {
			teamWithHighestScore.Team = competition[winnerIdx]
			teamWithHighestScore.Score = scores[competition[winnerIdx]]
		}
	}

	return teamWithHighestScore.Team
}
