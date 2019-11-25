package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

var stringFromFile string

func main() {
    buf, err := ioutil.ReadFile("/data/foo")
    if err != nil {
        fmt.Println("COULD NOT READ /data/foo see: " + err.Error())
    } else {
        stringFromFile = string(buf)
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Hello from %s! --- Content from /data/foo: %s --- Args: %v --- Env: %v", host, stringFromFile, os.Args, os.Environ())
}
