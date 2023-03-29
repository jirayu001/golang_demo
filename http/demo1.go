package main

import (
	"log"
	"net/http"
	"os"
)

type myHandler func()

func (myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
func inventory() {

}
func main() {
	log.SetFlags(0)
	log.Println(os.Getegid())
	http.ListenAndServe(":8080", myHandler(func() {

	}))
	//http.ListenAndServe(":8080", myHandler(inventory))

}
