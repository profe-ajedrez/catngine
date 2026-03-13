// Package main muestra un ejemplo de uso de catngine:
// el jugador (X) vs la máquina (O) en una partida de Tic-Tac-Toe por consola.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/profe-ajedrez/catngine"
)

func main() {
	b := catngine.NewMinimax()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Tic-Tac-Toe: tú (X) vs máquina (O)")
	fmt.Println("Ingresá las coordenadas como: x y  (columna fila, valores 0-2)")
	fmt.Println()

	for {
		fmt.Println(b.String())

		// Turno del jugador
		fmt.Print("Tu jugada (x y): ")
		scanner.Scan()
		parts := strings.Fields(scanner.Text())

		if len(parts) != 2 {
			fmt.Println("Ingresá dos números separados por espacio.")
			continue
		}

		x, errX := strconv.ParseInt(parts[0], 10, 8)
		y, errY := strconv.ParseInt(parts[1], 10, 8)

		if errX != nil || errY != nil {
			fmt.Println("Valores inválidos.")
			continue
		}

		if err := b.Set(int8(x), int8(y), catngine.P); err != nil {
			fmt.Printf("Jugada inválida: %v\n", err)
			continue
		}

		if b.Winner(catngine.P) {
			fmt.Println(b.String())
			fmt.Println("¡Ganaste!")
			return
		}

		if full(b) {
			fmt.Println(b.String())
			fmt.Println("Empate.")
			return
		}

		// Turno de la máquina
		move := b.Evaluate(catngine.F)
		_ = b.SetIndex(move, catngine.F)
		fmt.Printf("La máquina juega en la casilla %d\n\n", move)

		if b.Winner(catngine.F) {
			fmt.Println(b.String())
			fmt.Println("Ganó la máquina.")
			return
		}

		if full(b) {
			fmt.Println(b.String())
			fmt.Println("Empate.")
			return
		}
	}
}

// full devuelve true si no quedan casillas vacías.
func full(b *catngine.Minimax) bool {
	for _, v := range b.Board() {
		if v == catngine.E {
			return false
		}
	}
	return true
}
