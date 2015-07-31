package gotest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/simplq/handlers"
)

func TestDivision(t *testing.T) {
	if ans, err := Division(6, 3); ans != 2 || err != nil {
		t.Error("Test division failed")
	} else {
		t.Log("Test division passed")
	}
}
func TestHome(t *testing.T) {
	//mockDb := MockDb{}
	//homeHandle := homeHandler(mockDb)
	handle := handlers.ServeAndHandle("8080")
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handle.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}
