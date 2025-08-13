package forModel

var Products []Product
var FolderPath string = "files"
var FileName string = "products.json"
var Uri string = "http://localhost:8080"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
