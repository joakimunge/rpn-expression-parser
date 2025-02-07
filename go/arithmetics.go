package main

import "math"

type Arithmetic func(a float64, b float64) float64

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func modulo(x, y float64) float64 {
	return math.Mod(x, y)
}

func pow(x, y float64) float64 {
	return math.Pow(x, y)
}
