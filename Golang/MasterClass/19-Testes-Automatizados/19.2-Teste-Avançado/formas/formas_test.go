package formas_test

import (
	"math"
	. "teste-avancado/formas"
	"testing"
)

func TestArea(t *testing.T) {
	r := Rectangle{10, 12}
	t.Run("Retangulo", func(t *testing.T) {
		if r.Area() != float64(120) {
			t.Errorf("A área recebida %f é diferente da esperada 120", r.Area())

		}
	})
	c := Circle{10}
	t.Run("Curculo", func(t *testing.T) {
		if c.Area() != float64(math.Pi*100) {
			t.Fatalf("A área recebida %f é diferente da esperada 120", c.Area())

		}
	})
	tr := Triangle{10, 12}
	t.Run("Triangulo", func(t *testing.T) {
		if tr.Area() != float64(60) {
			t.Fatalf("A área recebida %f é diferente da esperada 120", c.Area())

		}
	})
}
