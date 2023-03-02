package mathint

import "math"

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Abs(a int) int {
	return int(math.Abs(float64(a)))
}

func Ceil(a int) int {
	return int(math.Ceil(float64(a)))
}

func Floor(a int) int {
	return int(math.Floor(float64(a)))
}

func Round(a int) int {
	return int(math.Round(float64(a)))
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}

func Mod(a, b int) int {
	return int(math.Mod(float64(a), float64(b)))
}
