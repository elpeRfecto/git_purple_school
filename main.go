package main

import "fmt"

func main() {
	const USD = 77.50
	const EUR = 88.02
	fmt.Println("Сколько денег нужно перевести в другую валюту?")
	var input int
	fmt.Scan(&input)
	usdToEur := float64(input) * USD / EUR
	usdToRub := float64(input) * USD / 1
	eurToRub := float64(input) * EUR / 1
	fmt.Printf("%v USD в EUR - %.2f\n", input, usdToEur)
	fmt.Printf("%v USD в RUB - %.2f\n", input, usdToRub)
	fmt.Printf("%v EUR в RUB - %.2f\n", input, eurToRub)
}
