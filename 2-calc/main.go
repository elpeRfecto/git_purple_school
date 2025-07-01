package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)



func main() {
	for {
		inputOfUser := userInput()
		userChoice := choice()
		sliceFloat := convertToFloat(strings.Split(inputOfUser, ", "))
		switch userChoice {
			case 1:
				average := sumOfSlice(sliceFloat)
				averageValue := average / float64(len(sliceFloat))
				fmt.Printf("Среднее значение будет - %.2f\n", averageValue)
			case 2:
				summary := sumOfSlice(sliceFloat)
				fmt.Printf("Сумма всех чисел в массиве - %.2f\n", summary)
			case 3:
				median := calcMed(sliceFloat)
				fmt.Printf("Медиана в массиве равна - %.2f\n", median)
			case 4:
				average := sumOfSlice(sliceFloat)
				averageValue := average / float64(len(sliceFloat))
				fmt.Printf("Среднее значение будет - %.2f\n", averageValue)
				summary := sumOfSlice(sliceFloat)
				fmt.Printf("Сумма всех чисел в массиве - %.2f\n", summary)
				median := calcMed(sliceFloat)
				fmt.Printf("Медиана в массиве равна - %.2f\n", median)
			}
		if !checkCalcAgain() {
			break
		}
	}
}

func userInput() string {
	fmt.Println("Введите числа через запятую")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	return input.Text()
}


func convertToFloat(stringSlice []string) []float64 {
	floatSlice := make([]float64, 0, 10)

	for _, value := range stringSlice {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			continue
		}
		floatSlice = append(floatSlice, num)
	}

	return floatSlice
}


func sumOfSlice(floatSlice []float64) float64 {
	sum := 0.0

	for _, value := range floatSlice {
		sum += float64(value)
	}

	return sum
}

func calcMed(floatSlcie []float64) float64 {
	sort.Float64s(floatSlcie)
	if len(floatSlcie) % 2 != 0 {
		return float64(floatSlcie[len(floatSlcie) / 2])
	}

	return ((float64(floatSlcie[len(floatSlcie) / 2 - 1])) + float64(floatSlcie[len(floatSlcie) / 2])) / 2
}

func choice() int {
	fmt.Println("Что нужно сделать?\n1 - вычислить среднее (AVG)\n2 - посчитать сумму (SUM)\n3 - сделать расчет медианы (MED)\n4 - всё сразу")
	// var value int
	// fmt.Scan(&value)
	var value string
	fmt.Scan(&value)
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Нужно ввести число")
	}


	return result
}

func checkCalcAgain() bool {
	fmt.Println("Ёще раз? (y/n)")

	var check string
	fmt.Scan(&check)
	
	if check == "y" || check == "Y" {
		return true
	}
	return false
}