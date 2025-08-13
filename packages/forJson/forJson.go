package forJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ignratnan/products-management/packages/forInput"
	"github.com/ignratnan/products-management/packages/forModel"
)

func WriteJson(folderpath string, filename string, data interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal the data: %w", err)
	}

	err = os.WriteFile(fullpath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write the file: %w", err)
	}

	return nil
}

func ReadJson(folderpath string, filename string, target interface{}) error {
	fullpath := filepath.Join(folderpath, filename)

	jsonData, err := os.ReadFile(fullpath)
	if err != nil {
		return fmt.Errorf("failed to read the file: %w", err)
	}

	err = json.Unmarshal(jsonData, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the data: %w", err)
	}

	return nil
}

func ShowData() {
	fmt.Println("All Products:")
	for _, product := range forModel.Products {
		fmt.Printf("ID: %d\tName: %s\tPrice: %.2f\n", product.ID, product.Name, product.Price)
	}
}

func EditData() {
	var editProducts []forModel.Product
	editID := forInput.InputNum("Please input the ID to edit: ")
	editName := forInput.InputProduct("Please input the product name: ")
	editPrice := forInput.InputPrice("Please input the product price: ")

	for _, product := range forModel.Products {
		if product.ID == editID {
			product.Name = editName
			product.Price = editPrice

		}
		editProducts = append(editProducts, product)
	}

	forModel.Products = editProducts
	WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
}

func DeleteData() {
	var deleteProducts []forModel.Product
	ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products)
	ShowData()
	fmt.Println("---")
	deleteID := forInput.InputNum("Please input the ID to edit: ")
	for _, product := range forModel.Products {
		if product.ID != deleteID {
			deleteProducts = append(deleteProducts, product)
		}
	}

	forModel.Products = deleteProducts
	WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
}
