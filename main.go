package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	Width  = 4
	Height = 4
)

type Graph map[string]map[string]int

func keyString(num1 int, num2 int) string {
	key := strconv.Itoa(num1) + "," + strconv.Itoa(num2)
	return key
}

func add(m Graph, path, target string) {
	mm, exist := m[path]
	if !exist {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[target] = 1
}

func generateGraphMap(grid [][]bool) Graph {
	m := make(map[string]map[string]int, Height)
	for i, _ := range grid {
		for j, _ := range grid[i] {
			if j == len(grid[i])-1 {
				fmt.Printf("%v, %v: reached row end\n", i, j)
				break
			}
			if grid[i][j+1] == false {
				fmt.Printf("%v, %v: valid space ahead\n", i, j)
				key := keyString(i, j)
				target := keyString(i, j+1)
				add(m, key, target)

			} else if grid[i][j+1] == true {
				fmt.Printf("%v, %v: wall encountered ahead\n", i, j)
			}
			if i < len(grid)-1 {
				if grid[i+1][j] == false {
					fmt.Printf("%v, %v: valid space below\n", i, j)
					key := keyString(i, j)
					target := keyString(i+1, j)
					add(m, key, target)
				}
			}

		}
	}
	return m
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
	graphMap := generateGraphMap(mapGrid)
	for i, v := range graphMap {
		fmt.Println(i, v)
	}
}
