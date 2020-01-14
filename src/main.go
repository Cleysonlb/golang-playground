package main

import "service"
import "net/http"
// import "fmt"
import "log"
// import "html"

func main() {

  http.HandleFunc("/bar", func(write http.ResponseWriter, request *http.Request) {
    service.Service(write)
  })
  
  log.Fatal(http.ListenAndServe(":3000", nil))
}