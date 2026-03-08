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
	Tittle("Currency Converter")

	// Formula as constants
	const usdToEur = 0.86
	const usdToMxn = 17.80
	const eurToUsd = 1.1627
	const eurToMxn = 20.78
	const mxnToEur = 0.04812
	const mxnToUsd = 0.05617

	// Select conversion type
	fmt.Println("Select conversion type:")
	fmt.Println("1: USD to EUR")
	fmt.Println("2: USD to MXN")
	fmt.Println("3: EUR to USD")
	fmt.Println("4: EUR to MXN")
	fmt.Println("5: MXN to USD")
	fmt.Println("6: MXN to EUR")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		// Get user input for amount in USD
		fmt.Println("Enter an amount in USD:")
		var usd float64
		fmt.Scanln(&usd)

		// Convert USD to EUR
		eur := usd * usdToEur

		// Display results
		fmt.Printf("%.2f USD is equal to %.2f EUR\n", usd, eur)

	case 2:
		// Get user input for amount in USD
		fmt.Println("Enter an amount in USD:")
		var usd float64
		fmt.Scanln(&usd)

		// Convert USD to MXN
		mxn := usd * usdToMxn

		// Display results
		fmt.Printf("%.2f USD is equal to %.2f MXN\n", usd, mxn)

	case 3:
		// Get user input for amount in EUR
		fmt.Println("Enter an amount in EUR:")
		var eur float64
		fmt.Scanln(&eur)

		// Convert EUR to USD
		usd := eur * eurToUsd

		// Display results
		fmt.Printf("%.2f EUR is equal to %.2f USD\n", eur, usd)

	case 4:
		// Get user input for amount in EUR
		fmt.Println("Enter an amount in EUR:")
		var eur float64
		fmt.Scanln(&eur)

		// Convert EUR to MXN
		mxn := eur * eurToMxn

		// Display results
		fmt.Printf("%.2f EUR is equal to %.2f MXN\n", eur, mxn)

	case 5:
		// Get user input for amount in MXN
		fmt.Println("Enter an amount in MXN:")
		var mxn float64
		fmt.Scanln(&mxn)

		// Convert MXN to USD
		usd := mxn * mxnToUsd

		// Display results
		fmt.Printf("%.2f MXN is equal to %.2f USD\n", mxn, usd)

	case 6:
		// Get user input for amount in MXN
		fmt.Println("Enter an amount in MXN:")
		var mxn float64
		fmt.Scanln(&mxn)

		// Convert MXN to EUR
		eur := mxn * mxnToEur

		// Display results
		fmt.Printf("%.2f MXN is equal to %.2f EUR\n", mxn, eur)
	}
}
