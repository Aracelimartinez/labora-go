package math

import (
	"testing"
	"fmt"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Add(4,6)
	}
}

func ExampleAdd() {
    fmt.Println(Add(4, 6))
    // Output: 10
}

type factorialTest struct{
	number, expected int
}

var factorialTests = []factorialTest{
	{0, 1},
	{2, 2},
	{3, 6},
	{4, 24},
	{5, 120},
}

func TestFactorial(t *testing.T){
	for _, test := range factorialTests{
		output := Factorial(test.number)
		if output != test.expected {
			t.Errorf("The result: %q is not equal to the expected: %q", output, test.expected)
		}
	}
}
