package generativerecursion

// GCD find the greatest common divisor of two numbers
// int [>= 1] int [>= 1] -> int [>= 1]
// Exmples: 2, 3 => 1; 6, 25 => 1; 18, 24 => 6
func GCD(a, b int) int {
	if a > b {
		return gcd(a, b)
	}
	return gcd(b, a)
}

func gcd(max, min int) int {
	if min == 0 {
		return max
	}

	return gcd(min, max%min)
}
