package main

import (
	"fmt"
	"io"
	"net/http"
)
var cache = map[string][]byte{}
func handler(w http.ResponseWriter, r *http.Request) {
	key:=r.URL.Path
	cachedBody,exists := cache[key]
	if exists{
		fmt.Fprintln(w, "Cache hit")
		w.Write(cachedBody)
		return
	}else{
		fmt.Fprintln(w, "Cache miss")
		response,err :=http.Get("http://dummyjson.com/products")
		if err!=nil{
			fmt.Println("Error:", err)
			return
		}
		body,err :=io.ReadAll(response.Body)
		if err!=nil{
			fmt.Println( "Error:", err)
			return
		}
		cache[key]=body
		w.Write(body)
	}
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	err:=http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
