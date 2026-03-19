package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/router"
	"server/services"

	"github.com/joho/godotenv"
)

// @Types declaration
//cuz Go is heavily --> each asset needs to be typed for use
type Config struct {
	Port string
}

type Application struct {
	Config Config //* holds config type
	Models services.Models
}


// ! Methods
func (app *Application) Serve() error {
	//& Belongs to Application type --> the one who would iniate its instance --> method belongs to it
	err := godotenv.Load() // loading our env vars

	// if failed to load env file
	if err != nil {
		log.Fatal("failed to load env : %w",err)
	}

	//# env loaded --> use "os" package to use them✅✅
	port := app.Config.Port
	fmt.Printf("Go API is listening on port : %s🚀🚀...",port)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s",port),
		// fixed -- add chi router to route api calls to chiRouter to excecute corresponding funcs
		Handler: router.Routes(),
	}
	return server.ListenAndServe()
}

func main() {
	err := godotenv.Load() // loading our env vars

	// if failed to load env file
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	// instance of config type that stores port num
	config := Config{
		Port: os.Getenv("PORT"),
	}

	// fixed : connection to the database 
	dbConnString := os.Getenv("DSN") //get connection string
	db,err := db.Connect2Postgress(dbConnString) // struct instance that stores db connection
	if err != nil {
		log.Fatal("failed to open db connection🛑🛑...")
	}
	defer db.DB.Close() //$ as func return instance which have DB connection <- used that which hold db connection

	// ! as this is initiating instance of Application so every method that uses this type --> belongs to this
	app := &Application{
		Config: config, //* as Application holds config data <- we feeded that onto this
		// fixed : add models type instance by calling func that needs db connection
		Models: services.NewModel(db.DB), //# executes the instance and returns Model
		
	}

	err = app.Serve()
	if err!= nil {
		log.Fatal("failed to start server")
	}
}