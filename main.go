package main

import (
	"fmt"
	"strconv"
	"strings"
)

var err error

/*
isCorrect() --- check is correct input data from user or not
Correct means data in format "number1symbolnumber2"

input:
	str --- input data from
return:
	true if it correct and false if not
*/



func IsCorrect(str string) bool {
	var length int = len(str)
	var symbols_count int = 0

	for i := 0; i < length; i++ {
		if string([]rune(str)[i]) == "+" || string([]rune(str)[i]) == "-" {
			symbols_count++
			if symbols_count > 1 || i == 0 || i == length {
				return false
			}
		} else {
			_, err = strconv.Atoi(string([]rune(str)[i]))

			if err != nil {
				return false
			}
		}
	}
	if symbols_count == 0 {
		return false
	}
	return true
}

/*
getNumbers() --- split data from user into 3 strintgs and make number one length
				by adding "0"
input:
	str --- input data from
return:
	number1, number2, + or -
*/

func GetNumbers(str string) (string, string, string) {
	var num []string
	var symbol string

	if strings.Contains(str, "+") {
		symbol = "+"
	} else {
		symbol = "-"
	}

	num = strings.Split(str, symbol)

	for len(num[0]) > len(num[1]) {
		num[1] = "0" + num[1]
	}

	for len(num[0]) < len(num[1]) {
		num[0] = "0" + num[0]
	}

	return num[0], num[1], "-"
}

/*
converToArray() --- convert string to array of int
input:
	str --- string to convert
return:
	arr --- array of int
*/

func ConverToArray(str string) []int {
	var arr []int
	var element int

	for i := 0; i < len(str); i++ {
		element, err = strconv.Atoi(string([]rune(str)[i]))
		arr = append(arr, element)
	}

	return arr
}

/*
converToString() --- convert array of int to string
input:
	arr --- array of int to convert
return:
	str --- string
*/

func ConverToString(arr []int) string {
	var str string
	var i int = 0

	for arr[i] == 0 && i < len(arr)-1 {
		i++
	}
	for i < len(arr) {
		str += strconv.Itoa(arr[i])
		i++
	}

	return str
}

/*
add() --- add numbers
input:
	numb1,numb2 -- numbers to add
return:
	numb1 -- sum
*/
func Add(numb1, numb2 []int) string {

	var length int = len(numb1)

	for i := length - 1; i > -1; i-- {
		numb1[i] += numb2[i]

		if numb1[i] > 9 && i != 0 {
			numb1[i] -= 10
			numb1[i-1]++
		}
	}

	return ConverToString(numb1)
}

/*
isNegative() --- check can be result of subtraction negative
input:
	numb1,numb2 --- numbers to subtraction
return:
	true if result will be negative and false if not
*/

func IsNegative(numb1, numb2 []int) bool {

	var i = 0

		for numb1[i] == numb2[i] && i != 0 {
			i++
		}
		if i != 0 {
			i++
		}

		if numb1[i] >= numb2[i] {
			return false
		} else {
			return true
		}
}

/*
sub() --- substract numbers
input:
	numb1,numb2 -- numbers to add
return:
	numb1 -- result of subtraction
*/
func Sub(numb1, numb2 []int) string {

	var length int = len(numb1)
	var prefix string = ""

	if IsNegative(numb1, numb2) {
		tmp := numb1
		numb1 = numb2
		numb2 = tmp
		prefix = "-"
	}

	for i := length - 1; i > -1; i-- {
		numb1[i] -= numb2[i]

		if numb1[i] < 0 && i != 0 {
			numb1[i] += 10
			numb1[i-1]--
		}
	}

	return prefix + ConverToString(numb1)
}

func main() {
	var input, num1, num2, symbol, result string

	fmt.Println("Enter expression: ")
	fmt.Scanln(&input)

	if IsCorrect(input) {
		num1, num2, symbol = GetNumbers(input)
		number1 := ConverToArray(num1)
		number2 := ConverToArray(num2)
		if symbol == "+" {
			result = Add(number1, number2)
		} else {
			result = Sub(number1, number2)
		}
		fmt.Println("Result\n" + result)
	} else {
		fmt.Println("Expression is not correct.")
	}
}
