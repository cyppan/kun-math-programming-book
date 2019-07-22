package secrets

import (
	"testing"
)

func TestSecrets(t *testing.T) {
	cases := []struct {
		secret, k, n int
	}{
		{42, 3, 5},
		{12, 2, 2},
	}
	for _, c := range cases {
		points := GenerateGroupKeys(c.secret, c.k, c.n)
		for j := c.k; j <= c.n; j++ {
			secret := DecodeSecret(points[j-c.k : j])
			if secret != c.secret {
				t.Errorf("decoded secret == %v, want %v", secret, c.secret)
			}
		}
	}
}
