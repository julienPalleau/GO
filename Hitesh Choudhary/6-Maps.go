// https://www.youtube.com/watch?v=p4LS3UdgJA4&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=8

package main

import (
	"fmt"
)

func main() {
	// make new
	// new - only allocates - no init of memory
	// make - allocate and init - non zero storage

	// when running the 3 lines of codes below we got "panic: assignment to entry in nil map"
	// this is because memory was allocate but not initialized 
	// var score map[string]int // this is a new
	// score["hitesh"] = 89
	// fmt.Println(score)

	// So let's modify the 3 lines of codes above as follow:
	fmt.Println("Let's create our map with make(map[string]int)")
	score := make(map[string]int)
	score["hitesh"] = 89
	fmt.Println(score)

	// Now we have been able to create our score we can add it more values
	fmt.Println("\nNow we have been able to create our score we can add it more values")
	fmt.Println("score[\"josh\"] = 34\nscore[\"ron\"] = 23\nscore[\"sam\"] = 56\nscore[\"vicky\"] = 78\n")
	score["josh"] = 34
	score["ron"] = 23
	score["sam"] = 56
	score["vicky"] = 78
	fmt.Println(score)

	// Get a value
	fmt.Println("\nGet a value:")
	getRonScore := score["ron"]
	fmt.Println("getRonScore := score[\"ron\"], getRonScore:",getRonScore)

	// Delete a value
	fmt.Println("\nDelete a value")
	fmt.Println("before we delete vicky, score:",score)
	delete(score, "vicky")
	fmt.Println("delete(score, \"vicky\"), score:",score)

	// How to iterate through a map?
	fmt.Println("\nHow to iterate through a map?")
	fmt.Println("for k, v := range score {\n	fmt.Printf(\"Score of %v is %v\", k, v)")
	for k, v := range score {
		fmt.Printf("Score of %v is %v\n", k, v)
	}
}