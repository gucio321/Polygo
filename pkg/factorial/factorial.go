// Package factorial contains Factorial method
package factorial

func Factorial(n int) (result int) {
	result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
