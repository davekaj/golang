package methodsandinterface

import "fmt"

type errNegativeSqrt float64

func (e errNegativeSqrt) Error() string { //*?
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

//Sqrt needs a comment
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errNegativeSqrt(x)
	}
	z := x / 4
	tempZ := 0.0
	for i := 0; i < 11; i++ {
		z -= (z*z - x) / (2 * z)
		if (tempZ-z) < 0.000001 && i != 0 {
			fmt.Println("x =", x, "guesses: ", i)
			i = 10
		}
		tempZ = z
	}
	return z, nil
}

//ErrorExcersize needs a comment
func ErrorExcersize() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
