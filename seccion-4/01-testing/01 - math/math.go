// math.go
package math

// Add nuestra función que suma dos enteros
func Add(x, y int) (res int) {
	return x + y
}

// Subtract resta dos enteros
func Subtract(x, y int) (res int) {
	return x - y
}

// Factorial calcula el factorial de un número.
func Factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * Factorial(n-1)
}
