package elo

import "math"

func CalculatorELO(gamesPlayed, wins, draws, losses, currentRating, opponentRating int) int {
	var K int
	if gamesPlayed < 30 {
		K = 40
	} else {
		K = 20
	}

	totalGames := wins + draws + losses
	if totalGames == 0 {
		return currentRating
	}

	totalScore := float64(wins) + float64(draws)*0.5

	EA := 1 / (1 + math.Pow(10, float64(opponentRating-currentRating)/400))

	expectedScore := float64(totalGames) * EA

	newRating := float64(currentRating) + float64(K)*(totalScore-expectedScore)

	return int(math.Round(newRating))
}
