package main

import "fmt"

func Tittle(text string) {
	dashtitle := ""
	for i := 0; i < (78-len(text))/2; i++ {
		dashtitle += "-"
	}
	fmt.Println(dashtitle + text + dashtitle)
}

func main() {

	Tittle("Temperature Converter")

	//Formula as constants
	const celsiusToFahrenheit = 1.8
	const fahrenheitToCelsius = 0.5556

	// Get user input for temperature in Celsius
	fmt.Println("Enter a temperature in Celsius:")
	var celsius float64
	fmt.Scanln(&celsius)

	// Convert Celsius to Fahrenheit
	fahrenheit := celsius*celsiusToFahrenheit + 32

	// Convert Fahrenheit back to Celsius for verification
	celsiusVerify := (fahrenheit - 32) * fahrenheitToCelsius

	// Display results
	fmt.Printf("%.2f°C is equal to %.2f°F\n", celsius, fahrenheit)
	fmt.Printf("Verification: %.2f°F is equal to %.2f°C\n", fahrenheit, celsiusVerify)

}
