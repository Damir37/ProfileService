package unit

import (
	"ProfileService/internal/pkg/elo"
	"fmt"
	"math"
	"testing"
)

func TestCalcELO(t *testing.T) {
	tests := []struct {
		name           string
		gamesPlayed    int
		wins           int
		draws          int
		losses         int
		currentRating  int
		opponentRating int
		expectedRating int
	}{
		{
			name:           "Новый игрок, выиграл 3, ничья 1, поражение 1",
			gamesPlayed:    5,
			wins:           3,
			draws:          1,
			losses:         1,
			currentRating:  1500,
			opponentRating: 1600,
			expectedRating: 1568,
		},
		{
			name:           "Опытный игрок, 7 побед, 3 ничьи",
			gamesPlayed:    50,
			wins:           7,
			draws:          3,
			losses:         0,
			currentRating:  1800,
			opponentRating: 1700,
			expectedRating: 1842,
		},
		{
			name:           "Игрок проигрывает все игры",
			gamesPlayed:    5,
			wins:           0,
			draws:          0,
			losses:         5,
			currentRating:  1500,
			opponentRating: 1600,
			expectedRating: 1428,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRating := elo.CalculatorELO(tt.gamesPlayed, tt.wins, tt.draws, tt.losses, tt.currentRating, tt.opponentRating)

			fmt.Printf("Тест: %s\n", tt.name)
			fmt.Printf("Игр сыграно: %d\n", tt.gamesPlayed)
			fmt.Printf("Побед: %d, Ничьих: %d, Поражений: %d\n", tt.wins, tt.draws, tt.losses)
			fmt.Printf("Текущий рейтинг: %d\n", tt.currentRating)
			fmt.Printf("Рейтинг соперника: %d\n", tt.opponentRating)
			fmt.Printf("Ожидаемый рейтинг: %d\n", tt.expectedRating)
			fmt.Printf("Фактический рейтинг: %d\n", actualRating)
			fmt.Println("----------------------------------------")

			if math.Abs(float64(actualRating-tt.expectedRating)) > 1 {
				t.Errorf("ожидался рейтинг %d, но получили %d", tt.expectedRating, actualRating)
			}
		})
	}
}
