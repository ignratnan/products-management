package main

import (
	"os"

	"github.com/ignratnan/products-management/packages/forMenu"
	"github.com/ignratnan/products-management/packages/forModel"
)

func main() {
	os.MkdirAll(forModel.FolderPath, 0755)
	forMenu.MainMenu()
}
