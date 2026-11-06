package main

import "fmt"

func main() {

	fmt.Println("===== GO CONDITION PRACTICE =====")

	// Multiple Data Types
	name := "Aftab"
	age := 25
	marks := 85
	cgpa := 3.50
	isStudent := true
	isActive := false

	// 1. IF ELSE CONDITION
	fmt.Println("\n--- IF ELSE ---")

	if marks >= 80 {
		fmt.Println("Grade: A+")
	} else if marks >= 70 {
		fmt.Println("Grade: A")
	} else if marks >= 60 {
		fmt.Println("Grade: A-")
	} else if marks >= 40 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: Fail")
	}

	// 2. NESTED IF ELSE
	fmt.Println("\n--- NESTED IF ---")

	if age >= 18 {
		fmt.Println("Adult")

		if isStudent {
			fmt.Println("You are a Student")
		} else {
			fmt.Println("You are not a Student")
		}

	} else {
		fmt.Println("Minor")
	}

	// 3. LOGICAL OPERATORS
	fmt.Println("\n--- LOGICAL OPERATORS ---")

	if age >= 18 && marks >= 80 {
		fmt.Println("Eligible")
	}

	if age < 18 || marks < 40 {
		fmt.Println("Not Eligible")
	}

	if !isActive {
		fmt.Println("Account is Inactive")
	}

	// 4. SWITCH WITH EXACT VALUE
	fmt.Println("\n--- SWITCH VALUE ---")

	switch age {
	case 18:
		fmt.Println("Age is 18")

	case 25:
		fmt.Println("Age is 25")

	case 30:
		fmt.Println("Age is 30")

	default:
		fmt.Println("Age Not Found")
	}

	// 5. SWITCH WITH MULTIPLE CASE VALUES
	fmt.Println("\n--- MULTIPLE CASE ---")

	day := 5

	switch day {
	case 1, 2, 3:
		fmt.Println("First Group")

	case 4, 5, 6:
		fmt.Println("Second Group")

	case 7:
		fmt.Println("Sunday")

	default:
		fmt.Println("Invalid Day")
	}

	// 6. SWITCH WITHOUT EXPRESSION
	fmt.Println("\n--- SWITCH CONDITION ---")

	switch {
	case age < 18:
		fmt.Println("Minor")

	case age >= 18 && age < 30:
		fmt.Println("Young Adult")

	case age >= 30 && age < 60:
		fmt.Println("Adult")

	default:
		fmt.Println("Senior")
	}

	// 7. FLOAT64 CONDITION
	fmt.Println("\n--- FLOAT CONDITION ---")

	if cgpa >= 3.50 {
		fmt.Println("Excellent CGPA")
	} else if cgpa >= 3.00 {
		fmt.Println("Good CGPA")
	} else {
		fmt.Println("Need Improvement")
	}

	// 8. STRING SWITCH
	fmt.Println("\n--- STRING SWITCH ---")

	role := "developer"

	switch role {
	case "admin":
		fmt.Println("Admin Dashboard")

	case "developer":
		fmt.Println("Developer Dashboard")

	case "user":
		fmt.Println("User Dashboard")

	default:
		fmt.Println("Unknown Role")
	}

	// 9. BOOLEAN CONDITION
	fmt.Println("\n--- BOOLEAN CONDITION ---")

	if isStudent && isActive {
		fmt.Println("Active Student")
	} else if isStudent && !isActive {
		fmt.Println("Inactive Student")
	} else {
		fmt.Println("Not a Student")
	}

	// 10. TYPE SWITCH
	fmt.Println("\n--- TYPE SWITCH ---")

	checkType(name)
	checkType(age)
	checkType(cgpa)
	checkType(isStudent)

	fmt.Println("\n===== END OF PROGRAM =====")


}

// TYPE SWITCH FUNCTION

func checkType(value any) {

	switch v := value.(type) {

	case string:
		fmt.Println(v, "is String")

	case int:
		fmt.Println(v, "is Integer")

	case float64:
		fmt.Println(v, "is Float64")

	case bool:
		fmt.Println(v, "is Boolean")

	default:
		fmt.Println("Unknown Type")
	}
}
