package forMenu

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ignratnan/products-management/packages/forApi"
	"github.com/ignratnan/products-management/packages/forCheck"
	"github.com/ignratnan/products-management/packages/forInput"
	"github.com/ignratnan/products-management/packages/forJson"
	"github.com/ignratnan/products-management/packages/forModel"
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
	fmt.Println("5. Run API")
	fmt.Println("6. Exit")
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
		RunApi()
	case 6:
		fmt.Println("Existing the system!")
	}
}

func AddProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Add Product")
	fmt.Println("---")
	var newProduct forModel.Product
	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products)
	newProduct.Name = forInput.InputProduct("Please input the product name: ")
	newProduct.Price = forInput.InputPrice("Please input the product price: ")
	newProduct.ID = forCheck.NewID()
	forModel.Products = append(forModel.Products, newProduct)
	forJson.WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
	fmt.Println("---")
	fmt.Println("The new product successfully added.")
	forInput.InputProduct("Click enter to back to Main Menu!")
	MainMenu()
}

func ShowProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Show Product")
	fmt.Println("---")
	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products)
	println(forModel.Products)
	forJson.ShowData()
	fmt.Println("---")
	forInput.InputProduct("Click enter to back to Main Menu!")
	MainMenu()
}

func EditProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Edit Product")
	fmt.Println("---")

	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products)
	forJson.ShowData()
	fmt.Println("---")
	forJson.EditData()
	fmt.Println("---")
	fmt.Println("The product successfully edited.")
	forInput.InputProduct("Click enter to back to Main Menu!")
	MainMenu()
}

func DeleteProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Delete Product")
	fmt.Println("---")
	forJson.DeleteData()
	fmt.Println("---")
	fmt.Println("The product successfully edited.")
	forInput.InputProduct("Click enter to back to Main Menu!")
	MainMenu()
}

func RunApi() {
	http.HandleFunc("/products", forApi.ProductsHandler)
	http.HandleFunc("/products/", forApi.ProductHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
