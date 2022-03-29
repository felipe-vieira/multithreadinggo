package main

import (
	"fmt"
	"math/rand"
	"time"
)

const matrixSize = 250

var (
	matrixA      = [matrixSize][matrixSize]int{}
	matrixB      = [matrixSize][matrixSize]int{}
	result       = [matrixSize][matrixSize]int{}
	workStart    = NewBarrier(matrixSize + 1)
	workComplete = NewBarrier(matrixSize + 1)
)

func main() {
	fmt.Println("working")
	start := time.Now()

	for row := 0; row < matrixSize; row++ {
		go workOutRow(row)
	}

	for i := 0; i < 100; i++ {
		generateRandomMatrix(&matrixA)
		generateRandomMatrix(&matrixB)
		workStart.Wait()
		workComplete.Wait()
	}

	elapsed := time.Since(start)
	fmt.Printf("elapsed : %s \n", elapsed)
}

func generateRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := 0; row < matrixSize; row++ {
		for col := 0; col < matrixSize; col++ {
			matrix[row][col] += rand.Intn(10) - 5
		}
	}
}

func workOutRow(row int) {
	for {
		workStart.Wait()
		for col := 0; col < matrixSize; col++ {
			for i := -0; i < matrixSize; i++ {
				result[row][col] += matrixA[row][i] * matrixB[i][col]
			}
		}
		workComplete.Wait()
	}
}
