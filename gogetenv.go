package main

import (
    "fmt"
    "net/http"
    "strings"
    "log"
	"os"
)

func printEnv(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // get arguments
    fmt.Println(r.Form)  // form info
    fmt.Println("path", r.URL.Path)
    fmt.Println(r.Form["url_long"])
	
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
	
	fmt.Println()
    for _, e := range os.Environ() {
        fmt.Fprintln(w, e)
    }
}

func main() {
    http.HandleFunc("/", printEnv)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}