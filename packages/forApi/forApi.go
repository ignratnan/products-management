package forApi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/ignratnan/products-management/packages/forCheck"
	"github.com/ignratnan/products-management/packages/forJson"
	"github.com/ignratnan/products-management/packages/forModel"
)

var mutex = &sync.Mutex{}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	forJson.ReadJson(forModel.FolderPath, forModel.FileName, &forModel.Products)
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(forModel.Products)
		log.Println("GET /products - All items sent")
	case http.MethodPost:
		var newProduct forModel.Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		nextID := forCheck.NewID()
		newProduct.ID = nextID
		forModel.Products = append(forModel.Products, newProduct)

		err := forJson.WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
		if err != nil {
			log.Printf("Error saving to file: %v", err)
			http.Error(w, "Failedto save date", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newProduct)
		log.Printf("POST /products - New product created with ID %d and save", newProduct.ID)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		for _, product := range forModel.Products {
			if product.ID == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(product)
				return
			}
		}
		http.Error(w, "Product not found", http.StatusNotFound)
	case http.MethodPut:
		var updatedProduct forModel.Product
		err = json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		for i, product := range forModel.Products {
			if product.ID == id {
				forModel.Products[i] = updatedProduct
				forModel.Products[i].ID = id

				err = forJson.WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
				if err != nil {
					log.Printf("Error saving to file: %v", err)
					http.Error(w, "Failed to save data", http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(forModel.Products[i])
				log.Printf("PUT /products/%d - Item updated and saved", id)
				return
			}
		}
		http.Error(w, "Item not found", http.StatusNotFound)
	case http.MethodDelete:
		for i, product := range forModel.Products {
			if product.ID == id {
				forModel.Products = append(forModel.Products[:i], forModel.Products[i+1:]...)

				err = forJson.WriteJson(forModel.FolderPath, forModel.FileName, forModel.Products)
				if err != nil {
					log.Printf("Error saving to file: %v", err)
					http.Error(w, "Failed to save data", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusNoContent)
				log.Printf("DELETE /products/%d - Item deleted and saved", id)
				return
			}
		}
		http.Error(w, "Item not found", http.StatusNotFound)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
