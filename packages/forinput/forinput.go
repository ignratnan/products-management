package forInput

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputNum(text string) int {
	var num int
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	number, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read the number: %w\n", err)
	}

	number = strings.TrimSpace(number)
	num, err = strconv.Atoi(number)
	if err != nil {
		fmt.Printf("Failed to convert the number: %w\n", err)
	}

	return num
}

func InputProduct(text string) string {
	var product string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	product, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read the product: %w", err)
	}

	product = strings.TrimSpace(product)

	return product
}

func InputPrice(text string) float64 {
	var prices float64
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(text)
	price, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read the product: %w", err)
	}

	price = strings.TrimSpace(price)
	prices, err = strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Printf("Failed to convert the price: %w", err)
	}

	return prices
}
