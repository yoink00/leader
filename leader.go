package main

import "math"

var peers []string

func getQuorum(p []string) int {
	if len(p) == 0 {
		return 0
	}

	half := math.Ceil(float64(len(p)) / 2.0)
	if len(p)&1 == 0 {
		half++
	}

	return int(half)
}

func main() {

}
