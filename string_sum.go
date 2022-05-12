package string_sum

import (
	"errors"
	"strconv"
	"strings"
	"fmt"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	in := strings.ReplaceAll(input," ","")
	if in == "" {return "", errorEmptyInput}
	in = strings.ReplaceAll(in,"+"," +")
	in = strings.ReplaceAll(in,"-"," -")
	for _, v := range in {
		if v != '+' && v != '-' && !unicode.IsNumber(v) {
			return "", fmt.Errorf("not a valid input")
		}	
	}
	oper := strings.Fields(in)
	if len(oper) != 2 {return "", errorNotTwoOperands}
	s1, er1 :=strconv.Atoi(oper[0])
	if er1 != nil {return "", er1}
	s2, er2 :=strconv.Atoi(oper[1])
	if er2 != nil {return "", er2}
	out := s1+s2
	return strconv.Itoa(out), nil
}
