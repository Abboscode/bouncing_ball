package main

import (
	"fmt"
	"math"
	"time"
)

// initialize array
func initArray(ar *[][]string) {
	const WINDDOW_WIDTH = 100
	const WINDOW_HEIGHT = 60
	*ar = make([][]string, WINDOW_HEIGHT)
	for i := 0; i < WINDOW_HEIGHT; i++ {
		(*ar)[i] = make([]string, WINDDOW_WIDTH)

		for j := 0; j < WINDDOW_WIDTH; j++ {

			(*ar)[i][j] = " "

		}

	}
}

func printArray(arr [][]string) {

	for _, strings := range arr {

		fmt.Print(strings)
		fmt.Println()

	}

}

func drawCircle(radius int, x int, y int, window *[][]string) {
	h := cap(*window)
	width := cap((*window)[0])

	//xy := math.Round(math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2)))
	//radius to draw from
	//	radiusxy := float64(radius) + xy

	for i := 0; i < 180; i++ {
		xx := int(math.Round(float64(radius) * math.Cos(float64(i))))
		yy := int(math.Round(float64(radius) * math.Sin(float64(i))))
		xxx := int(x + xx)

		yyy := int(y + yy)
		if yyy < width && yyy > 0 && xxx < h && xxx > 0 {
			(*window)[xxx][yyy] = "*"
		}
	}

}

func main() {

	const ac = 10

	var window = [][]string{}
	//initiaalize array

	//printArray

	for i := 0; i < 80; i++ {
		initArray(&window)
		//printArray(window)
		v0 := ac * time.Duration(i) * 8000000
		t := v0 / ac
		drawCircle(2, i, 4, &window)

		printArray(window)

		time.Sleep((time.Second / 2) - t)

	}

}
