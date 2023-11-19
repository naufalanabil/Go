package main

import (
	"fmt"
)

// Student struct to represent student data
type Student struct {
	NPM   string
	Name  string
	Class string
	// Map to store grades for each subject
	SubjectGrades map[string]int
}

// Function to calculate the final grade for a student
func calculateFinalGrade(student Student) float64 {
	// Calculate the total grade for all subjects
	var totalGrade int
	for _, grade := range student.SubjectGrades {
		totalGrade += grade
	}

	// Calculate the average grade
	averageGrade := float64(totalGrade) / float64(len(student.SubjectGrades))

	// Return the average grade
	return averageGrade
}

// Function to convert numeric grade to letter grade (A-E)
func convertToLetterGrade(grade float64) string {
	switch {
	case grade >= 80:
		return "A"
	case grade >= 70:
		return "B"
	case grade >= 60:
		return "C"
	case grade >= 50:
		return "D"
	default:
		return "E"
	}
}

func main() {
	// Example usage
	student := Student{
		NPM:   "51422215",
		Name:  "Naufal Ardra Anabil",
		Class: "2IA14",
		SubjectGrades: map[string]int{
			"Algoritma & Pemrograman 3*":          98,
			"Matematika Lanjut 1":                 81,
			"Matematika Informatika 3":            85,
			"Struktur Data**":                     90,
			"Statistika 1":                        72,
			"Praktikum Algoritma & Pemrograman 3": 98,
			"Organisasi Sistem Komputer*/**":      92,
			"Informatika Kesehatan":               95,
			"Pengantar Sains Data":                95,
			"Bahasa Indonesia":                    90,
			"Komputasi Big Data":                  98,
			"Praktikum Komputasi Big Data":        95,
		},
	}

	// Calculate the final grade
	finalGrade := calculateFinalGrade(student)

	// Convert the final grade to letter grade (A-E)
	letterGrade := convertToLetterGrade(finalGrade)

	// Display the result
	fmt.Printf("Final grade for %s (NPM: %s, Class: %s): %.2f (%s)\n", student.Name, student.NPM, student.Class, finalGrade, letterGrade)
}
