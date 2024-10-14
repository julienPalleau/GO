// https://www.youtube.com/watch?v=xWaDl_ycxXE&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=7


/*
First let's have a look to slice introduction: https://go.dev/blog/slices-intro
*/
// package main
// import "fmt"

// func main() {
// 	var s []byte
// 	s = make([]byte, 5)
// 	fmt.Println("taille de s",len(s))
// 	fmt.Println("Capacite de s",cap(s))

// 	// Shrink to a length shorter than its capacity
// 	fmt.Println("\nLet's shrink s to a length shorter than its capacity with s=s[2:4]")
// 	s = s[2:4]
// 	fmt.Println("taille de s",len(s))
// 	fmt.Println("Capacite de s",cap(s))
	
// 	// Grow s to its capacity by slicing it again
// 	fmt.Println("\nNow we can grow s to its capacity by slicing it again")
// 	s = s[:cap(s)]
// 	fmt.Println("taille de s",len(s))
// 	fmt.Println("Capacite de s",cap(s))

// 	fmt.Println("\nA slice cannot be grown beyond its capacity ! Slice cannot be re-sliced below zero !")

// 	// Growing slices (the copy and append functions)
// 	fmt.Println("\nGrowing slices (the copy and append functions)")
// 	fmt.Println("Lets copy in t the slice s:",s)
// 	fmt.Println("t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0\n for i := range s {\n t[i] = s[i]\n}\ns=t")
// 	s[0]=0
// 	s[1]=1
// 	s[2]=2
// 	t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
// 	for i := range s {
// 		t[i] = s[i]
// 	}
// 	s=t
// 	fmt.Println(s)
// 	fmt.Println(t)

// 	// Using copy, is much simpler !
// 	fmt.Println("\nUsing copy, we can simplify the code snippet above:")
// 	fmt.Println("u := make([]byte, len(s), (cap(s)+1)*2)\n copy(u, s)\n s = u")
// 	s[0]=10
// 	s[1]=11
// 	s[2]=12
// 	u := make([]byte, len(t), (cap(t)+1)*2) // +1 in case cap(s) == 0
// 	fmt.Println("u: ", u)
// 	fmt.Println("s: ", s)
// 	copy(u, s)
// 	u = s
// 	fmt.Printf("we do: copy(u, s) we got: u: %v, s: %v",u, s)

// 	// Append
// 	fmt.Println("\n\nAppend")
// 	a := make([]int, 1)
// 	fmt.Println("We have a:", a)
// 	a = append(a, 1, 2, 3)
// 	fmt.Println("We appended 1, 2, 3 with append(a, 1, 2, 3) a: ", a)

// 	// To append one slice to another, use ... to expand the second argument to a list of arguments.
// 	fmt.Println("\n\nTo append one slice to another, use ... to expand the second argument to a list of arguments.")
// 	c := []string{"John", "Paul"}
// 	d := []string{"George", "Ringo", "Pete"}
// 	c = append(c, d...) // equivalent to "append(a, b[0], b[1], b[2])"
// 	fmt.Println("We append John, Paul to George, Ringo, Pete with append()",c)
// }

/*
Let's go back to youtube lesson on slices
*/
package main

import(
	"fmt"
	"sort"
)

func main() {
	var things = []string{"maleta", "ropa", "reloj"}
	fmt.Println(things)

	things = append(things, "bolso")
	fmt.Println(things)

	things = append(things[1:])
	fmt.Println(things)

	things = append(things[1: len(things)-1])
	fmt.Println(things)

	heros := make([]string, 3, 3)
	heros[0] = "thor"
	heros[1] = "ironman"
	heros[2] = "spiderman"
	fmt.Println(heros)
	fmt.Println(cap(heros))

	heros = append(heros, "deadpool") //no error even if we declare our slice for 3 elments. Go reserve more memory on the fly.
	fmt.Println(heros)
	fmt.Println(cap(heros))

	// Sort in golang: https://yourbasic.org/golang/how-to-sort-in-go/

	myints := []int{4, 7, 3, 2, 8}

	isSorted := sort.IntsAreSorted(myints)
	fmt.Println("Are ints sorted: ", isSorted)

	sort.Ints(myints)
	fmt.Println("Sorted int numbers", myints)
}