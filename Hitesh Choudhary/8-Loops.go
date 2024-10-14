// https://www.youtube.com/watch?v=0A5fReZUdRk&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=10
package main

import(
	"fmt"
)

func main() {
	start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(things)

	fmt.Println("First for loop: ")
	fmt.Println("for i:=0;  i < 10; i++{\n fmt.Println(i + start)\n}")
	for i:=0; i < 10; i++{
		fmt.Println(i + start)
	}

	fmt.Println("\nSecond for loop: ")
	fmt.Println("for i:=0; i<len(things); i++{\n fmt.Println(things[i])")
	for i:=0; i<len(things); i++{
		fmt.Println(things[i])
	}

	fmt.Println("\nThird for loop: ")
	fmt.Println("for i:=range things {\n fmt.Println(things[i])")
	for i:=range things {
		fmt.Println(things[i])
	}

	fmt.Println("\nFourth for loop & break:")
	fmt.Println("for start < 100 {\nstart += start\nif start == 32 {\nbreak\n}\nfmt.Println(\"START is now: \", start)")
	start = 1
	for start < 100 {
		start += start
		if start == 32 {
			break
		}
		fmt.Println("START is now: ", start)
	}

	fmt.Println("\nFourth for loop & continue:")
	fmt.Println("for start < 100 {\nstart += start\nif start == 32 {\ncontinue\n}\nfmt.Println(\"START is now: \", start)")
	start = 1
	for start < 100 {
		start += start
		if start == 32 {
			continue
		}
		fmt.Println("START is now: ", start)
	}

	fmt.Println("\nFith for loop & lable:")
	fmt.Println("for start < 100 {\nstart += start\nif start == 32 {\ngoto lcolabel\n}\nfmt.Println(\"START is now: \", start)\nlcolabel: fmt.Println(\"LearnCodeOnLine.in\")")
	start = 1
	for start < 100 {
		start += start
		if start == 32 {
			goto lcolabel
		}
		fmt.Println("START is now: ", start)
	}
	lcolabel: fmt.Println("LearnCodeOnLine.in")
}