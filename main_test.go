package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//TestHello TestHello
func TestHello(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	reader := strings.NewReader(`{"title":"The Go Standard Library","content":"It contains many packages."}`)
	r, _ := http.NewRequest(http.MethodPost, "/hello", reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
	}
}

func TestAllwaysisok(t *testing.T) {
	if allwaysisok() {
		t.Log("OK")
	} else {
		t.Error("Bad")
	}
}
