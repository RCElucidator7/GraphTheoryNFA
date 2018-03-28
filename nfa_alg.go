package main

import (
	"fmt"
)

//state structure
type state struct {
	
	symbol rune

	//edges that point at other states
	edge1 *state
	edge2 *state
}

type NFA struct{
	//Value for the inital state
	initial *state
	//Value for the accepted states
	accept *state
}

func regex(postfix string) *NFA {

	//variable for an array of pointers for NFA struct
	stack := []*NFA{}
	
}

func main() {
	NFA := regex("ab.c*|")
	fmt.Println(NFA)
}