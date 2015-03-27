package main

import (
    "net/http"
    "fmt"
    "benford"
    "io"
    "encoding/json"
)

func main() {
    http.HandleFunc("/benford", computeBenford)
    http.Handle("/", http.FileServer(http.Dir("front")))
    fmt.Println("Server started and listening to port 3001")
    http.ListenAndServe(":3001", nil)
}

func DecodeJson(r *http.Request, o interface{}) (error) {
    raw := make([]byte, r.ContentLength)
    _, err := io.ReadFull(r.Body, raw)
    if err == nil {
        err = json.Unmarshal(raw, &o)
    }
    return err;
}

func computeBenford(res http.ResponseWriter, req *http.Request) {
    type Data struct {
        Body string 
    }
    var d Data
    if err := DecodeJson(req, &d); err != nil {
    }
	json, _ := json.Marshal(benford.Calculate_from_string(d.Body))
    res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}
