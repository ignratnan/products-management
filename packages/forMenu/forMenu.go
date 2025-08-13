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

	inputMain := forInput.InputNum("Please input the number: ") //For input selected number
	switch inputMain {
	case 1:
		AddProduct() //Show "Add Product"
	case 2:
		ShowProduct() //Show "Show Product"
	case 3:
		EditProduct() //Show "Edit Product"
	case 4:
		DeleteProduct() //Show "Delete Product"
	case 5:
		RunApi() //Run the server
	case 6:
		fmt.Println("Existing the system!") //Exit the system
	}
}

func AddProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Add Product")
	fmt.Println("---")

	var newProduct forModel.Product                                              //Declare variable to store new product
	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products) //Read JSON file and safe the data to Products
	newProduct.Name = forInput.InputProduct("Please input the product name: ")   //For input product name
	newProduct.Price = forInput.InputPrice("Please input the product price: ")   //For input product price
	newProduct.ID = forCheck.NewID()                                             //For check new ID
	forModel.Products = append(forModel.Products, newProduct)                    //Add new product to Products
	forJson.WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products) //Write all data to JSON file

	fmt.Println("---")
	fmt.Println("The new product successfully added.")
	forInput.InputProduct("Click enter to back to Main Menu!") //To stop the code before back to Main Menu
	MainMenu()
}

func ShowProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Show Product")
	fmt.Println("---")

	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products) //Read JSON file and safe the data to Products
	forJson.ShowData()                                                           //To show all product

	fmt.Println("---")
	forInput.InputProduct("Click enter to back to Main Menu!") //To stop the code before back to Main Menu
	MainMenu()
}

func EditProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Edit Product")
	fmt.Println("---")

	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products) //Read JSON file and safe the data to Products
	forJson.ShowData()                                                           //To show all product
	fmt.Println("---")
	forJson.EditData() //To edit the data

	fmt.Println("---")
	fmt.Println("The product successfully edited.")
	forInput.InputProduct("Click enter to back to Main Menu!") //To stop the code before back to Main Menu
	MainMenu()
}

func DeleteProduct() {
	fmt.Println("#########--Product Management--#########")
	fmt.Println("---")
	fmt.Println("*Delete Product")
	fmt.Println("---")
	forJson.DeleteData() //To delete the data
	fmt.Println("---")
	fmt.Println("The product successfully edited.")
	forInput.InputProduct("Click enter to back to Main Menu!") //To stop the code before back to Main Menu
	MainMenu()
}

func RunApi() {
	http.HandleFunc("/products", forApi.ProductsHandler) //Declare handle function for sub URI /products
	http.HandleFunc("/products/", forApi.ProductHandler) //Declare handle function for sub URI /products/

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) //To run the server
}
