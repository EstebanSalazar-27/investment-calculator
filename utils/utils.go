package utils

import "math"

func expotenialRateFormForYears(anualReturn float64, years float64) float64 {
	var result = math.Pow(1+anualReturn/100, years)
	return result
}
