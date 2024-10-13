// https://www.youtube.com/watch?v=xWaDl_ycxXE&list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N&index=7

package main
import "fmt"

main() {
	var s []byte
	s = make([]byte, 5)
	fmt.Println("taille de s",len(s))
	fmt.Println("Capacite de s",cap(s))
}