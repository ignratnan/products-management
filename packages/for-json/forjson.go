package forjson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func WriteJson(folderpath string, filename string, data interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Errorf("failed to marshal the data: %w", err)
	}

	err = os.WriteFile(fullpath, jsonData, 0644)
	if err != nil {
		fmt.Errorf("failed to write the file: %w", err)
	}

	return nil
}

func ReadJson(folderpath string, filename string, target interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := os.ReadFile(fullpath)
	if err != nil {
		fmt.Errorf("failed to read the file: %w", err)
	}

	err = json.Unmarshal(jsonData, target)
	if err != nil {
		fmt.Errorf("failed to unmarshal the data: %w", err)
	}

	return nil
}

func mainMenu() {
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
	inputMain := inputNum("Please input the number: ")
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
