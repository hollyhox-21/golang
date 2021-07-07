package main

import (
	"errors"
	"fmt"
	"github.com/hollyhox-21/notpad/pkg/models"
	"net/http"
	"strconv"
	"text/template"
)

func (a *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	s, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}

	//fmt.Printf("%v\n", s)

	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		a.serverError(w, err)
		return
	}
	err = tmpl.Execute(w, s)
	if err != nil {
		a.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (a *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}
	//w.Write([]byte("Show snippet"))
	//fmt.Fprintf(w, "View snippet with ID = %d...", id)
	s, err := a.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			a.notFound(w)
		} else {
			a.serverError(w, err)
		}
		return
	}
	fmt.Fprintf(w, "%v",s)
}

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("forbidden method")) эквивалент http.Error()
		//http.Error(w, "forbidden method", 405)
		a.clientError(w, http.StatusMethodNotAllowed) // обертка над обработкой ошибки
		return
	}
	title := r.PostFormValue("title")
	content := r.PostFormValue("note")

	_, err := a.snippets.Insert(title, content, "7")
	if err != nil {
		a.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *application) deleteSnippet(w http.ResponseWriter, r *http.Request)  {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		a.clientError(w, http.StatusMethodNotAllowed) // обертка над обработкой ошибки
		return
	}
	ids := r.PostFormValue("id")

	id, err := strconv.Atoi(ids)
	if err != nil {
		a.serverError(w, err)
		return
	}

	err = a.snippets.Delete(id)
	if err != nil {
		a.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
