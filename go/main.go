package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


// Response define a estrutura do JSON que será retornado pelo endpoint
type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

// projetoKorpHandler processa as requisições da rota /projeto-korp
func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	// Garante que o endpoint responda apenas a requisições GET
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"erro": "Método não permitido"})
		return
	}

	// Cria o objeto de resposta com o horário atual dinâmico em UTC
	res := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339), // Formato padrão ISO 8601 / RFC3339
	}

	// Define os cabeçalhos da resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Codifica e envia o JSON
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Erro ao codificar resposta JSON: %v", err)
	}
}

func main() {
	// Registra o handler para o endpoint solicitado
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Servidor http-server-projeto-korp iniciado com sucesso na porta 8080...")
	
	// Inicia o servidor na porta 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
