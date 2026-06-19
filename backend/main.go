package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Captura o nome do nó configurado nas variáveis de ambiente do Docker
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "" {
		serverName = "No_Heterogeneo_Desconhecido"
	}

	// Porta padrão interna do container
	port := ":8080"

	// Rota Principal: Retorna o ID/Nome do container para validação de redirecionamento
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Recebido tráfego do Balanceador no nó: %s", serverName)
		_, err := fmt.Fprintf(w, "=== Resposta do Servidor Backend ===\nIdentificação do Nó: %s\nStatus: Processado com Sucesso\n", serverName)
		if err != nil {
			log.Printf("Erro ao responder requisição: %v", err)
		}
	})

	// REQUISITO: Mecanismo de Health Check para o balanceador validar a saúde do nó
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintln(w, "Healthy")
	})

	log.Printf("Iniciando Nó de Destino [%s] na porta %s...", serverName, port)
	// Tratamento de erro nativo ao iniciar o servidor HTTP
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Falha crítica ao iniciar o servidor backend: %v", err)
	}
}
