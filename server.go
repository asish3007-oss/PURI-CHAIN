package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MineRequest struct {
	Data       string json:"data"
	Difficulty int    json:"difficulty"
}

func handleMine(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var req MineRequest
	json.Unmarshal(body, &req)

	addBlock(req.Data, req.Difficulty)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

func handleChain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Blockchain)
}

func startServer() {
	http.HandleFunc("/mine", handleMine)
	http.HandleFunc("/chain", handleChain)

	log.Println("🔥 PURI-CHAIN node running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
