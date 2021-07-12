package main

import (
	"database/sql"
	"flag"
	"github.com/hollyhox-21/notpad/pkg/models/mysql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)


type application struct {
	infoLog *log.Logger
	errorLog *log.Logger
	snippets *mysql.SnippetModel
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	//create flags to start
	addr := flag.String("addr", ":4000", "Сетевой адрес HTTP")
	flag.Parse()

	dsn := flag.String("dsn", "web:password@tcp(docker.for.mac.localhost:3306)/snippetbox?parseTime=true", "Название MySQL источника данных")

	//creat std logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		infoLog: infoLog,
		errorLog: errorLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
	}

	app.startServer(srv)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err = nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
