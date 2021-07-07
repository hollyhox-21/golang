package main

import "net/http"

func (a *application) routes() *http.ServeMux {
	//create handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.home)
	mux.HandleFunc("/snippet", a.showSnippet)
	mux.HandleFunc("/snippet/create", a.createSnippet)
	mux.HandleFunc("/snippet/delete", a.deleteSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}

func (a *application) startServer(srv *http.Server) {
	a.infoLog.Printf("Starting web-server on %s", srv.Addr)
	err := srv.ListenAndServe()
	a.errorLog.Fatal(err)
}