package dcp

import "math"

/*

Problem:
You are given a 2-d matrix where each cell represents number of coins in that cell.
Assuming we start at matrix[0][0], and can only move right or down, find the maximum
number of coins you can collect by the bottom right corner.

For example, in this matrix

	0 3 1 1
	2 0 0 4
	1 5 3 1

The most we can collect is 0 + 2 + 1 + 5 + 3 + 1 = 12 coins.

*/

// Algorítimo usado: Busca em largura - Breadth-First Search - BFS (https://pt.wikipedia.org/wiki/Busca_em_largura)

func bfs(matrix [][]int, i, j, n, m, sum int) int {
	/*
		matrix = matriz
		i, j   = coordenadas da posição atual
		n, m   = tamanho da matriz (n externa | m interna)
		sum    = soma das moedas (cumulativa)

	*/

	// Chegamos ao final, retonar o valor da última posição (canto inferior direito) + o cumulativo (sum)
	if i == n-1 && j == m-1 {
		return sum + matrix[i][j]
	}

	// Não podemos ir para baixo, retorna para direita, soma cumulativo + posição atual
	if i == n-1 {
		return bfs(matrix, i, j+1, n, m, sum+matrix[i][j])
	}

	// Não podemos mais ir para direita, retorna para baixo, soma cumulativo + posição atual
	if j == m-1 {
		return bfs(matrix, i+1, j, n, m, sum+matrix[i][j])
	}

	// Se ainda podemos ir pra cima ou para baixo, retorna o maior valor (calculando indo para cima ou para baixo)
	return int(math.Max(float64(bfs(matrix, i+1, j, n, m, sum+matrix[i][j])), float64(bfs(matrix, i, j+1, n, m, sum+matrix[i][j]))))
}

func getMaxCoins(matrix [][]int) int {
	// começamos sempre na posição 0 0
	// sum sempre vai iniciar com 0
	return bfs(matrix, 0, 0, len(matrix), len(matrix[0]), 0)
}
