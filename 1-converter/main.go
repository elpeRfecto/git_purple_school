package main

import (
	"fmt"
	"strings"
)

// стоимость валюты актульна на 24 июня 2025г. по курсу Google
const USD = 78.38
const EUR = 91.14


// формулы для конвертации
const usdToEu = USD / EUR
const usdToRu = USD / 1

const eurToUs = EUR / USD
const eurToRu = EUR / 1

const rubToUsd = 1 / USD
const rubToEur = 1 / EUR


var allowedCurrencies = map[string]bool{
    "USD": true,
    "EUR": true,
    "RUB": true,
}

var conversionRates = map[string]map[string]float64{
    "USD": {
        "EUR": usdToEu,
        "RUB": usdToRu,
    },
    "EUR": {
        "USD": eurToUs,
        "RUB": eurToRu,
    },
    "RUB": {
        "USD": rubToUsd,
        "EUR": rubToEur,
    },
}


func main() {
	for {
		res1, res2, res3 := userInput()

		val1, _ := checkUserInput(res1)
		val2, _ := checkUserInput(res3)

		if val1 && val2 {
			fmt.Println(convertCurrency(res1, res2, res3))
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
	if curOrig == curTar {
        return fmt.Sprintf("Нельзя сконвертировать из %v в %v", curOrig, curTar)
    }
    if rates, ok := conversionRates[curOrig]; ok {
        if rate, ok := rates[curTar]; ok {
            return fmt.Sprintf("%.2f %s в %s - %.2f", input, curOrig, curTar, input*rate)
        }
    }
    return "Конвертация невозможна"
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

func checkUserInput(currency string) (bool, string) {
	if allowedCurrencies[currency] {
        return true, ""
    }
    fmt.Printf("Нет такой валюты как - %v\n", currency)
    return false, ""
}