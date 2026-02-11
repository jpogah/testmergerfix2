package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var helloPage = template.Must(template.New("hello").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>Hello Web</title>
</head>
<body>
  <main>
    <h1>Hello, world!</h1>
  </main>
</body>
</html>`))

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := helloPage.Execute(w, nil); err != nil {
			log.Printf("failed to render hello page: %v", err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := ":" + port
	log.Printf("hello-web listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
