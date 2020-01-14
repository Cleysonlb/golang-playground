package service

import "fmt"
import "net/http"
import "encoding/json"

type objeto struct {
	Name string
	Age int
}

func Service(write http.ResponseWriter) {
	person := objeto{
		Name:  "Cleyson",
		Age: 27,
	}

	ret, _ := json.Marshal(person)

	fmt.Fprintf(write, string(ret))
}