package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	addAndSelectService, _ := url.Parse("http://localhost:3001")
	deleteService, _ := url.Parse("http://localhost:3002")
	completeService, _ := url.Parse("http://localhost:3003")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			httputil.NewSingleHostReverseProxy(addAndSelectService).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httputil.NewSingleHostReverseProxy(addAndSelectService).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			httputil.NewSingleHostReverseProxy(deleteService).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/complete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			httputil.NewSingleHostReverseProxy(completeService).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Load Balancer started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
