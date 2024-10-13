// https://www.youtube.com/watch?v=YMfWwKSG3Vw&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=4
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var myString string
	// fmt.Sscanln(&myString)
	// fmt.Println(myString)

	// How can we discard error code returned with _
	// var name string = "hitesh"
	// var a, _ = fmt.Println(name)
	// fmt.Println(a)

	// How we can take input that has got some spaces in that
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your full name: ")
	// myName, _ := reader.ReadString('\n')
	// fmt.Println(myName)

	// How we can take input and convert it in floating value and remove spaces
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your rating: ")
	myrating, _ := reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumRating + 2)
}