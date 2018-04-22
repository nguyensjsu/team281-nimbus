
/*
	Starbucks API in Go 

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


// MongoDB Configuration
var mongodb_server = "localhost:27017"
var mongodb_database = "Tea"
var mongodb_collection = "starbucksTea"



// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}


// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", homepageHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory", inventoryTea(formatter)).Methods("GET")
	mx.HandleFunc("/cartItemsTeas", cartHandlerTea(formatter)).Methods("GET")
	mx.HandleFunc("/searchInventoryTeas", searchInventoryTea(formatter)).Methods("GET")
	//Below - PUT : Status 0 - 1
	mx.HandleFunc("/updateTea", UpdateTeas(formatter)).Methods("PUT")
	

}


// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}



// Ping Application For Tea Section
func homepageHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks Tea Section!"})
	}
}

//API returns  inventory with status 0 
func inventoryTea(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		//establishing session with DB
		fmt.Println("Inventory details:")
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

		var result []*bson.M

        err = c.Find(bson.M{"items":"espresso"}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)

	}
}


//API returns  inventory with status 1
func cartHandlerTea(formatter *render.Render) http.HandlerFunc{

	return func(w http.ResponseWriter, req *http.Request) {

		//establishing session with DB
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

		var result []*bson.M

        err = c.Find(bson.M{"status":1}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)

	}
}


// API seacrhes inventory for a specific item and returns it
func searchInventoryTea(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M

        err = c.Find(bson.M{"item" : "latte"}).One(&result)

        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Inventory details:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}


//API adds item to cart - Updates status from 0 to 1
func UpdateTeas(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		
			var lat []bson.M
			json.NewDecoder(req.Body).Decode(&simulation)		
	    	
			session, err := mgo.Dial(mongodb_server)
	        if err != nil {
	                panic(err)
	        }
	        defer session.Close()
	        session.SetMode(mgo.Monotonic, true)
	        c := session.DB(mongodb_database).C(mongodb_collection)
	  		
	        var lat []bson.M
	        err = c.Find(bson.M{"status":0}).All(&lat)
	        if err != nil {
			log.Fatal(err)
			
	        } 

		formatter.JSON(w, http.StatusOK, lat)
	}