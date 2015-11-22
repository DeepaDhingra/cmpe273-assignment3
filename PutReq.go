package main
import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "strconv"
    	
        )
        
        type UberApiPutResponse struct {
        
	 Id     bson.ObjectId `json:"id" bson:"_id"`
     Status string `json:"status"` 
     Starting_from_location_id string  `json:"starting_from_location_id"`
     Next_destination_location_id string  `json:"next_destination_location_id"`
     Best_route_location_ids []string `json:"best_route_location_ids"`
     Total_uber_costs int `json:"total_uber_costs"`
     Total_uber_duration int `json:"total_uber_duration"`
     Total_distance float64 `json:"total_distance"` 
     Uber_wait_time_eta int `json:"uber_wait_time_eta"`
}

var BestRoutePutArray[] string
var localRoutePutArray[] string


var called int


var uberRspMap UberApiPutResponse
// To hold array of multiple stocks + MAP
type Holder struct {
	Hold UberApiPutResponse
}

var AccessHolder Holder

var M = make(map[int]Holder)
 
 
func (uc UserController) PutUser(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {  

users1 := UberApiResponse{
 }
    
    fmt.Println("Foo! i am caled");
    id := p.ByName("trip_id")
    fmt.Println("id is",id)

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        rw.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)
    
    usersput := UberApiResponse{	}
 
 	userPutRes := UberApiPutResponse{}
  	
 	   StartResultCord := UserLocation{
        }

EndResultCord := UserLocation{
        }
 	
    // Fetch user
    if (called==0){
    fmt.Println("i am called")
    if err := uc.session.DB("mongodatabase").C("CMPE273Assignment3").FindId(oid).One(&usersput); err != nil {
        rw.WriteHeader(404)
        return
        }
       fmt.Println("user put is",usersput.Status)
  
   //Make a call to Uber Sandbox's POST API for getting ETA
   
 
     
        userPutRes.Id = oid
    userPutRes.Status = "Requesting" 
    userPutRes.Starting_from_location_id = usersput.Starting_from_location_id 
    userPutRes.Next_destination_location_id = usersput.Best_route_location_ids[0] //calculate from arrayindex
    userPutRes.Best_route_location_ids = usersput.Best_route_location_ids
    userPutRes.Total_uber_costs = usersput.Total_uber_costs
    userPutRes.Total_uber_duration = usersput.Total_uber_duration
    userPutRes.Total_distance = usersput.Total_distance
    startID:=userPutRes.Starting_from_location_id
    nextID:=  userPutRes.Best_route_location_ids[0]
   
     
   startoid := bson.ObjectIdHex(startID)
   nextoid := bson.ObjectIdHex(nextID)
      uc.session.DB("mongodatabase").C("CMPE273").FindId(startoid).One(&StartResultCord)  
      uc.session.DB("mongodatabase").C("CMPE273").FindId(nextoid).One(&EndResultCord)  
   
  slat:=strconv.FormatFloat(StartResultCord.Coordinate.Latitude, 'f', 2, 32)
  slong:=strconv.FormatFloat(StartResultCord.Coordinate.Longitude, 'f', 2, 32)
  elat:=strconv.FormatFloat(EndResultCord.Coordinate.Latitude, 'f', 2, 32)
  elong:=strconv.FormatFloat(EndResultCord.Coordinate.Longitude, 'f', 2, 32)
  
         eta := CallUberAPI(slat,slong,elat,elong)
    	userPutRes.Uber_wait_time_eta = eta
     	fmt.Println("ETA in calling in PUT", eta)
    
		uberRspMap = userPutRes
 		AccessHolder.Hold = uberRspMap;
 
 		M[called] = Holder{AccessHolder.Hold}
        uc.session.DB("mongodatabase").C("CMPE273Assignment3").FindId(oid).One(&users1)
		users1.Status = "Requesting"
        uc.session.DB("mongodatabase").C("CMPE273Assignment3").UpdateId(users1.Id, bson.M{"$set": users1})
    }
    

   
      if (called>0 && called < len(M[0].Hold.Best_route_location_ids)){    
      
      
     	startID:=M[0].Hold.Best_route_location_ids[called-1]
   		nextID:=  M[0].Hold.Best_route_location_ids[called]
   
   		startoid := bson.ObjectIdHex(startID)
   		nextoid := bson.ObjectIdHex(nextID)
        uc.session.DB("mongodatabase").C("CMPE273").FindId(startoid).One(&StartResultCord)  
        uc.session.DB("mongodatabase").C("CMPE273").FindId(nextoid).One(&EndResultCord)  
    
    	slat:=strconv.FormatFloat(StartResultCord.Coordinate.Latitude, 'f', 2, 32)
  		slong:=strconv.FormatFloat(StartResultCord.Coordinate.Longitude, 'f', 2, 32)
  		elat:=strconv.FormatFloat(EndResultCord.Coordinate.Latitude, 'f', 2, 32)
  		elong:=strconv.FormatFloat(EndResultCord.Coordinate.Longitude, 'f', 2, 32)
      
      
         eta := CallUberAPI(slat,slong,elat,elong)
           AccessHolder.Hold.Uber_wait_time_eta = eta
    
       AccessHolder.Hold.Next_destination_location_id = M[0].Hold.Best_route_location_ids[called] 
   		  AccessHolder.Hold.Status = "Requesting"
   		
     
   		 M[0] = Holder{AccessHolder.Hold}   
   
    	uc.session.DB("mongodatabase").C("CMPE273Assignment3").FindId(oid).One(&users1)
		users1.Status = "Requesting"
	
		 fmt.Printf("Update Users are ",users1)  
		
		uc.session.DB("mongodatabase").C("CMPE273Assignment3").UpdateId(users1.Id, bson.M{"$set": users1})
	
        } 
        
   			if (called == len(M[0].Hold.Best_route_location_ids)) {
    
    		AccessHolder.Hold.Next_destination_location_id = M[0].Hold.Starting_from_location_id
   			AccessHolder.Hold.Status = "Finished"
   			AccessHolder.Hold.Uber_wait_time_eta=0;
     		M[0] = Holder{AccessHolder.Hold}
     
     
           uc.session.DB("mongodatabase").C("CMPE273Assignment3").FindId(oid).One(&users1)
		   users1.Status = "Finished"
           uc.session.DB("mongodatabase").C("CMPE273Assignment3").UpdateId(users1.Id, bson.M{"$set": users1})
     
     
    }
         if (called > len(M[0].Hold.Best_route_location_ids)) {
          
		called = 0
		rw.WriteHeader(404)
        return
    }
        
uj, _ := json.Marshal(M[0].Hold)

    // Write content-type, statuscode, payload
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(200)
    fmt.Fprintf(rw, "%s", uj)
    called++
}
    

