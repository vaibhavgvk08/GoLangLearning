package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/test", YourHandler)

    log.Fatal(http.ListenAndServe(":8000", r))
}
