package main

import "service"
import "net/http"
import "log"

func main() {

  http.HandleFunc("/", func(write http.ResponseWriter, request *http.Request) {
    service.Service(write)
  })
  
  log.Fatal(http.ListenAndServe(":3000", nil))
}