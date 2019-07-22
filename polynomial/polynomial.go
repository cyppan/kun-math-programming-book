package polynomial

import (
	"fmt"
	"math"
	"strings"
)

// Polynomial is defined given its ordered coefficients
type Polynomial []float64

func (p Polynomial) String() string {
	var sb strings.Builder
	for i, a := range p {
		sb.WriteString(fmt.Sprint(a))
		if i > 0 {
			sb.WriteString(fmt.Sprintf("x^%v", i))
		}
		if i < len(p)-1 {
			sb.WriteString(" + ")
		}
	}
	return sb.String()
}

// Y computes the polynomial value at x
func (p Polynomial) Y(x float64) (y float64) {
	for i, coef := range p {
		y += coef * math.Pow(x, float64(i))
	}
	return
}

// Add two polynomials
func (p *Polynomial) Add(p2 Polynomial) {
	if len(p2) == 0 {
		return
	}
	var coefs Polynomial
	if len(p2) > len(*p) {
		coefs = make(Polynomial, len(p2))
		copy(coefs, *p)
	} else {
		coefs = *p
	}
	for i := range coefs {
		if len(p2) >= i+1 {
			coefs[i] += p2[i]
		}
	}
	*p = coefs
}

// Multiply two polynomials
func (p *Polynomial) Multiply(p2 Polynomial) {
	// given two polynoms degrees n1 and n2, the product is of degree n1 + n2
	if len(*p) == 0 || len(p2) == 0 {
		*p = make(Polynomial, 0)
		return
	}
	n := len(*p) + len(p2) - 1
	result := make(Polynomial, n)
	for i := range result {
		for j := 0; j <= i; j++ {
			k := i - j
			if j <= len(*p)-1 && k <= len(p2)-1 {
				result[i] += (*p)[j] * p2[k]
			}
		}
	}
	*p = result
}

// A Point is a simple X, Y coordinates container
type Point struct{ X, Y float64 }

// InterpolatePolynomial interpolates the n-polynomial corresponding to the given n+1 points
func InterpolatePolynomial(points []Point) Polynomial {
	if len(points) == 0 {
		panic("InterpolatePolynomial takes at least one point as its input")
	}
	p := Polynomial([]float64{})
	for i := 0; i < len(points); i++ {
		product := Polynomial([]float64{1})
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			product.Multiply(Polynomial([]float64{-points[j].X / (points[i].X - points[j].X), 1 / (points[i].X - points[j].X)}))
		}
		product.Multiply(Polynomial([]float64{points[i].Y}))
		p.Add(product)
	}
	return p
}
