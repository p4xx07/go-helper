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

func Ceil(a float64) int {
	return int(math.Ceil(a))
}

func Floor(a float64) int {
	return int(math.Floor(a))
}

func Round(a float64) int {
	return int(math.Round(a))
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
