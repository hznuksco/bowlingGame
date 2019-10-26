package bowling

import "testing"

func rollMany(n, pins int, g *Game) {
	for i := 0; i < n; i++ {
		g.Roll(pins)
	}
}

func rollSpare(g *Game) {
	g.Roll(5)
	g.Roll(5)
}

func rollStrike(g *Game) {
	g.Roll(10)
}

func TestGame_Score(t *testing.T) {
	tests := []struct {
		name   string
		arg    []int
		want   int
	}{
		{"TestAllZero", []int{20, 0}, 0},
		{"TestAllOne", []int{20, 1}, 20},
		{"TestPerfectGame", []int{12, 10}, 300},
	}
	for _, tt := range tests {
		g := NewGame()
		t.Run(tt.name, func(t *testing.T) {
			rollMany(tt.arg[0], tt.arg[1], g)
			if got := g.Score(); got != tt.want {
				t.Errorf("Score() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("TestSpare", func(t *testing.T) {
		want := 16
		g := NewGame()
		rollSpare(g) //spare
		g.Roll(3)
		rollMany(17, 0, g)

		if got := g.Score(); got != want {
			t.Errorf("Score() = %v, want %v", got, want)
		}
	})

	t.Run("TestStrick", func(t *testing.T) {
		want := 24
		g := NewGame()
		rollStrike(g) //strick
		g.Roll(3)
		g.Roll(4)
		rollMany(16, 0, g)

		if got := g.Score(); got != want {
			t.Errorf("Score() = %v, want %v", got, want)
		}
	})

}
