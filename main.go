package goarea

import "math"

const Pi = 3.1416

func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Retangulo(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2.0
}
