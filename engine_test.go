package catngine

import (
	"testing"
)

func TestMinimaxMap(t *testing.T) {
	tetsCases := []struct {
		x              int8
		y              int8
		msg            string
		expectedErr    error
		msgMapped      string
		expectedMapped int8
	}{
		{
			x:              9,
			y:              9,
			msg:            "MinimaxMap: expected ErrOutOfBoardBounds error. %v received",
			expectedErr:    ErrOutOfBoardBounds,
			msgMapped:      "",
			expectedMapped: 0,
		},
		{
			x:              1,
			y:              2,
			msg:            "MinimaxMap: expected no error. %v received",
			expectedErr:    nil,
			msgMapped:      "",
			expectedMapped: 7,
		},
		{
			x:              -1,
			y:              -2,
			msg:            "MinimaxMap: expected ErrOutOfBoardBounds error. %v received",
			expectedErr:    ErrOutOfBoardBounds,
			msgMapped:      "",
			expectedMapped: 0,
		},
		{x: -1, y: 1, msg: "MinimaxMap: expected ErrOutOfBoardBounds error. %v received", msgMapped:      "", expectedErr: ErrOutOfBoardBounds, expectedMapped: 0},

	}

	b := NewMinimax()

	for _, cs := range tetsCases {
		expected, err := b.m(cs.x, cs.y)

		if err != cs.expectedErr {
			t.Errorf(cs.msg, err)
			t.FailNow()
		}

		if expected != cs.expectedMapped {
			t.Errorf(cs.msgMapped, expected)
			t.FailNow()
		}
	}

}

func TestMinimaxWinner(t *testing.T) {
	b := NewMinimax()

	_ = b.Set(0, 0, P) // columna 0, fila 0 → índice 0
	_ = b.Set(0, 1, F) // columna 0, fila 1 → índice 3
	_ = b.Set(1, 0, P) // columna 1, fila 0 → índice 1
	_ = b.Set(1, 1, F) // columna 1, fila 1 → índice 4
	_ = b.Set(2, 0, P) // columna 2, fila 0 → índice 2

	if !b.Winner(P) {
		t.Errorf("MinimaxMap: expected true. %v received", false)
		t.FailNow()
	}

	if b.Winner(F) {
		t.Errorf("MinimaxMap: expected false. %v received", true)
		t.FailNow()
	}

}

func TestMinimaxEvaluate(t *testing.T) {
	testCase := []struct {
		Minimaxer func() *Minimax
		expected  int8
	}{
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(0, 0, P)
				_ = b.Set(1, 0, F)
				_ = b.Set(0, 1, P)

				return b
			},
			expected: 2,
		},
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(1, 1, P)
				_ = b.Set(1, 0, F)
				_ = b.Set(2, 0, P)

				return b
			},
			expected: 2,
		},
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(1, 1, P)
				_ = b.Set(1, 2, F)
				_ = b.Set(2, 1, P)

				return b
			},
			expected: 1,
		},
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(1, 1, F)
				_ = b.Set(1, 2, P)
				_ = b.Set(2, 1, F)
				_ = b.Set(0, 2, P)

				return b
			},
			expected: 1,
		},
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(2, 0, F)
				_ = b.Set(1, 2, P)
				_ = b.Set(2, 1, F)
				_ = b.Set(0, 2, P)

				return b
			},
			expected: 8,
		},
		{
			Minimaxer: func() *Minimax {
				b := NewMinimax()

				_ = b.Set(0, 2, F)
				_ = b.Set(2, 2, P)
				_ = b.Set(1, 1, F)
				_ = b.Set(1, 2, P)

				return b
			},
			expected: 6,
		},
	}

	for i, cs := range testCase {
		t.Logf("test case %v", i+1)
		b := cs.Minimaxer()

		//t.Log("before")
		//t.Logf("\n%v", b.String())
		nextMove := b.Evaluate(F)

		if nextMove != cs.expected {
			t.Logf("expected %v. %v received", cs.expected, nextMove)
			t.FailNow()
		}

		_ = b.SetI8(nextMove, F)

		//t.Log("after")
		//t.Logf("\n%v", b.String())
		//t.Log("------------------------------------")
	}
}

var mapResult int8
var winnerResult bool
var evalResult int8

func BenchmarkMinimaxMap(b *testing.B) {
    bd := NewMinimax()
    for i := 0; i < b.N; i++ { // era <=
        mapResult, _ = bd.m(1, 2)
    }
}

func BenchmarkMinimaxWinner(b *testing.B) {
    bd := NewMinimax()
    // ... setup con las coordenadas corregidas del Winner test ...
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        winnerResult = bd.Winner(P)
    }
}

func BenchmarkMinimaxEvaluate(b *testing.B) {
    // ... setup igual ...
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        evalResult = bd.Evaluate(F)
    }
}
