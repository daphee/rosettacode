package main

import (
	"fmt"
	"github.com/soniah/evaler"
	"time"
	"math/rand"
	"strings"
	"os"
	"bufio"
	"strconv"
	"regexp"
)

const (
	ERROR_MSG string = "Boooo! That wasn't a valid expression. Please use every number exactly once. Allowed operators are + - / * ( )"
	REGEXP string = "[^0-9]*"
)

var (
	reg *regexp.Regexp = regexp.MustCompile(REGEXP)
)

func generate_numbers () [4]string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numbers := [4]string{}
	for i:=0; i < 4; i++ {
		numbers[i] = strconv.Itoa((r.Intn(10) + 1))
	}	

	return numbers
}

//just checks if every number is used once
//evaler does the rest
func valid_expression(line string, numbers [4]string) bool {
	numbers_string := strings.Trim(reg.ReplaceAllString(line, " "), " ")	

	for _, num := range numbers {
		new_string := strings.Replace(numbers_string, num, "", 1)
		if new_string == numbers_string {
			return false
		} 
		numbers_string = new_string
	}

	return true
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for true {
		numbers := generate_numbers()
		fmt.Println("Your numbers: ", strings.Join(numbers[:], " "))

		line,_,_:= stdin.ReadLine()
		if valid_expression(string(line), numbers) {
			result, err := evaler.Eval(string(line))
			if err != nil {
				fmt.Println(ERROR_MSG)
			} else if r, _ := evaler.BigratToInt(result); r == 24 {
				fmt.Println("You got it!")
			} else {
				fmt.Println("That wasn't correct :(")
			}

		} else {
			fmt.Println(ERROR_MSG)
		}
	}			
}