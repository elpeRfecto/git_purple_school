package main

import (
	"fmt"
	"strings"
)

// стоимость валюты актульна на 12 июня 2025г. по курсу Google
const USD = 80.00
const EUR = 92.78


// формулы для конвертации
const usdToEu = USD / EUR
const eurToUs = EUR / USD
const usdToRu = USD / 1
const eurToRu = EUR / 1
const rubToUsd = 1 / USD
const rubToEur = 1 / EUR


func main() {
	for {
		res1, res2, res3 := userInput()

		val1, _ := checkUserInput(res1)
		val2, _ := checkUserInput(res3)

		if val1 && val2 {
			fmt.Println(convertCurrency(res1, res2, res3))
			checkCurrency(res1, res3)
		}
		if !checkCovertAgain() {
			break
		}
	}
}

// Считываются исходная валюта, сколько денег надо сконвертировать и в какую валюту пересчитать
func userInput() (string, float64, string) {
	var currencyOriginal string
	var input float64
	var currencyTarget string

	fmt.Println("Какая у вас валюта? (USD, EUR, RUB)")
	fmt.Scan(&currencyOriginal)

	fmt.Println("Сколько денег нужно перевести в другую валюту?")
	fmt.Scan(&input)

	fmt.Println("В какую валюту конвертировать? (USD, EUR, RUB)")
	fmt.Scan(&currencyTarget)
	
	return strings.ToUpper(currencyOriginal), input, strings.ToUpper(currencyTarget)
}

// Здесь происходит расчет 
func convertCurrency(curOrig string, input float64, curTar string) string {
	var result string

	switch curOrig {
	case "USD":
		if curTar == "EUR" {
				result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * usdToEu)
		} else {
			result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * usdToRu)
		}
	case "EUR":
		if curTar == "USD" {
			result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * eurToUs)
		} else {
			result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * eurToRu)
		}
	case "RUB":
		if curOrig == "USD" {
			result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * rubToUsd)
		} else {
			result = fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input * rubToEur)
		}
	}
	return result
}

// функция для выхода из бесконечного цикла в функции main
func checkCovertAgain() bool {
	fmt.Println("Сделать еще расчет? (y/n)")

	var check string
	fmt.Scan(&check)
	
	if check == "y" || check == "Y" {
		return true
	}
	return false
}

// функция для подсказок по валютам (из какой в какую)
func checkCurrency(currencyOrig string, currencyTar string) {
	if currencyOrig == currencyTar {
		fmt.Printf("Нельзя сконвертировать из %v в %v\n", currencyOrig, currencyTar)
		switch currencyOrig {
		case "USD":
			fmt.Printf("%v можно пересчитать в: EUR или RUB\n", currencyOrig)
		case "EUR":
			fmt.Printf("%v можно пересчитать в: USD или RUB\n", currencyOrig)
		case "RUB":
			fmt.Printf("%v можно пересчтать в: USD или EUR\n", currencyOrig)
		}
	}
}

func checkUserInput(currency string) (bool, string) {
	switch currency {
	case "USD": 
		return true, ""
	case "EUR":
		return true, ""
	case "RUB":
		return true, ""
	default:
		fmt.Printf("Нет такой валюты как - %v\n", currency)
	}
	
	return false, ""
}