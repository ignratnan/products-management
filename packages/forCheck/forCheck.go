package forCheck

import (
	"github.com/ignratnan/products-management/packages/forModel"
)

func LastID() int {
	return 0
}

func NewID() int {
	var nextID int
	if len(forModel.Products) > 0 {
		nextID = forModel.Products[len(forModel.Products)-1].ID + 1
	} else {
		nextID = 1
	}
	return nextID
}
