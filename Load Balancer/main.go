package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"sync"
	"time"
)

// LoadBalancer represents a load balancer with multiple backend servers
type LoadBalancer struct {
	servers []*Server
	mu      sync.Mutex
}

// Server represents a backend server with a URL and a health check
type Server struct {
	URL     *url.URL
	Healthy bool
}

func newLoadBalancer(servers []*Server) *LoadBalancer {
	return &LoadBalancer{servers: servers}
}

func (lb *LoadBalancer) getHealthyServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	for _, server := range lb.servers {
		if server.Healthy {
			return server
		}
	}
	return nil
}

func (lb *LoadBalancer) healthCheck() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	startTime := time.Now()
	for range ticker.C {
		// Clear the terminal screen
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Print the running time
		runningTime := time.Since(startTime)
		hours := int(runningTime.Hours())
		minutes := int(runningTime.Minutes()) % 60
		seconds := int(runningTime.Seconds()) % 60
		fmt.Printf("Running since: %d:%02d:%02d\n", hours, minutes, seconds)

		// Print the server status
		for _, server := range lb.servers {
			status := "active"
			if !server.Healthy {
				status = "deactive"
			}
			fmt.Printf("%s: %s\n", server.URL.Host, status)
		}

		// Check the health of each server
		for _, server := range lb.servers {
			resp, err := http.Get(server.URL.String() + "/health")
			if err != nil {
				log.Printf("Error checking health of %s: %v", server.URL, err)
				server.Healthy = false
			} else if resp.StatusCode != http.StatusOK {
				log.Printf("Health check failed for %s: %d", server.URL, resp.StatusCode)
				server.Healthy = false
			} else {
				server.Healthy = true
			}
		}
	}
}

func main() {
	addServers := []*Server{
		{URL: &url.URL{Scheme: "http", Host: "localhost:3001"}, Healthy: true},
		{URL: &url.URL{Scheme: "http", Host: "localhost:3002"}, Healthy: true},
	}
	deleteServers := []*Server{
		{URL: &url.URL{Scheme: "http", Host: "localhost:3003"}, Healthy: true},
		{URL: &url.URL{Scheme: "http", Host: "localhost:3004"}, Healthy: true},
	}
	completeServers := []*Server{
		{URL: &url.URL{Scheme: "http", Host: "localhost:3005"}, Healthy: true},
		{URL: &url.URL{Scheme: "http", Host: "localhost:3006"}, Healthy: true},
	}

	addLB := newLoadBalancer(addServers)
	deleteLB := newLoadBalancer(deleteServers)
	completeLB := newLoadBalancer(completeServers)

	go addLB.healthCheck()
	go deleteLB.healthCheck()
	go completeLB.healthCheck()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			server := addLB.getHealthyServer()
			if server == nil {
				http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
				return
			}
			httputil.NewSingleHostReverseProxy(server.URL).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			server := addLB.getHealthyServer()
			if server == nil {
				http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
				return
			}
			httputil.NewSingleHostReverseProxy(server.URL).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			server := deleteLB.getHealthyServer()
			if server == nil {
				http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
				return
			}
			httputil.NewSingleHostReverseProxy(server.URL).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/complete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			server := completeLB.getHealthyServer()
			if server == nil {
				http.Error(w, "No healthy servers available", http.StatusServiceUnavailable)
				return
			}
			httputil.NewSingleHostReverseProxy(server.URL).ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Load Balancer started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
