// https://www.youtube.com/watch?v=sZoRSbokUE8&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=3

package main
import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string
	superman = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	// thorRating := 3.
	// fmt.Printf("%v, %T", thorRating, thorRating)

	var Ironman, CaptAmerica string = "I am Ironman", "I am capt America"
	fmt.Println(Ironman, CaptAmerica)

	var defaultValue1 int
	fmt.Printf("defaultValue d'un int %v\n", defaultValue1)

	var defaultValue2 string
	fmt.Printf("defaultValue2 d'une string %v\n", defaultValue2)

	var(
		spiderman = "I am spiderman"
		age = 18
		powers = "web slings, spidy sense"
		antman = "I am antman"
	)
	fmt.Println(spiderman, age, powers, antman)
}
