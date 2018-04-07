package main
import (
    "fmt"
    "net"
    "net/http"
    "net/http/fcgi"

)

func main() {
    // http.HandleFunc("/", handler)
    // http.ListenAndServe(":8080", nil)
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        return
    }
    http.HandleFunc("/", handler)
    fcgi.Serve(l, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}
