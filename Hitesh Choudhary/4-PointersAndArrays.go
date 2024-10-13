// https://www.youtube.com/watch?v=xbtwQnookZ4&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=6

package main

import (
	"fmt"
)

// func main(){
// 	var p *int

// 	if p != nil {
// 		fmt.Println("P is having a value: ", *p)
// 	} else {
// 		fmt.Println("Watchout for nil values")
// 	}
// }

// func main(){
// 	var lifebooster float64 = 99.2
// 	lifeboosterRef := &lifebooster

// 	lifebooster = lifebooster * 2.2

// 	fmt.Println(lifebooster)
// 	fmt.Println(lifeboosterRef)
// 	fmt.Println(*lifeboosterRef)
// }

// myArray
		func main(){
			var numbers [3]string
			numbers[0] = "uno"
			numbers[1] = "dos"
			numbers[2] = "tres"

			fmt.Println(numbers)

			var colors = [4]string{"rojo", "gris", "azul", "verde"}
			fmt.Println(colors)
			fmt.Println(colors)
			fmt.Println(len(colors))

		}