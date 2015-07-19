package main

import (
	"fmt"
	"github.com/soniah/evaler"
	"os"
	"strconv"
)

////
//Array functions
////

//works on copy
func remove(pool []int, n int) []int {
	//TODO: Better way to do this
	c := append([]int{}, pool[0:n]...)
	return append(c, pool[n+1:]...)
}

func permutate(arr []int) [][]int {
	if len(arr) == 1 {
		return [][]int{arr}
	}

	//recursively permutate:
	//take each element and permutate the remaining list
	//combine the just took element with each of the permutated lists
	//also add the smaller groups
	permutations := [][]int{}
	for i, num := range arr {
		for _, sub_perm := range permutate(remove(arr, i)) {
			permutations = append(permutations, append(sub_perm, num))

			//for this use case also add the smaller groups
			//this way we automatically get EVERY possible combination
			//permutations = append(permutations, sub_perm)
		}
	}
	return permutations
}

func repetative_choose(num_elements int, from []int) [][]int {
	r := [][]int{}
	if num_elements == 1 {
		for _, poss := range from {
			r = append(r, []int{poss})
		}
		return r
	}

	for _, poss := range from {
		subs := repetative_choose(num_elements-1, from)
		for _, sub := range subs {
			r = append(r, append(sub, poss))
		}
	}

	return r
}

//Too lazy for trees :)
//digits are the digits used
//operations is a possible permutation of operations
//will generate all possible orders of operations
//and return a list of all expressions found this way
func generate_possible_expressions(digits []int, operations []string) []string {
	_operations_indices := make([]int, len(operations))
	for i := 0; i < len(operations); i++ {
		_operations_indices[i] = i
	}

	bracket_permutations := permutate(_operations_indices)

	r := []string{}

	for _, brackets := range bracket_permutations {
		//array of [operator,left,right, nesting level]
		//nesting level is more of a "have been used" flag
		d := make([][]string, len(operations))
		for i, op := range operations {
			d[i] = []string{op, strconv.Itoa(digits[i]), strconv.Itoa(digits[i+1]), "0"}
		}

		expression := ""
		for i, op_num := range brackets {
			sub_expression := "(" + d[op_num][1] + d[op_num][0] + d[op_num][2] + ")"

			//notify operators to the left and to the right of the new product
			//go multiple levels
			//TODO: These two loops can be merged. leave them duplicated for debugging purposes
			for i := 1; op_num-i >= 0; i++ {
				if d[op_num-i][3] == "0" {
					d[op_num-i][2] = sub_expression
					break
				}
			}
			for i := 1; op_num+i < len(brackets); i++ {
				if d[op_num+i][3] == "0" {
					d[op_num+i][1] = sub_expression
					break
				}
			}

			d[op_num][3] = "1"

			if i == len(brackets)-1 {
				expression = sub_expression
			}
		}

		r = append(r, expression)
	}

	return r
}

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Error: Exactly four numerical arguments are expected.")
		return
	}

	input_numbers := []int{}
	operations := []string{"+", "-", "*", "/"}
	for _, num := range os.Args[1:] {
		if num, err := strconv.Atoi(num); err != nil {
			fmt.Println("Error: Exactly four numerical arguments are expected")
			return
		} else {
			input_numbers = append(input_numbers, num)
		}
	}

	digit_permutations := permutate(input_numbers)
	operation_permutations := repetative_choose(3, []int{0, 1, 2, 3})

	expressions := []string{}

	//combine every digit permutation with every operation permutation and with every bracket permutation
	for _, digits := range digit_permutations {
		for _, operation_indices := range operation_permutations {
			//operation index -> operation string
			//cause no generic permutation :(
			ops := make([]string, 3)
			for i, operation_index := range operation_indices {
				ops[i] = operations[operation_index]
			}

			expressions = append(expressions, generate_possible_expressions(digits, ops)...)
		}
	}

	fmt.Println("Got", len(expressions), "expressions. Now calculating")

	c := 0
	duplicates := make(map[string]bool)
	for _, expr := range expressions {
		if duplicates[expr] {
			continue
		}

		if res, err := evaler.Eval(expr); err == nil {
			if res2 := evaler.BigratToFloat(res); res2 == 24.0 {
				fmt.Println(expr)
				duplicates[expr] = true
				c++
			}
		}
	}
	fmt.Println("Got", c, "solutions")
}
