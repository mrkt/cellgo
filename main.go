// http.service project main.go
package main

import (
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("hello!"))
}

func main() {

	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001", nil)

}

