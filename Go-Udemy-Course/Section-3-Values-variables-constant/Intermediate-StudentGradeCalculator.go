package main

import (
	"fmt"
)

func Tittle(text string) {
	dashtitle := ""
	for i := 0; i < (78-len(text))/2; i++ {
		dashtitle += "-"
	}
	fmt.Println(dashtitle + text + dashtitle)
}

func main() {

	Tittle("Student Grade Calculator")

	// Constants for grade thresholds
	const (
		APlus = 97
		A     = 93
		BPlus = 90
		B     = 87
		CPlus = 83
		C     = 80
		DPlus = 77
		D     = 73
		F     = 0
	)

	// Constants for weight of each component
	const quizWeight = 0.2
	const midtermWeight = 0.3
	const finalExamWeight = 0.5

	// Get user input for student's score
	fmt.Println("Enter student's score in quiz (0-100):")
	var quizScore float64
	fmt.Scanln(&quizScore)

	fmt.Println("Enter student's score in midterm (0-100):")
	var midtermScore float64
	fmt.Scanln(&midtermScore)

	fmt.Println("Enter student's score in final exam (0-100):")
	var finalExamScore float64
	fmt.Scanln(&finalExamScore)

	// Calculate the final grade
	finalGrade := quizScore*quizWeight + midtermScore*midtermWeight + finalExamScore*finalExamWeight

	// Determine letter grade based on final grade
	var letterGrade string
	switch {
	case finalGrade >= APlus:
		letterGrade = "A+"
	case finalGrade >= A:
		letterGrade = "A"
	case finalGrade >= BPlus:
		letterGrade = "B+"
	case finalGrade >= B:
		letterGrade = "B"
	case finalGrade >= CPlus:
		letterGrade = "C+"
	case finalGrade >= C:
		letterGrade = "C"
	case finalGrade >= DPlus:
		letterGrade = "D+"
	case finalGrade >= D:
		letterGrade = "D"
	default:
		letterGrade = "F"
	}

	// Display results
	fmt.Printf("Final Grade: %.2f\n", finalGrade)
	fmt.Printf("Letter Grade: %s\n", letterGrade)
}
