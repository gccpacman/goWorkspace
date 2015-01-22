/* Exercise: Loops and Functions #43 */
package main

import (
	"fmt"
	"math"
)

func Abs(x float64) float64 {
	if x < 0 {
		x = -x
	}
	return x
}

func Sqrt(x float64) (float64, int) {

	z := float64(1)
	tmp := float64(0)
	times := int(0)

	for {
		times++
		z = z - (z*z-x)/(2*z)
		if Abs(z-tmp) < 1e-15 {
			//fmt.Printf("loop %d times\n", times)
			break
		}
		tmp = z
	}

	return z, times
}

func main() {

	for i := float64(1); i < 10; i++ {
		attempt, times := Sqrt(i)

		expected := math.Sqrt(i)

		fmt.Printf("sqrt %g times %d --- attempt= %g (expected = %g) error = %g\n", i, times, attempt, expected, Abs(attempt-expected))
	}

}
