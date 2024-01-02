package main

import (
	"fmt"
	"math"
)

func expotenialRateFormForYears(anualReturn float64, years float64) float64 {
	var result = math.Pow(1+anualReturn/100, years)
	return result
}

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var anualReturnRate float64
	var investmentYears float64
	fmt.Print("Type ur investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print(("Type anual rate: "))
	fmt.Scan(&anualReturnRate)
	fmt.Print(("Type investment years duration: "))
	fmt.Scan(&investmentYears)
	var returnValue float64 = investmentAmount * expotenialRateFormForYears(anualReturnRate, investmentYears)
	var realProfitWithInflation = returnValue / expotenialRateFormForYears(inflationRate, investmentYears)

	fmt.Println("Return for year is gonna be ", returnValue)
	fmt.Print("Final return later than 10 years of investment is ", realProfitWithInflation)
}
