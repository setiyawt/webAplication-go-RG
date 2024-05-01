package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path) // membuka file
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string // membaca tiap baris

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(lines) == 0 {
		return []string{}, nil
	}

	return lines, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	var income, expense int
	var lastDate string

	for _, transaction := range data {
		parts := strings.Split(transaction, ";")
		if len(parts) != 3 {
			continue
		}

		date, transactionType, amountStr := parts[0], parts[1], parts[2]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			continue
		}

		if transactionType == "income" {
			income += amount
		} else if transactionType == "expense" {
			expense += amount
		}

		lastDate = date
	}

	var result string
	if income > expense {
		result = lastDate + ";profit;" + strconv.Itoa(income-expense)
	} else {
		result = lastDate + ";loss;" + strconv.Itoa(expense-income)
	}

	return result
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
