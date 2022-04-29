package function

import (
	"fmt"
	"net/http"
)

func EntryPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
}
