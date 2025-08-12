package forMenu

import (
	"fmt"

	"github.com/ignratnan/products-management/packages/forCheck"
	"github.com/ignratnan/products-management/packages/forInput"
)

func MainMenu() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Main Menu")
	fmt.Println("---")
	fmt.Println("1. Add Product")
	fmt.Println("2. Show Product")
	fmt.Println("3. Edit Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Exit")
	fmt.Println("---")
	inputMain := forInput.InputNum("Please input the number: ")
	switch inputMain {
	case 1:
		AddProduct()
	case 2:
		ShowProduct()
	case 3:
		EditProduct()
	case 4:
		DeleteProduct()
	case 5:
		fmt.Println("Existing the system!")
	}
}

func AddProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Add Product")
	fmt.Println("---")
	inputProduct := forInput.InputProduct("Please input the product name: ")
	inputPrices := forInput.InputPrice("Please input the product price: ")
	inputID := forCheck.NewID()

}

func ShowProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Show Product")
	fmt.Println("---")
}

func EditProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Edit Product")
	fmt.Println("---")
}

func DeleteProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Delete Product")
	fmt.Println("---")
}
