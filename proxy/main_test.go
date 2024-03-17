package main

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestAPIHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantBody   string
	}{
		{
			name:       "200 OK",
			args:       args{r: &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/api/"}}}, //эт ваще не имеет значения))
			wantStatus: 200,
			wantBody:   "Hello from API",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.args.r.Method, tt.args.r.URL.String(), nil)
			rr := httptest.NewRecorder()
			APIHandler(rr, req)
			assert.Equal(t, tt.wantStatus, rr.Code)
			assert.Equal(t, tt.wantBody, rr.Body.String())
		})
	}
}

func TestNewReverseProxy(t *testing.T) {
	type args struct {
		host string
		port string
	}
	tests := []struct {
		name string
		args args
		want *ReverseProxy
	}{
		{
			name: "any",
			args: args{host: "any", port: "any"},
			want: &ReverseProxy{host: "any", port: "any"},
		},
		{
			name: "any",
			args: args{host: "asdff", port: "1234"},
			want: &ReverseProxy{host: "asdff", port: "1234"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReverseProxy(tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReverseProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseProxy_ReverseProxy(t *testing.T) {
	type fields struct {
		host string
		port string
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantStatus int
	}{
		{
			name:   "api",
			fields: fields{host: "localhost", port: "1313"},
			args: args{
				method: "GET",
				path:   "/api/",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:   "bad",
			fields: fields{host: "localhost", port: "1313"},
			args: args{
				method: "GET",
				path:   "/some/path",
			},
			wantStatus: http.StatusNotFound,
		},
		{
			name:   "good",
			fields: fields{host: "localhost", port: "1313"},
			args: args{
				method: "GET",
				path:   "/",
			},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rp := NewReverseProxy(tt.fields.host, tt.fields.port)
			r := chi.NewRouter()
			r.Use(rp.ReverseProxy)
			r.Get("/api/*", APIHandler)
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(tt.args.method, tt.args.path, nil)

			r.ServeHTTP(rr, req)
			assert.Equal(t, tt.wantStatus, rr.Code)
		})
	}
}
