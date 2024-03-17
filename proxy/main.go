package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")

	r.Use(proxy.ReverseProxy)
	r.Use(middleware.Recoverer)

	r.Get("/api/*", APIHandler)

	http.ListenAndServe(":8080", r)
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.String(), "/api/") {
			next.ServeHTTP(w, r)
			return
		}
		rev := httputil.ReverseProxy{Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(&url.URL{Scheme: "http", Host: rp.host + ":" + rp.port})
			r.SetXForwarded()
		}}
		//httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: rp.host + ":" + rp.port}).ServeHTTP(w, r)
		rev.ServeHTTP(w, r)

	})
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

//
//const content = ``
//
//func WorkerTest() {
//	t := time.NewTicker(1 * time.Second)
//	var b byte = 0
//	for {
//		select {
//		case <-t.C:
//			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
//			if err != nil {
//				log.Println(err)
//			}
//			b++
//		}
//	}
//}
