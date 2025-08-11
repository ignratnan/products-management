package formenu

import (
	"fmt"

	"github.com/ignratnan/products-management/packages/forInput"
)

func MainMenu() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("Main Menu")
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
		addProduct()
	case 2:
		showProduct()
	case 3:
		editProduct()
	case 4:
		deleteProduct()
	case 5:
		fmt.Println("Existing the system!")
	}
}
