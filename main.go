package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	Width  = 4
	Height = 4
)


type Graph map[string]map[string]int

func generateGraphMap(grid [][]bool) {
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if j == len(grid[i])-1 {
				fmt.Printf("%v, %v: reached row end\n",i,j)
				break
			}
			if grid[i][j+1] == true {
				fmt.Printf("%v, %v: wall encountered ahead\n",i,j)
				break
			} 
			if i < len(grid)-1 {
				if grid[i+1][j] == true {
					fmt.Printf("%v, %v: wall encountered below\n",i,j)
				}
			} 

		}
	}
}

func generateMap() [][]bool {
	rand.Seed(time.Now().UnixNano())

	// Create an empty map
	grid := make([][]bool, Height)
	for i := range grid {
		grid[i] = make([]bool, Width)
	}

	// Add random walls
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			// Adjust the probability to control the density of walls
			if rand.Float64() < 0.2 {
				grid[y][x] = true
			}
		}
	}

	return grid
}

func printMap(grid [][]bool) {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if grid[y][x] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	mapGrid := generateMap()
	printMap(mapGrid)
	generateGraphMap(mapGrid)
	fmt.Println(mapGrid)
}
