// https://www.youtube.com/watch?v=XOTDzICtIfw&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=5

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userRating string

	//Front end
	//take name as input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name:")
	name, _ = reader.ReadString('\n')

	//take rating from user and convert it to float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Dosa center between 1 and 5: ")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	//Back End
	fmt.Printf("\n\nHello %s, \n Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v\n\n", 
	name, mynumRating, time.Now().Format(time.Stamp))

	if mynumRating == 5 {
		fmt.Println("Bonus for team for 5 stars service")
	} else if mynumRating == 4 || mynumRating == 3 {
		fmt.Println("We are always improving")
	} else if mynumRating < 3 {
		fmt.Println("Need Serious work on our side")
	}
}