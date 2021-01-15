package main

import (
	"fmt"
	"strconv"
)

// five
var m int = 60

// six
/*
var actorName string = "Elisabeth Sladen"
var companion string = "Sarah Jane Smith"
var doctorNumber int = 3
var season int = 11
*/
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

// seven
var n int = 34 // lowercase - Package access
var O int = 34 // uppercase - Public access

func main() {
	// one
	var i int
	i = 42
	// two
	var j int = 27
	// three
	k := 13

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	// four
	var l float32 = 27

	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", m, m)

	// eight
	var p int = 77
	fmt.Printf("%v, %T\n", p, p)
	var q float32
	q = float32(p)
	fmt.Printf("%v, %T\n", q, q)
	var r string
	r = string(p)
	fmt.Printf("%v, %T\n", r, r)
	r = strconv.Itoa(p)
	fmt.Printf("%v, %T\n", r, r)
}
