package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gopkg.in/mgo.v2"

	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/controllers"
	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/view"
)

func main() {
	// for better login on crash
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hire Server initializing")

	// initialize html templates
	view.Initialize()

	controllers.HC = controllers.NewHireController(getMongoConnection())

	http.HandleFunc("/", index)
	http.HandleFunc("/users", controllers.Index)
	http.HandleFunc("/users/show", controllers.Show)
	http.HandleFunc("/users/create", controllers.Create)
	http.HandleFunc("/users/create/process", controllers.CreateProcess)
	http.HandleFunc("/users/update", controllers.Update)
	http.HandleFunc("/users/update/process", controllers.UpdateProcess)
	http.HandleFunc("/users/delete/process", controllers.DeleteProcess)

	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()

	// wait for interrupt signal to allow for graceful shutdown
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)

	<-ch
	shutdownServer(controllers.HC)
}

// getMongoConnection is used to establish the connection to the database
func getMongoConnection() (*mgo.Session, *mgo.Collection, context.Context, context.CancelFunc) {
	log.Println("Starting database")

	client, err := mgo.Dial("mongodb://127.0.0.1:27017")
	//client, err := mgo.Dial("mongodb://mongo:27017")
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 4000*time.Second)

	// opens MongDB database and its collection (creates them if they don't exitst)
	collection := client.DB("mydb").C("hires")

	return client, collection, ctx, cancel
}

// shutdownServer is used for graceful shutdown
func shutdownServer(c *controllers.HireController) {
	log.Println("Stopping database connection")
	c.CancelFunc()
}

// index servers as a web redirector for /users
func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
