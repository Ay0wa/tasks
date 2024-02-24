package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var list []string
	var oper string
	for {
		fmt.Println("Введите ЧислоОперацияЧисло - exit - выход")
		st := bufio.NewScanner(os.Stdin)
		st.Scan()
		if st.Text() == "exit" {
			break
		}
		if strings.Contains(st.Text(), "+") {
			oper = "+"
		} else if strings.Contains(st.Text(), "-") {
			oper = "-"
		} else if strings.Contains(st.Text(), "/") {
			oper = "/"
		} else if strings.Contains(st.Text(), "*") {
			oper = "*"
		} else {
			fmt.Println("Введите корректное значение!")
			break
		}
		list = strings.Split(st.Text(), oper)
		a, error1 := strconv.ParseInt(list[0], 10, 64)
		b, error2 := strconv.ParseInt(list[1], 10, 64)
		if error1 != nil || error2 != nil {
			fmt.Println("Введите корректное значение!")
		} else {
			fmt.Println(a, oper, b, "=", result(a, b, oper))
		}

	}

}

func result(a int64, b int64, oper string) int64 {
	if oper == "+" {
		var res int64 = a + b
		return res
	} else if oper == "-" {
		var res int64 = a - b
		return res
	} else if oper == "*" {
		var res int64 = a * b
		return res
	} else if oper == "/" {
		var res int64 = a / b
		return res
	}
	return 0
}
