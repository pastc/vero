package internal

import "math"

//func Factorial(n int) int {
//	if n == 0 || n == 1 {
//		return 1
//	}
//
//	return n * Factorial(n-1)
//}

// Binomial
//
// n = Row number
//
// r = Column number
func Binomial(n, k int) int {
	if n < 0 || k < 0 {
		return 0
	}
	if n < k {
		return 0
	}
	// (n,k) = (n, n-k)
	if k > n/2 {
		k = n - k
	}
	b := 1
	for i := 1; i <= k; i++ {
		b = (n - k + i) * b / i
	}
	return b
}

// BinomialDistribution calculates the binomial distribution probability.
func BinomialDistribution(n, r int) float64 {
	return float64(Binomial(n, r)) * math.Pow(0.5, float64(n)) * 100
}

//// BinomialPascal calculates the binomial distribution using Pascal's triangle. Very fast
//func BinomialPascal(n, r int) float64 {
//	m := 0
//	for i := 0; i <= n; i++ {
//		m += Binomial(n, i)
//	}
//	return float64(Binomial(n, r)) / float64(m) * 100
//}
