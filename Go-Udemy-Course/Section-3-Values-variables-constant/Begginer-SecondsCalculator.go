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
	Tittle("Seconds Calculator")
	// Constants for time conversion
	const secondsInMinute = 60
	const secondsInHour = 3600
	const secondsInDay = 86400

	// Get user input for time in minutes
	fmt.Println("Enter time in Days, Hours, Minutes and Seconds:")
	var days, hours, minutes, seconds int
	fmt.Println("Days")
	fmt.Scanln(&days)
	fmt.Println("Hours")
	fmt.Scanln(&hours)
	fmt.Println("Minutes")
	fmt.Scanln(&minutes)
	fmt.Println("Seconds")
	fmt.Scanln(&seconds)

	//Calculate total secondsº
	totalSeconds := days*secondsInDay + hours*secondsInHour + minutes*secondsInMinute + seconds

	// Display the calculated total seconds
	fmt.Printf("Total time in seconds: %d\n", totalSeconds)
}
