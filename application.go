package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(Handler))

	go func() {
		cert := "/etc/letsencrypt/live/DOMAIN_NAME/fullchain.pem"
		prvkey := "/etc/letsencrypt/live/DOMAIN_NAME/privkey.pem"
		err := http.ListenAndServeTLS(":https", cert, prvkey, mux)
		if err != nil {
			log.Fatalf("http.ListendAndServeTLS() failed with %s", err)
		}
	}()

	http.ListenAndServe(":8080", http.HandlerFunc(redirect))
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}

// Handler comment
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
