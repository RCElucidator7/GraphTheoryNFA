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

	for _, r := range postfix {
		switch r {
			case '.':
				fragment2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fragment1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				fragment1.accept.edge1 = fragment2.initial

				stack = append(stack, &NFA{initial: fragment1.initial, accept: fragment2.accept})

			case '|':
				fragment2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fragment1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				accept := state{}
				initial := state{edge1: fragment1.initial, edge2: fragment2.initial}
				fragment1.accept.edge1 = &accept
				fragment2.accept.edge1 = &accept

				stack = append(stack, &NFA{initial: &initial, accept: &accept})

			case '*':
				fragment := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				accept := state{}
				initial := state{edge1: fragment.initial, edge2: &accept}
				
				//Join accept state of fragment edge1 to the initial state
				fragment.accept.edge1 = fragment.initial
				//Join accept state of fragment edge2 to the accept state
				fragment.accept.edge2 = &accept

				stack = append(stack, &NFA{initial: &initial, accept: &accept})
			default:
				accept := state{}
				initial := state{symbol: r, edge1: &accept}

				stack = append(stack, &NFA{initial: &initial, accept: &accept})
		}
	}
	
	return stack[0]
}

func main() {
	NFA := regex("ab.c*|")
	fmt.Println(NFA)
}