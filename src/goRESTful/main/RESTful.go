package main

import (
	"github.com/drone/routes"
	"net/http"
	"fmt"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are getting user %s.", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are modifying user %s.", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are deleting user %s", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprint(w, "you are adding user %s", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	mux.Del("user/:uid", deleteuser)
	mux.Put("user/:uid", adduser)
	http.Handle("/", mux)
	http.ListenAndServe(":8811", nil)
}
