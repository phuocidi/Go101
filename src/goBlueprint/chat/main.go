package main

import (
	"flag"
	// "github.com/gorilla/pat"
	// "github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/gplus"
	_ "goBlueprint/trace"
	"log"
	"net/http"
	_ "os"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse() // parse the flags

	// p := pat.New()

	goth.UseProviders(
		gplus.New("208960536631-ac8p6vrd49t9i9lhgmasm4c8eqp4vnqe.apps.googleusercontent.com",
			"qVX8UOu6a7qZTD0ml7WH1lJ-",
			"http://localhost:8080/auth/callback/google",
		),
	)

	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))

	fs := http.FileServer(http.Dir("js"))
	http.Handle("/js/", http.StripPrefix("/js/", fs))

	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	http.Handle("/room", r)
	// run the chat room in the background
	go r.run()

	log.Println("Starting web server on: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
