package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
	"sync"
)

type cacheEntry struct {
	body      []byte
	expiresAt time.Time
}

var cache = map[string]cacheEntry{}
var mu sync.RWMutex

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	mu.RLock()
	entry, exists := cache[key]
	mu.RUnlock()
	if exists && time.Now().Before(entry.expiresAt) {
		fmt.Printf("%s %s Cache hit\n", r.Method, r.URL.Path)
		w.Write(entry.body)
		return
	} else {
		fmt.Printf("%s %s Cache miss\n", r.Method, r.URL.Path)
		baseURL := "http://dummyjson.com"
		fullURL := baseURL + key
		response, err := http.Get(fullURL)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		if(response.StatusCode==200){
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			mu.Lock()
			cache[key] = cacheEntry{body: body, expiresAt: time.Now().Add(5 * time.Minute)}
			mu.Unlock()
			w.WriteHeader(response.StatusCode)
			w.Write(body)
		}else{
			w.WriteHeader(response.StatusCode)
	    	w.Write(body)
		}
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
