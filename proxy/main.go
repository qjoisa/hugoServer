package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
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

const content = ``

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	var tn string
	var b byte = 0
	fileData, err := os.ReadFile("/app/static/tasks/_index.md")
	fmt.Println(string(fileData))
	if err != nil {
		log.Fatal(err)
	}
	nTime := bytes.LastIndex(fileData, []byte("Текущее время:"))
	nCounter := bytes.LastIndex(fileData, []byte("Счетчик:"))
	for {
		tn = time.Now().Format("2006-01-02 15:04:05")
		select {
		case <-t.C:
			res := append(fileData[:nTime], append([]byte(tn), fileData[nTime+len(tn):]...)...)
			res = append(fileData[:nCounter], append([]byte(tn), fileData[nCounter+len(tn):]...)...)
			err := os.WriteFile("/app/static/_index.md", res, 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
