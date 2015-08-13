package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"code.google.com/p/go-uuid/uuid"
	"github.com/gorilla/sessions"
	"github.com/simplq/handlers/validate"
)

var store = sessions.NewCookieStore(uuid.New())

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) { // use a swith statement instead?
	fmt.Println("the address is:", r.URL.Path)
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
	r.ParseForm() // parse arguments, you have to call this by yourself
	/*fmt.Println(r.Form) // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) */
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Herro thur") // send data to client side
}
func getLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/login.html")
	// get cookie
	//r.Cookie("username")
	if err != nil {
		//fmt.Println("Login err template ", err)
	} else {
		t.Execute(w, nil)
	}
}
func postLogin(w http.ResponseWriter, r *http.Request) { // submit login information
	r.ParseForm()
	if isValid := validateLogin(w, r); isValid {
		t, err := template.ParseFiles("web/home.html")
		if err != nil {
			//fmt.Println("(ignore if running a test) Login err template ", err)
		} else {
			// set cookie
			// create session along with session id, include session id with cookie below
			//http.SetCookie(w, &http.Cookie{Name: r.Form.Get("email"), Value: r.Form.Get("password"), Expires: time.Now().Add(30 * time.Minute)})
			//http.SetCookie(w, &http.Cookie{Name: "simplqSessionId, uuid.Uuid(), Expires: time.Now().Add(30 * time.Minute)})
			t.Execute(w, nil)
		}
	} else {
		fmt.Errorf("Invalid login")
		t, err := template.ParseFiles("web/login.html")
		if err != nil {
			fmt.Println("(ignore if running a test) Login err template ", err)
		} else {
			t.Execute(w, nil)
		}
	}

}

// ---- helper methods -------- //
func validateLogin(w http.ResponseWriter, r *http.Request) bool {
	email := r.Form.Get("email")
	isEmail := validate.IsEmailAddress(email)
	password := r.Form.Get("password")
	isPassword := validate.IsPassword(password)
	//fmt.Println(isEmail, isPassword, "-------------")
	return isEmail && isPassword
	//fmt.Fprintf(w, "Email: "+email+" "+strconv.FormatBool(isEmail)+"  Password:") // send data to client side
	//fmt.Fprintf(w, password+" "+strconv.FormatBool(isPassword))                   // send data to client side
	//w.Write([]byte("aaa"))
}

type Handle func(w http.ResponseWriter, r *http.Request)
type Mux struct { // mux is a handler. It only has functions
}

func GetMux() *Mux {
	return &Mux{}
}
func ServeAndHandle(port string) http.Handler {
	m := GetMux()
	err := http.ListenAndServe(":"+port, m) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return m
}

func pathMethod(w http.ResponseWriter, r *http.Request, path string, method string, h Handle) bool {
	status := r.URL.Path == path && r.Method == method
	if status {
		h(w, r)
	}
	return status
}
