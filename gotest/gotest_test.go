package gotest

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/simplq/handlers"
)

const addr = "http://localhost:8080/"

//const addr = "http://::1/"

func TestIndex(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpCreateRequest(t, h, "GET", "")

}
func TestGetLogin(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpCreateRequest(t, h, "GET", "login")
}
func TestPostLogin(t *testing.T) {
	ts, h := helpPrepServer()
	defer ts.Close()
	helpPostLogin(t, h, "login")
}

// ----- Helper Functions ------ //
func helpPostLogin(t *testing.T, h *handlers.Mux, path string) {
	req, err := http.PostForm(addr+path,
		url.Values{"username": {"zlaw777"}, "password": {"12345"}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(req.Body)
	// add ability to test forms
	//req.PostForm = http.PostForm("http://example.com/form",
	//url.Values{"username": {"zlaw777"}, "id": {"123"}})
	//w := httptest.NewRecorder()
	//h.ServeHTTP(w, req)
	/*if w.Code != 200 {
		t.Fatal("Code is", w.Code)
	}
	fmt.Println(w.Body.String()) */
}
func helpCreateRequest(t *testing.T, h *handlers.Mux, method string, path string) {
	req, err := http.NewRequest(method, addr+path, nil)
	if err != nil {
		fmt.Println("AEUAU")
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Fatal("Code is", w.Code)
	}
	fmt.Println(w.Body.String())
}
func helpPrepServer() (*httptest.Server, *handlers.Mux) {
	h := handlers.GetMux()
	server := httptest.NewServer(h)
	server.URL = addr
	//server.Start()
	//server.ListenAndServe()
	return server, h
}
