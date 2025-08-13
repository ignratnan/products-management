package forRequest

import (
	"log"
	"net/http"

	"github.com/ignratnan/products-management/packages/forModel"
)

func ForGet(sub string) {
	resp, err := http.Get(forModel.Uri + sub)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

}
