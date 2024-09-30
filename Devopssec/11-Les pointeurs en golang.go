// 11-POINTEURS
// https://devopssec.fr/article/pointeurs-golang#begin-article-section

package main

import "fmt"

func main() {
	println()
	var ptr *int

	var a int = 20
	var ap *int
	ap = &a
	fmt.Printf("Adresse de votre variable a : %p\n", &a)
	fmt.Printf("Valeur de votre variable (pointeur) ap: %p\n", ap)
	fmt.Printf("Valeur à l'adresse %p: %d\n", ap, *ap)
	println()

	if ptr == nil { //est ce que le pointeur est vide ?
		fmt.Println("Aucune adresse mémoire")
		fmt.Printf("Adresse mémoire stockée %v\n", ptr)
	} else {
		fmt.Println("Votre adresse mémoire est %p\n", ptr)
	}
	println()

	fmt.Printf("Valeur de la variable a : %d\n", a)
	*ap = 30
	fmt.Printf("Valeur de la variable a : %d \n", a)
	println()

	var b int = 20
	var pb *int
	pb = &b

	if pb == nil {
		fmt.Printf("Aucune adresse mémoire")
	} else {
		fmt.Printf("Votre adresse mémoire est %p\n", pb)
	}
}