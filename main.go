package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
   http.HandleFunc("/", Index)
   http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Olá Devs Norte!\n\nAmbiente de execução: [%s]", os.Getenv("ENVIRONMENT"))
}

