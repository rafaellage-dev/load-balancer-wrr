package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Backend struct {
	URL    string
	Weight int
}

var (
	backends = []Backend{
		{URL: "http://backend_potente:8080", Weight: 3},
		{URL: "http://backend_medio:8080", Weight: 2},
		{URL: "http://backend_fraco:8080", Weight: 1},
	}
	serverPool []string
	mu         sync.Mutex
	index      int
)

func init() {
	for _, b := range backends {
		for i := 0; i < b.Weight; i++ {
			serverPool = append(serverPool, b.URL)
		}
	}
}

func getNextServer() string {
	mu.Lock()
	defer mu.Unlock()
	server := serverPool[index]
	index = (index + 1) % len(serverPool)
	return server
}

func handler(w http.ResponseWriter, r *http.Request) {
	target := getNextServer()

	resp, err := http.Get(target)
	if err != nil {
		http.Error(w, "Erro ao contatar backend", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Fprint(w, string(body))
}

func main() {
	fmt.Println("=== Custom Load Balancer (Cenario B - WRR) ativo na porta :8000 ===")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}