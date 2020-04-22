package main



import (

    "fmt"

    "net/http"

)



func main() {

    http.HandleFunc("/ping", PingServer)

    http.HandleFunc("/", HelloServer)

    http.ListenAndServe(":8080", nil)

}



func PingServer(w http.ResponseWriter, r *http.Request) {

fmt.Fprintf(w, "backend: Pong")

}



func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello  there %s", r.URL.Path[1:] )
}
