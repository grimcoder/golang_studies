// package main 

// import (
// 	"fmt"
// 	"net/http"
// 	"github.com/gorilla/mux"
// )

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world")
// }

// func main() {
// 	r:=mux.NewRouter()
// 	r.HandleFunc("/", helloWorld)

// 	http.Handle("/", r)
// 	http.ListenAndServe(":8000", nil)
// }