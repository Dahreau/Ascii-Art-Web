package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/ascii-art", AsciiArtHandler)
	mux.HandleFunc("/download", DownloadHandler)

	server := &http.Server{
		Addr:              ":8080",           //adresse du server (le port choisi est à titre d'exemple)
		Handler:           mux,               // listes des handlers
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		ReadTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1 MB // maxinmum de bytes que le serveur va lire
	}
	log.Printf("http://localhost%v", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
