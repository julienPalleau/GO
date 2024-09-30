//https://go.dev/tour/basics/8

/* 
Variables
The var statement declares a list of variables; as in function argument lists, the type is last.
A var statement can be at package or function level. We see both in this example.
uint8 8 bits 0 à 255
uint16 16 bits 0 à 65535
uint32 32 bits 0 à 4294967295
uint64 64 bits 0 à 18446744073709551615
int8 8 bits -128 à 127
int16 16 bits -32768 à 32767
int32 32 bits -2147483648 à 2147483647
int64 64 bits -9223372036854775808 à 9223372036854775807 
float32 32 bits Environ 7 chiffres décimaux
float64 64 bits Environ 16 chiffres décimaux
byte Identique à uint8
rune Identique à int32
32 ou 64 bits non signé
int Même taille que uint mais signé 
booléen true ou false
string 
*/



package main

import "fmt"
var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}