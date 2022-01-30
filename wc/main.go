package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func wordCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wordCount := len(strings.Fields(string(body)))
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("words: %d\n", wordCount)))
}

func timeMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		log.Printf("%s %s %s\n", r.Method, r.URL, time.Since(start))
	}
}

func main() {
	http.HandleFunc("/", timeMiddleware(wordCountHandler))
	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
