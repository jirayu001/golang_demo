package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//type myHandler func(http.ResponseWriter, *http.Request)

//func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	m(w, r)
//}

type inventory map[string]float64

func (iv inventory) items(w http.ResponseWriter, r *http.Request) {
	for k, v := range iv {
		fmt.Fprintf(w, "%s : %.2f\n", k, v)
	}
}

func (iv inventory) price(w http.ResponseWriter, r *http.Request) {
	searchItem := r.URL.Query().Get("item")
	price, ok := iv[searchItem]
	if !ok {
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "No tiem : %s", searchItem)
		return

	}
	fmt.Fprintf(w, "%.2f\n", price)
}

func (iv inventory) notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "sorry, no such page: %s", r.URL.Path)
}

/*func (iv inventory) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("r.url", r.URL.Path)
	//w.Write([]byte("Hello2"))

	switch r.URL.Path {
	case "/items":
		//w.Write([]byte("handle item"))
		for k, v := range iv {
			fmt.Fprintf(w, "%s : %.2f\n", k, v)
		}
	case "/price":
		//w.Write([]byte("handle price"))
		searchItem := r.URL.Query().Get("item")
		price, ok := iv[searchItem]
		if !ok {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "No tiem : %s", searchItem)
			return

		}
		fmt.Fprintf(w, "%.2f\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "sorry, no such page: %s", r.URL.Path)
	}
}*/
func main() {
	log.SetFlags(0)
	log.Println(os.Getegid())
	inven := inventory{
		"apple":  1.25,
		"orange": 0.99,
	}

	//mux := http.NewServeMux()
	//mux.Handle("/items", http.HandlerFunc(inven.items))
	//mux.Handle("/price", http.HandlerFunc(inven.price))
	//mux.Handle("/", http.HandlerFunc(inven.notFound))

	http.HandleFunc("/items", (inven.items))
	http.HandleFunc("/price", (inven.price))
	http.HandleFunc("/", (inven.notFound))

	//http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", nil)

}
