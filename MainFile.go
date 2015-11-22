package main

import (
 	"fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "gopkg.in/mgo.v2"
		)
		
type UserController struct {  
    session *mgo.Session
}
func NewUserController(s *mgo.Session) *UserController {  
    return &UserController{s}
}

//Below function --connection with Remote mongolab
func RemoteMongoSession() *mgo.Session {
	mongolab_uri := "mongodb://mongoDeepa:welcome1@ds048878.mongolab.com:48878/mongodatabase"
	session, err := mgo.Dial(mongolab_uri)
	//Check if connection error
  	if err != nil {
    	fmt.Println("Connection problem with mongo, go error %v\n", err)
  	}	
	return session
}

func main() {
    router := httprouter.New()
// Get a UserController instance
 //uc := NewUserController(getSession())    
  uc := NewUserController(RemoteMongoSession())    
 
 
   router.POST("/trips", uc.createpost) 
   router.GET("/trips/:trip_id", uc.GetUser)
   router.PUT("/trips/:trip_id/request", uc.PutUser)
   server := http.Server{
   Addr:        "0.0.0.0:8080",
   Handler: router,
    }
    server.ListenAndServe()
}