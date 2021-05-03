package main

import "fmt"

type Tabuleiro [][]int8

func main() {

	tabuleiros := explorar(tabuleiroInicial(7), make([]Tabuleiro, 0))

	fmt.Printf("Quantidade %d\n", len(tabuleiros))

	for _, tabuleiro := range tabuleiros {
		for _, linha := range tabuleiro {
			for _, valor := range linha {
				var output string
				if valor == 1 {
					output = "Q"
				} else {
					output = "-"
				}
				fmt.Printf(output + " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func tabuleiroInicial(n int) Tabuleiro {
	tabuleiro := make(Tabuleiro, n)
	for i := 0; i < n; i++ {
		tabuleiro[i] = make([]int8, n)
	}
	return tabuleiro
}

func tabuleiroFilho(b Tabuleiro, x int, y int) Tabuleiro {
	length := len(b)
	filho := tabuleiroInicial(length)
	for y, row := range b {
		for x, val := range row {
			filho[y][x] = val
		}
	}
	y1 := y - x
	y2 := y + x
	for i := 0; i < length; i++ {
		filho[y][i] = 2
		filho[i][x] = 2
		if y1 >= 0 && y1 < length {
			filho[y1][i] = 2
		}
		if y2 >= 0 && y2 < length {
			filho[y2][i] = 2
		}
		y1 = y1 + 1
		y2 = y2 - 1
	}
	filho[y][x] = 1
	return filho
}

func proximo(tabuleiro Tabuleiro) []Tabuleiro {
	children := make([]Tabuleiro, 0)
	achou := false
	for y, row := range tabuleiro {
		for x, val := range row {
			if val == 0 {
				achou = true
				children = append(children, tabuleiroFilho(tabuleiro, x, y))
			}
		}
		if achou {
			break
		}
	}
	return children
}

func criterion(tabuleiro Tabuleiro) bool {
	for _, row := range tabuleiro {
		hasQueen := false
		for _, val := range row {
			if val == 1 {
				hasQueen = true
				break
			}
		}
		if !hasQueen {
			return false
		}
	}
	return true
}

func explorar(tabuleiro Tabuleiro, boards []Tabuleiro) []Tabuleiro {
	for _, successor := range proximo(tabuleiro) {
		if criterion(successor) {
			boards = append(boards, successor)
		} else {
			boards = explorar(successor, boards)
		}
	}
	return boards
}
