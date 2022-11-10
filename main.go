package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Введите число больше нуля")
	var (
		numStr string
		num    int
		err    error
	)

	for {
		fmt.Scan(&numStr)
		num, err = strconv.Atoi(numStr)
		if err != nil && len(numStr) > 0 {
			fmt.Println("Не похоже на число", err)
			continue
		}
		if num < 1 {
			fmt.Println("Число должно быть больше 0. Повторите ввод")
			continue
		}
		break
	}

	var arr []string
	for i := 1; i < num; i++ {
		if getSimple(i) {
			arr = append(arr, strconv.Itoa(i))
		}
	}

	file, err := os.OpenFile("Output.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		file, err = os.Create("Output.txt")
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(time.Now().String() + "\n"))
	_, err = file.WriteString("чисел в последовательности: " + strconv.Itoa(len(arr)) + "\n")
	_, err = file.WriteString(strings.Join(arr, ", "))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат записан в файл Output.txt")
	}
}

func getSimple(i int) bool {
	if i <= 1 {
		return false
	} else if i <= 3 {
		return true
	} else if i%2 == 0 || i%3 == 0 {
		return false
	}
	n := 5
	for n*n <= i {
		if i%n == 0 || i%(n+2) == 0 {
			return false
		}
		n = n + 6
	}
	return true
}
