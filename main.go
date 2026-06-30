package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type cacheEntry struct {
	body      []byte
	expiresAt time.Time
}

var cache = map[string]cacheEntry{}

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	entry, exists := cache[key]
	if exists && time.Now().Before(entry.expiresAt) {
		fmt.Println(w, "Cache hit")
		w.Write(entry.body)
		return
	} else {
		fmt.Println(w, "Cache miss")
		response, err := http.Get("http://dummyjson.com/products")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cache[key] = cacheEntry{body: body, expiresAt: time.Now().Add(5 * time.Minute)}
		w.Write(body)
	}
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
