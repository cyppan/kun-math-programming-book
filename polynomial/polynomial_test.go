package polynomial

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddPolynomial(t *testing.T) {
	cases := []struct {
		p1, p2, result Polynomial
	}{
		{Polynomial([]float64{}), Polynomial([]float64{}), Polynomial([]float64{})},
		{Polynomial([]float64{}), Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{1.0, 2.0})},
		{Polynomial([]float64{2.5, 1.5}), Polynomial([]float64{1.0}), Polynomial([]float64{3.5, 1.5})},
		{Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{2.0, 4.0})},
	}
	for _, c := range cases {
		original := fmt.Sprint(c.p1)
		c.p1.Add(c.p2)
		if len(c.p1) != len(c.result) || !reflect.DeepEqual(c.p1, c.result) {
			t.Errorf("%q + %q == %q, want %q", original, c.p2, c.p1, c.result)
		}
	}
}

func TestMultiplyPolynomial(t *testing.T) {
	cases := []struct {
		p1, p2, result Polynomial
	}{
		{Polynomial([]float64{}), Polynomial([]float64{}), Polynomial([]float64{})},
		{Polynomial([]float64{}), Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{})},
		{Polynomial([]float64{1}), Polynomial([]float64{1, 2}), Polynomial([]float64{1, 2})},
		{Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{1.0, 2.0}), Polynomial([]float64{1, 4, 4})},
	}
	for _, c := range cases {
		original := fmt.Sprint(c.p1)
		c.p1.Multiply(c.p2)
		if len(c.p1) != len(c.result) || !reflect.DeepEqual(c.p1, c.result) {
			t.Errorf("(%q) * (%q) == %q, want %q", original, c.p2, c.p1, c.result)
		}
	}
}

func TestInterpolate(t *testing.T) {
	cases := []struct {
		points []Point
		p      Polynomial
	}{
		{[]Point{Point{1, 1}}, Polynomial([]float64{1})},
		{[]Point{Point{1, 1}, Point{2, 0}}, Polynomial([]float64{2, -1})},
		// {[]Point{Point{1, 1}, Point{2, 4}, Point{7, 9}}, Polynomial([]float64{-2.666666, 3.999999, -0.3333333})},
	}
	for _, c := range cases {
		p := InterpolatePolynomial(c.points)
		if len(p) != len(c.p) || !reflect.DeepEqual(p, c.p) {
			t.Errorf("InterpolatePolynomial(%v) == %q, want %q", c.points, p, c.p)
		}
	}
}
