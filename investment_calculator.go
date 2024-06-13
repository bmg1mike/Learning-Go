package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	writeBalanceToFile();

	fmt.Print("Investment Aount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	errors.New("Failed to read file")


	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func getBalanceFromFile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return balance
}

func writeBalanceToFile() {
	balance := 1000;
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
