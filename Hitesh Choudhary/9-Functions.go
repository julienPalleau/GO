// https://www.youtube.com/watch?v=feU9DQNoKGE&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=11
package main

import (
	"fmt"
)

func main() {
	superman()
	
	message, result := multiplyme(3, 6)
	fmt.Println(message, result)
	myresult, mylength, myname := adder(3, 6, 4, 2, 4, 12, 23)
	fmt.Println("sum result: ", myresult, mylength, myname)
}

func superman(){
	fmt.Println("I am superman")
}

func multiplyme(v1, v2 int) (string, int){
	return "result: ", (v1 * v2)
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values{
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}