package catngine

import (
	"errors"
	"fmt"
)

const (
	// E Empty. representa el estado de una casilla vacía
	E = int8(0)

	// P Player. representa el estado de una casilla de la cual se ha apoderado el jugador humano
	P = int8(1)

	// F Foe. representa el estado de una casilla de la cual se ha apoderado la máquina
	F = int8(2)
)

var (
	// ErrOutOfBoardBounds es el error que se gatilla cuando se intenta acceder a una casilla fuera de los límites del mundo del juego
	ErrOutOfBoardBounds = errors.New("out of board bounds")

	// ErrNoEmptyCell se gatilla cuando se intenta adueñarse de una casilla que ya tiene dueño
	ErrNoEmptyCell = errors.New("that cell is not empty")
)


type Minimax struct {
	g    []int8
}

func NewMinimax() *Minimax {
	return &Minimax{g: make([]int8, 9)}
}

func (b *Minimax) m(x, y int8) (int8, error) {
    if x < 0 || x > 2 || y < 0 || y > 2 {
        return 0, ErrOutOfBoardBounds
    }
    i := 3*y + x
    return i, nil
}

// Board expone el slice interno del tablero
func (m *Minimax) Board() []int8 {
 	return m.g
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


	return nil
}

func (b *Minimax) SetIndex(i int8, p int8) error {
	if i < 0 || i > 8 {
		return ErrOutOfBoardBounds
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
    s := [3]string{" ", "X", "O"}
    return fmt.Sprintf(
        " %s | %s | %s \n---+---+---\n %s | %s | %s \n---+---+---\n %s | %s | %s \n",
        s[b.g[0]], s[b.g[1]], s[b.g[2]],
        s[b.g[3]], s[b.g[4]], s[b.g[5]],
        s[b.g[6]], s[b.g[7]], s[b.g[8]],
    )
}

func (b *Minimax) Winner(p int8) bool {
	return checkWin(b, p)
}

func isDraw(b *Minimax) bool {
    for i := 0; i < 9; i++ {
        if b.g[i] == E {
            return false
        }
    }
    return !checkWin(b, P) && !checkWin(b, F)
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

func miniMax(b *Minimax, depth int8, p int8) int8 {
	mark := int8(20)

	if checkWin(b, P) {
		return 11 - depth
	}

	if checkWin(b, F) {
		return depth - 11
	}

	if isDraw(b) {
		return 0
	}

	if p == P {
		mark = -20
	}

	for i := 0; i <= 8; i++ {
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
