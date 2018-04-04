package main

import (
	"fmt"
)

//Function that translates from infix to postfix notation
func infixToPostfix(infix string) string {
	//Assign special characters int values
	special := map[rune]int{'*': 10 , '.': 9 , '|': 8, '+': 7, '?': 6}


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
				for stack[len(stack)-1] != '('{
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
	//Input a.b.c*
	fmt.Println("Infix: a.b.c*")
	fmt.Println("Postfis: ", infixToPostfix("(a.b.c*)"))
	//Should return ab.c*
}