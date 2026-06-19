package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)


type Backend struct {
	URL    *url.URL
	Weight int
	Alive  bool
	mux    sync.RWMutex
}


func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}


func (b *Backend) IsAlive() bool {
	b.mux.RLock()
	alive := b.Alive
	b.mux.RUnlock()
	return alive
}


type ServerPool struct {
	backends []*Backend
	sequence []*Backend
	current  uint64
}


func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, 1) % uint64(len(s.sequence)))
}


func (s *ServerPool) GetNextPeer() *Backend {
	next := s.NextIndex()
	l := len(s.sequence) + next


	for i := next; i < l; i++ {
		idx := i % len(s.sequence)
		if s.sequence[idx].IsAlive()
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return s.sequence[idx]
		}
	}
	return nil
}


func (s *ServerPool) HealthCheck() {
	client := http.Client{
		Timeout: 2 * time.Second,
	}

	for _, b := range s.backends {
		healthURL := b.URL.String() + "/health"
		resp, err := client.Get(healthURL)

		if err != nil || resp.StatusCode != http.StatusOK {
			if b.IsAlive() {
				b.SetAlive(false)
				log.Printf("[HEALTH CHECK] ALERTA: Nó caído detectado! Removendo do tráfego: %s", b.URL)
			}
		} else {
			if !b.IsAlive() {
				b.SetAlive(true)
				log.Printf("[HEALTH CHECK] SUCESSO: Nó se recuperou e voltou a ficar ativo: %s", b.URL)
			}
			resp.Body.Close()
		}
	}
}


func lbHandler(pool *ServerPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		peer := pool.GetNextPeer()
		if peer != nil {

			proxy := httputil.NewSingleHostReverseProxy(peer.URL)


			r.Header.Set("X-Forwarded-Host", r.Host)
			r.Header.Set("X-Custom-LoadBalancer", "Weighted-Round-Robin")

			proxy.ServeHTTP(w, r)
			return
		}


		http.Error(w, "Erro 503: Todos os nós de destino estão indisponíveis no momento.", http.StatusServiceUnavailable)
	}
}

func main() {

	nodes := []struct {
		address string
		weight  int
	}{
		{"http://backend_potente:8080", 3},
		{"http://backend_medio:8080", 2},
		{"http://backend_fraco:8080", 1},
	}

	pool := ServerPool{}

	for _, node := range nodes {
		parsedURL, err := url.Parse(node.address)
		if err != nil {
			log.Fatalf("Erro crítico ao parsear URL do nó: %v", err)
		}

		backendNode := &Backend{
			URL:    parsedURL,
			Weight: node.weight,
			Alive:  true,
		}

		pool.backends = append(pool.backends, backendNode)


		for i := 0; i < node.weight; i++ {
			pool.sequence = append(pool.sequence, backendNode)
		}
	}


	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			pool.HealthCheck()
		}
	}()


	serverPort := ":8000"
	server := &http.Server{
		Addr:    serverPort,
		Handler: lbHandler(&pool),
	}

	log.Printf("=== Custom Load Balancer (Cenário B - WRR) ativo na porta %s ===", serverPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Falha crítica no Load Balancer: %v", err)
	}
}
