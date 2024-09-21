package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

func (r circulo) area() float64 {
	return math.Pi * math.Pow(r.raio, 2)
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

func main() {
	fmt.Println("Interfaces")

	r := retangulo{10, 15}

	escreverArea(r)

	c := circulo{10}

	escreverArea(c)
}
