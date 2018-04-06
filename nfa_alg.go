//Ryan Conway
//ID - G00332826
package main

import (
	"fmt"
)

//state structure
type state struct {
	//rune used to detect characters/symbols
	symbol rune

	//edges that point at other states
	edge1 *state
	edge2 *state
}

//NFA struct
type NFA struct {
	//Value for the inital state
	initial *state
	//Value for the accepted states
	accept *state
}

func regex(postfix string) *NFA {

	//Variable for an array of pointers for NFA struct
	stack := []*NFA{}

	for _, r := range postfix {
		switch r {
		case '.':
			//Pop two fragments off the nfa stack and gets rid of last element on the stack
			//Fragment 2 comes off first as it was added last
			fragment2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fragment1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			//Set Fragment1's accept state to the initial state of fragment2
			fragment1.accept.edge1 = fragment2.initial

			//Push pointer to the stack
			stack = append(stack, &NFA{initial: fragment1.initial, accept: fragment2.accept})

		case '|':
			//Pop two fragments off the nfa stack and gets rid of last element on the stack
			//Fragment 2 comes off first as it was added last
			fragment2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fragment1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			accept := state{}
			//Point both edges to the initial state
			initial := state{edge1: fragment1.initial, edge2: fragment2.initial}
			fragment1.accept.edge1 = &accept
			fragment2.accept.edge1 = &accept

			//Push pointer to the stack
			stack = append(stack, &NFA{initial: &initial, accept: &accept})

		case '*':
			//Pop fragment off the stack
			fragment := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			accept := state{}
			//Point the initial at new initial and accept states
			initial := state{edge1: fragment.initial, edge2: &accept}

			//Join accept state of fragment edge1 to the initial state
			fragment.accept.edge1 = fragment.initial
			//Join accept state of fragment edge2 to the accept state
			fragment.accept.edge2 = &accept

			//Push states to the stack
			stack = append(stack, &NFA{initial: &initial, accept: &accept})
		case '+':
			//Pop fragment off the stack
			fragment := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			accept := state{}
			//initial := state{edge1: fragment.initial}
			middle := state{edge1: fragment.initial, edge2: &accept}

			fragment.accept.edge1 = &middle

			//Push states to the stack
			stack = append(stack, &NFA{initial: fragment.initial, accept: &accept})
		case '?':
			//Pop fragment off the stack
			fragment := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			accept := state{}
			initial := state{edge1: fragment.initial, edge2: &accept}

			fragment.accept.edge1 = &accept

			//Push states to the stack
			stack = append(stack, &NFA{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			//Push states to the stack
			stack = append(stack, &NFA{initial: &initial, accept: &accept})
		}
	}

	//Error handler was throwing errors
	/*//Error handling
	if len(stack) != 1 {
		fmt.Println("Error: ", len(stack), stack)
	}*/

	return stack[0]
}

//Function used to get the current state, checking if the state has an edge and follows the edge to get the next state
func addState(list []*state, s *state, a *state) []*state {
	list = append(list, s)

	//Checks if s.symbol has a 0 value rune, meaning it has an edge
	if s != a && s.symbol == 0 {
		list = addState(list, s.edge1, a)

		//Checks if it has another edge, and then adds it
		if s.edge2 != nil {
			list = addState(list, s.edge2, a)
		}
	}
	return list
}

func postmatch(post string, s string) bool {
	match := false
	nfa := regex(post)

	//Linked list of states
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

//Adding the infix to postfix function from the shunt_alg file to demonstrate how both work together
//Function that translates from infix to postfix notation
func infixToPostfix(infix string) string {
	//Assign special characters int values
	special := map[rune]int{'*': 10, '.': 9, '|': 8, '+': 7, '?': 6}

	//Empty array of runes
	pFix := []rune{}

	stack := []rune{}

	//Loop over string and convert to postfix
	for _, r := range infix {
		//range is used to convert type string into type rune
		switch {
		case r == '(':
			stack = append(stack, r)
		case r == ')':
			for stack[len(stack)-1] != '(' {
				pFix = append(pFix, stack[len(stack)-1])

				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1]
		case special[r] > 0:
			for len(stack) > 0 && special[r] <= special[stack[len(stack)-1]] {
				pFix = append(pFix, stack[len(stack)-1])

				stack = stack[:len(stack)-1]
			}
			stack = append(stack, r)
		default:
			pFix = append(pFix, r)
		}
	}
	//Takes top element of stack
	for len(stack) > 0 {
		pFix = append(pFix, stack[len(stack)-1])

		stack = stack[:len(stack)-1]
	}

	return string(pFix)
}

func main() {
	//Regular expression example
	NFA := regex("ab.c*|")
	fmt.Println(NFA)

	//Examples
	fmt.Println("Regular Expression match function:")
	fmt.Println("Input : ab.c*|")
	//Should Return True
	fmt.Println("Output:", postmatch("ab.c*|", "cccc"))
	fmt.Println()

	fmt.Println("Input : ab.c*|")
	//Should Return Flase
	fmt.Println("Output:", postmatch("ab.c*|", "abc"))
	fmt.Println()

	fmt.Println("Input : abd|.c*")
	//Should Return true
	fmt.Println("Output:", postmatch("abd|.c*", "ab"))
	fmt.Println()

	fmt.Println("Input : abd|.c*")
	//Should Return false
	fmt.Println("Output:", postmatch("abd|.c*", "abd"))
	fmt.Println()

	fmt.Println("Converting (a.b.c)* to a postfix notation")
	value := infixToPostfix("(a.b.c)*")
	fmt.Println("Postfix value: ", value)

	fmt.Println("Checking if the postfix value matches the string 'abc'")
	matchValue := postmatch(value, "abc")
	//Should return true
	fmt.Println("Does the string match? : ", matchValue)
}
