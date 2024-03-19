package main

import (
	"bytes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	go WorkerTest()
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
		rev.ServeHTTP(w, r)
	})
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

func WorkerTest() {
	t := time.NewTicker(5 * time.Second)
	var tn string
	var b int = 0
	var counterLine, timeLine int
	fileData, err := os.ReadFile("/app/static/tasks/_index.md")
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(fileData, []byte("\n"))
	for i, l := range lines {
		if bytes.Contains(l, []byte("Счетчик:")) {
			counterLine = i
		}
		if bytes.Contains(l, []byte("Текущее время:")) {
			timeLine = i
		}
	}

	for {
		select {
		case <-t.C:
			tn = time.Now().Format("2006-01-02 15:04:05")
			lines[timeLine] = []byte("Текущее время: " + tn)
			lines[counterLine] = []byte("Счетчик: " + strconv.Itoa(b))
			err := os.WriteFile("/app/static/tasks/_index.md", bytes.Join(lines, []byte("\n")), 644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
