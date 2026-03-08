package main

import (
	"fmt"
	"math"
)

func Tittle(text string) {
	dashtitle := ""
	for i := 0; i < (78-len(text))/2; i++ {
		dashtitle += "-"
	}
	fmt.Println(dashtitle + text + dashtitle)
}

func main() {

	Tittle("Metric System BMI Calculator")

	//Constants for BMI calculation
	const bmiConstant = 703
	// Get user input for weight in kilograms
	fmt.Println("Enter your weight in kilograms:")
	var weightKg float64
	fmt.Scanln(&weightKg)

	// Get user input for height in centimeters
	fmt.Println("Enter your height in centimeters:")
	var heightCm float64
	fmt.Scanln(&heightCm)

	// Convert height from centimeters to meters
	heightM := heightCm / 100

	// Calculate BMI using the formula: BMI = weight (kg) / (height (m) * height (m))
	exponent := 2.0
	bmi := weightKg / math.Pow(heightM, exponent)

	// Display the calculated BMI
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are underweight.")
	} else if bmi >= 18.5 && bmi < 25 {
		fmt.Println("You have a normal weight.")
	} else if bmi >= 25 && bmi < 30 {
		fmt.Println("You are overweight.")
	} else {
		fmt.Println("You are obese.")
	}
}
