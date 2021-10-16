package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dinel13/lanjukang/config"
	"github.com/dinel13/lanjukang/db/driver"
	"github.com/dinel13/lanjukang/handlers"
	"github.com/dinel13/lanjukang/pkg/utilities"
)

const portNumber = ":8080"

// app is a global variable for this package
var app config.AppConfig

func main() {
	app.AppName = "Lanjukang"
	app.AppVersion = "1.0.0"
	app.JwtSecret = "secret"

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	// send mail
	to := []string{"andirunawa13@gmail.com", "baharuddinbutdam@gmail.com"}
	subject := "Test email"
	body := "This is a test email"

	// make chanel as receiver for sending email
	mailEror := make(chan error)
	go func() {
		mailEror <- utilities.SendMail(to, subject, body)
	}()
	errMail := <-mailEror
	if errMail != nil {
		log.Fatal(errMail)
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=lanjukang user=din password=postgres")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	log.Println("Connected to database!")

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)

	return db, nil
}
