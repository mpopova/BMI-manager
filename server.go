package main

import (
	"fmt"
	"html/template"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

type Profile struct {
		username   string
		gender string
		age string
	}

func getUserInfo(request *http.Request) (userName string, gender string, age string) {
	if cookie, err := request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
			gender = cookieValue["gender"]
			age = cookieValue["age"]
		}
	}
	return userName, gender, age
}

func setSession(userName string, gender string, age string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
		"gender" : gender,
		"age" : age,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func renderHtml(w http.ResponseWriter, tmpl string) {
    tmpl = fmt.Sprintf("templates/%s", tmpl)
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, "")
    if err != nil {
        log.Print("template executing error: ", err)
    }
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

// login handler
func loginHandler(response http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/bmi")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		username string
		password string
		gender string
		age string
	)

	name := request.FormValue("name")
	pass := request.FormValue("password")

	rows, err := db.Query("select username, password, gender, age from users where username = ?", name)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&username, &password, &gender, &age)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	redirectTarget := "/"
	if name == username && pass == password {
		// .. check credentials ..
		setSession(name, gender, age, response)
		redirectTarget = "/calculate"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

// logout handler
func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

// register handler
func registerHandler(response http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/bmi")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := request.FormValue("name")
	pass := request.FormValue("password")
	gender := request.FormValue("gender")
	age := request.FormValue("age")

	stmt, err := db.Prepare("INSERT INTO users(username, password, gender, age) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(user, pass, gender, age)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	redirectTarget := "/"
	if rowCnt == 0 {
		redirectTarget = "/registration"
	}

	http.Redirect(response, request, redirectTarget, 302)
}

func profileHandler(response http.ResponseWriter, request *http.Request) {
	username, gender, age := getUserInfo(request)

	fmt.Println("response Body:")
}

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
	renderHtml(response, "index.html")
}

func calculatePageHandler(response http.ResponseWriter, request *http.Request) {
	userName, gender, age := getUserInfo(request)
	if userName != "" {
		renderHtml(response, "calculate.html")
	} else { 
		http.Redirect(response, request, "/", 302)
	}
}

func registerPageHandler(response http.ResponseWriter, request *http.Request){
	renderHtml(response, "register.html")
}

func profilePageHandler(response http.ResponseWriter, request *http.Request) {
	userName, gender, age := getUserInfo(request)
	if userName != "" {
		renderHtml(response, "profile.html")
	} else { 
		http.Redirect(response, request, "/", 302)
	}
}

var router = mux.NewRouter()

func main() {
	router.HandleFunc("/", indexPageHandler)
	router.HandleFunc("/calculate", calculatePageHandler)
	router.HandleFunc("/registration", registerPageHandler)
	router.HandleFunc("/profile", profilePageHandler)

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler)
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/userinfo", profileHandler).Methods("GET")
	http.Handle("/", router)
	

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js")))) 

	http.ListenAndServe(":8080", nil)
}