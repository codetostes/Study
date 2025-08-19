package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A Área da forma é %0.3f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type triangulo struct {
	base   float64
	altura float64
}

func (t triangulo) area() float64 {
	return (t.base * t.altura) / 2
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	t := triangulo{10, 15}
	escreverArea(t)
}

// Interfaces são muito utilizadas em GO quando se precisa ter uma certa flexibilidade com os tipos
