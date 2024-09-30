// https://www.youtube.com/watch?v=LHhsNa_Kgns
package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	// // Defining Variables
	// var name string = "Tim"
	// var foot_size int = 38
	// var age uint = 29 // uint means unsigned integer
	// var taille float64 = 3.14
	// var sex bool = true // true or false
	// fmt.Println(name, age, taille, foot_size, sex)

	// // Implicit variable definition
	// nickname := "Joe"
	// fmt.Println(nickname, age, taille, foot_size, sex)

	// // Look at print more in detail
	// fmt.Printf("Hello %v, you are %v years old!", nickname, age) // %v takes the value of variable name like the old python way to print

	// // Ask our user what is his name
	// var name_asked string
	// var age_asked int

	// fmt.Scan(&name_asked)
	// fmt.Scan(&age_asked)

	// fmt.Println(name_asked)
	// fmt.Println(age_asked)

	var name string

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)
	fmt.Println(age >= 10)


	// CONDITIONALS

	if age >=10 {
		fmt.Println("Yay you can play!")
	} else {
		fmt.Println("You can not play!")
		return
	}

	score := 0
	num_questions := 2

	// User input with Scan()

	fmt.Printf("What is better, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2) // The scan cut the answer on white space, so for now we need to get the full answer ex: RTX 3090
	fmt.Println(answer, answer2)

	if answer + " " + answer2 == "RTX 3090" || answer + " " + answer2 == "rtx 3090" || answer + " " + answer2 == "rtX 3090" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have?")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct !")
		score++
	} else {
		fmt.Println("Incorrect !")
	}

	fmt.Printf("You scored %v out of %v.", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100 // if you divide an int by an int you get an int
	fmt.Printf("You scored: %v%%.", percent)
}