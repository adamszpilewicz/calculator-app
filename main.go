package main

import (
	"log"
	"net/http"
	"os"

	"my-order-app/internal/config"
	"my-order-app/internal/handler"
)

func main() {
	packSizes, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	h := handler.NewHandler(packSizes)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/api/calculate", h.Calculate)
	http.HandleFunc("/api/packsizes", h.GetPackSizes)
	http.HandleFunc("/api/packsizes/add", h.AddPackSize)
	http.HandleFunc("/api/packsizes/delete", h.DeletePackSize)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Printf("Listening on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
