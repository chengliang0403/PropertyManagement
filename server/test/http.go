package main

import (
	"fmt"
	"net/http"
	"strings"
	// "log"
)

func sayHi(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Hi Hi")
}

// func main()  {
//   http.HandleFunc("/", sayHi)
//   err := http.ListenAndServe(":9090", nil)
//   if err != nil {
//     log.Fatal("ListenAndServe:", err)
//   }
// }
