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

//NFA struct
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
	
	//Error handling
	if len(stack) != 1 {
		fmt.Println("Error: ", len(stack), stack)
	}

	return stack[0]
}

func addState(list []*state, s *state, a *state) []*state {
	list = append(list, s)

	if s != a && s.symbol == 0 {
		list = addState(list, s.edge1, a)
		if s.edge2 != nil {
			list = addState(list, s.edge2, a)
		}
	}
	return list
}

func postmatch(post string, s string) bool {
	match := false
	nfa := regex(post)

	//Licked list of states
	current := []*state{nfa.initial}
	next := []*state{}

	current = addState(current[:], nfa.initial, nfa.accept)

	for _, r := range s {
		for _, curs := range current {

			if curs.symbol == r {
				//If the current character from the input is equal to the symbol of the current state
				if curs.symbol == r {
					next = addState(next[:], curs.edge1, nfa.accept)
				}
			}
		}
		//Set current states to next states
		current = next
		//Clear next states
		next = []*state{}
	}

	for _, curs := range current {
		if curs == nfa.accept {
			match = true
			break
		}
	}

	return match
}

func main() {
	NFA := regex("ab.c*|")
	fmt.Println(NFA)

	//Should Return True
	fmt.Println(postmatch("ab.c*|", "cccc"))

	//Should Return Flase
	fmt.Println(postmatch("ab.c*|", "abc"))

}