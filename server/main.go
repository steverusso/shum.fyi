// This is free and unencumbered software released
// into the public domain. Please see the UNLICENSE
// file or unlicense.org for more information.
package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

//go:embed public
var website embed.FS

func main() {
	host := flag.String("host", "localhost", "The web host.")
	port := flag.Int("p", 0, "The web port. If specified, TLS will not be used.")
	flag.Parse()

	website, err := fs.Sub(website, "public")
	if err != nil {
		log.Fatalf("couldn't get 'public' subdir of embedded fs: %v", err)
	}

	h := http.FileServer(http.FS(website))

	if *port != 0 {
		addr := fmt.Sprintf("%s:%d", *host, *port)
		if err := http.ListenAndServe(addr, h); err != nil {
			log.Fatalf("couldn't listen and serve: %v", err)
		}
	} else {
		// Serve over HTTPS using the `autocert` package.
		go func() {
			if err := http.Serve(autocert.NewListener("shum.fyi"), h); err != nil {
				log.Fatalf("couldn't serve autocert listener: %v", err)
			}
		}()

		// Redirect all requests on port 80 to HTTPS on port 443.
		redir := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+":443"+r.RequestURI, http.StatusMovedPermanently)
		})

		if err := http.ListenAndServe(":80", redir); err != nil {
			log.Fatalf("couldn't redirect port 80: %v", err)
		}
	}
}
