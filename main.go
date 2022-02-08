package main

import (
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
)

func HandleGetMethod(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Message from Method: GET")
}

func HandlePutMethod(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Message from Method: PUT")
}

func HandlePostMethod(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Message from Method: POST")
}

func HandleDeleteMethod(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Message from Method: DELETE")
}

func main() {
        r := mux.NewRouter()
        r.HandleFunc("/", HandleGetMethod).Methods(http.MethodGet)
        r.HandleFunc("/", HandlePostMethod).Methods(http.MethodPost)
        r.HandleFunc("/", HandlePutMethod).Methods(http.MethodPut)
        r.HandleFunc("/", HandleDeleteMethod).Methods(http.MethodDelete)

        srv := http.Server{
                Addr:    ":8081",
                Handler: r,
        }

        log.Println("Listening...")
        srv.ListenAndServe()
}
