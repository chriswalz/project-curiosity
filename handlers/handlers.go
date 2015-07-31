package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) { // use a swith statement instead?
	if false {
	} else if pathMethod(w, r, "/", "GET", getIndex) {
	} else if pathMethod(w, r, "/login", "GET", getLogin) {
	} else if pathMethod(w, r, "/login", "POST", postLogin) {
	} else {
		http.NotFound(w, r)
		return
	}

}

func getIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // parse arguments, you have to call this by yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Herro thur") // send data to client side
}
func getLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/login.html")
	if err != nil {
		fmt.Println("Login err template ", err)
	} else {
		t.Execute(w, nil)
	}
}
func postLogin(w http.ResponseWriter, r *http.Request) { // submit login information
	fmt.Println("parsing code")
	r.ParseForm()
	fmt.Println(r.Form["username"][0])
	fmt.Println(r.Form["password"][0])
}

// ---- helper methods -------- //
type handle func(w http.ResponseWriter, r *http.Request)
type mux struct {
}

func ServeAndHandle(port string) http.Handler {
	m := &mux{}
	err := http.ListenAndServe(":"+port, m) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return m

}

func pathMethod(w http.ResponseWriter, r *http.Request, path string, method string, h handle) bool {
	status := r.URL.Path == path && r.Method == method
	if status {
		h(w, r)
	}
	return status
}
