package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Backend struct {
	URL    string
	Weight int
	Alive  bool
}

var (
	backends = []*Backend{
		{URL: "http://backend_potente:8080", Weight: 3, Alive: true},
		{URL: "http://backend_medio:8080", Weight: 2, Alive: true},
		{URL: "http://backend_fraco:8080", Weight: 1, Alive: true},
	}
	serverPool []*Backend
	mu         sync.Mutex
	index      int
)

func init() {
	for _, b := range backends {
		for i := 0; i < b.Weight; i++ {
			serverPool = append(serverPool, b)
		}
	}
}

func healthCheck() {
	for {
		time.Sleep(5 * time.Second)
		for _, b := range backends {
			resp, err := http.Get(b.URL)
			mu.Lock()
			if err != nil || resp.StatusCode != http.StatusOK {
				b.Alive = false
			} else {
				b.Alive = true
			}
			mu.Unlock()
		}
	}
}

func getNextServer() *Backend {
	mu.Lock()
	defer mu.Unlock()

	for i := 0; i < len(serverPool); i++ {
		server := serverPool[index]
		index = (index + 1) % len(serverPool)
		if server.Alive {
			return server
		}
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	target := getNextServer()
	
	if target == nil {
		http.Error(w, "Erro: 503 - Todos os servidores estao offline (Health Check falhou)", http.StatusServiceUnavailable)
		return
	}

	resp, err := http.Get(target.URL)
	if err != nil {
		http.Error(w, "Erro ao contatar backend", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Fprint(w, string(body))
}

func main() {
	fmt.Println("=== Load Balancer WRR com Health Check Ativo na porta :8000 ===")
	
	go healthCheck()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}