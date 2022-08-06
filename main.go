package main

import (
	"fmt"
	"net/http"
	"os"
   "encoding/json"
)

func main() {
   http.HandleFunc("/", Index)
   http.HandleFunc("/status", Status)

   http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request){
   fmt.Fprintf(w, "Olá Devs Norte!\n\nAmbiente de execução: [%s]", os.Getenv("ENVIRONMENT"))
}

func Status(w http.ResponseWriter, r *http.Request){
   w.WriteHeader(200)
   w.Header().Set("Content-Type","application/json")
   resp := make(map[string]string)

   resp["status"] = "UP"

   jsonResp, _ := json.Marshal(resp)
   w.Write(jsonResp)
}
