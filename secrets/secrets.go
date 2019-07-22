package secrets

import (
	"math/rand"

	"github.com/cyppan/kun-math-programming-book/polynomial"
)

// Point is a type alias
type Point = polynomial.Point

// GenerateGroupKeys ...
func GenerateGroupKeys(secret, k, n int) []Point {
	points := make([]Point, k)
	points[0] = Point{0, float64(secret)}
	for i := 1; i < k; i++ {
		points[i] = Point{float64(i), float64(rand.Intn(999999999))}
	}
	p := polynomial.InterpolatePolynomial(points)
	points = append(points[1:k], Point{float64(k), p.Y(float64(k))})
	if n > k {
		for j := k + 1; j <= n; j++ {
			points = append(points, Point{float64(j), p.Y(float64(j))})
		}
	}
	return points
}

// DecodeSecret ...
func DecodeSecret(points []Point) int {
	p := polynomial.InterpolatePolynomial(points)
	return int(p.Y(0))
}
