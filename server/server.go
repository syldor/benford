package main

import (
    "net/http"
    "fmt"
    "benford"
    "encoding/json"
)

func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("public")))
    fmt.Println("Server started and listening to port 3001")
    http.ListenAndServe(":3001", nil)
}

func GenerateMarkdown(res http.ResponseWriter, req *http.Request) {
	json, _ := json.Marshal(benford.Calculate_from_string(req.FormValue("body")))
	res.Write([]byte(json))
	fmt.Println(benford.Calculate_from_string(req.FormValue("body")))
}