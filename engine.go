package catngine

import (
	"errors"
	"fmt"
)

const (
	// E representa el estado de una casilla vacia
	E = int8(0)

	// P representa el estado de una casilla de la cual se ha apoderado el jugador humano
	P = int8(1)

	// F representa el estado de una casilla de la cual se ha apoderado la máquina
	F = int8(2)
)

var (
	// ErrOutOfMinimaxBounds es el error que se gatilla cuando se intenta acceder a una casilla fuera de los límites del mundo del juego
	ErrOutOfMinimaxBounds = errors.New("out of Minimax bounds")

	// ErrNoEmptyCell se gatilla cuando se intenta adueñarse de una casilla que ya tiene dueño
	ErrNoEmptyCell = errors.New("that cell is not empty")

	MinimaxWinningStates = [][]int8{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
)

var _ Catngine = (*Minimax)(nil)

type Catngine interface {
	Evaluate(p int8) int8
	Set(x, y int8, p int8) error
	String() string
	Winner(p int8) bool
}

type Minimax struct {
	turn int8
	g    []int8
}

func NewMinimax() *Minimax {
	return &Minimax{g: make([]int8, 9), turn: 1}
}

func (m *Minimax) Turn() int8 {
	return m.turn / int8(2)
}

func (b *Minimax) m(x, y int8) (int8, error) {
	i := y + 3*x

	if i < 0 || i > 8 {
		return 0, ErrOutOfMinimaxBounds
	}
	return i, nil
}

func (b *Minimax) Set(x, y int8, p int8) error {
	i, err := b.m(x, y)

	if err != nil {
		return err
	}

	if b.g[i] != E {
		return ErrNoEmptyCell
	}

	b.g[i] = p

	b.turn++

	return nil
}

func (b *Minimax) SetI8(i int8, p int8) error {
	if i < 0 || i > 8 {
		return ErrOutOfMinimaxBounds
	}

	if b.g[i] != E {
		return ErrNoEmptyCell
	}

	b.g[i] = p

	return nil
}

func (b *Minimax) Evaluate(p int8) int8 {
	j := int8(0)

	score := int8(20)
	if p == P {
		score = -20
	}

	for i := int8(0); i < 9; i++ {
		if b.g[i] == E {
			b.g[i] = p

			if p == P {
				scoreMM := miniMax(b, 0, F)
				b.g[i] = E
				if scoreMM > score {
					score = scoreMM
					j = i
				}
			}

			if p == F {
				scoreMM := miniMax(b, 0, P)
				b.g[i] = E
				if scoreMM < score {
					score = scoreMM
					j = i
				}
			}
		}
	}

	return j
}

func (b *Minimax) String() string {
	return fmt.Sprintf(" %v | %v | %v \n---+---+---\n %v | %v | %v \n---+---+---\n %v | %v | %v \n", b.g[0], b.g[1], b.g[2], b.g[3], b.g[4], b.g[5], b.g[6], b.g[7], b.g[8])
}

func (b *Minimax) Winner(p int8) bool {
	return checkWin(b, p)
}

func checkWin(b *Minimax, p int8) bool {
	return (b.g[0] == p && b.g[1] == p && b.g[2] == p) ||
		(b.g[3] == p && b.g[4] == p && b.g[5] == p) ||
		(b.g[6] == p && b.g[7] == p && b.g[8] == p) ||
		(b.g[0] == p && b.g[3] == p && b.g[6] == p) ||
		(b.g[1] == p && b.g[4] == p && b.g[7] == p) ||
		(b.g[2] == p && b.g[5] == p && b.g[8] == p) ||
		(b.g[0] == p && b.g[4] == p && b.g[8] == p) ||
		(b.g[2] == p && b.g[4] == p && b.g[6] == p)
}

type States []*Minimax

func miniMax(b *Minimax, depth int8, p int8) int8 {
	mark := int8(20)

	if checkWin(b, P) {
		return 11 - depth
	}

	if checkWin(b, F) {
		return depth - 11
	}

	if p == P {
		mark = -20
	}

	maxTrack := 8

	if b.Turn() <= 1 {
		maxTrack = 2
	}

	for i := 0; i <= maxTrack; i++ {
		if b.g[i] == E {
			if p == P {
				b.g[i] = P
				m2 := miniMax(b, depth+1, F)
				if mark < m2 {
					mark = m2
				}
			}

			if p == F {
				b.g[i] = F
				m2 := miniMax(b, depth+1, P)
				if mark > m2 {
					mark = m2
				}
			}
			b.g[i] = E
		}
	}

	return mark
}
