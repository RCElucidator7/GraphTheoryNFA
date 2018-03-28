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

	for _, r := postfix {
		switch r {
			case '.':
				fragtwo := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fragone := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				fragone.accept.edge1 = fragtwo.initial

				stack = append(stack, &NFA{initial: fragone.initial, accept: fragtwo.accept})

			case '|':
				fragtwo := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fragone := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				accept := state{}
				initial := state{edge1: fragone.initial, edge2: fragtwo.initial}
				fragone.accept.edge1 = &accept
				fragtwo.accept.edge1 = &accept

				stack = append(stack, &NFA{initial: &initial, accept: &accept})

			case '*':

			default:
		}
	}
	
	return stack[0]
}

func main() {
	NFA := regex("ab.c*|")
	fmt.Println(NFA)
}