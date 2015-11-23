package main

import (
 	"fmt"
    "github.com/julienschmidt/httprouter"
	"github.com/anweiss/uber-api-golang/uber"
    "net/http"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
     "io/ioutil"
     "strconv"
     "log"
	)
		
type UberApi struct {
		Prices []struct {
		Estimate string `json:"estimate"`
		LowEstimate int `json:"low_estimate"`
		HighEstimate int `json:"high_estimate"`
		Duration int `json:"duration"`
		Distance float64 `json:"distance"`
	} `json:"prices"`
}

type UberPOSTRequest struct {

	Starting_from_location_id string `json:"starting_from_location_id"`
	Location_ids []string `json:"location_ids"`
}

type UberApiResponse struct {
        
	 Id     bson.ObjectId `json:"id" bson:"_id"`
     Status string `json:"status"` 
     Starting_from_location_id string  `json:"starting_from_location_id"`
     Best_route_location_ids []string `json:"best_route_location_ids"`
     Total_uber_costs int `json:"total_uber_costs"`
     Total_uber_duration int `json:"total_uber_duration"`
     Total_distance float64 `json:"total_distance"` 
}
	
type (  
    UserLocation struct {
        Id     bson.ObjectId `json:"id" bson:"_id"`
        Name   string        `json:"name" bson:"name"`
        Address string       `json:"address" bson:"address"`
        City    string       `json:"city" bson:"city"`
        State    string      `json:"state" bson:"state"`
        Zip string		 `json:"zip" bson:"zip"`
        Coordinate struct {   
        	Latitude float64  `json:"lat" bson:"lat"`
        	Longitude float64 `json:"lng" bson:"lng"`
        }

    }    
)	

type ( 
CostRoot struct {
Cost int
RootLoc[] string 
}
)


func (uc UserController) createpost(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {


body, err := ioutil.ReadAll(req.Body)

 u := UberPOSTRequest{
        }
        
        
 users := UberApiResponse{
 }
 
     err = json.Unmarshal(body, &u)
      if err != nil {
        fmt.Println("error is",err)
    }

  
 //PERM LOGIC
 
  var in []string
  in = u.Location_ids
  l:=permutations(in)
  fmt.Println("\n",l)
  fmt.Println("len is",len(l))
   StartResult := UserLocation{
        }

EndResult := UserLocation{
        }
        var count int
  		count = len(l)
        FinalCostArray := make([] int, count)
        FinalDurationArray := make([] int, count)
        FinalDistanceArray := make([] float64, count)
        
    var    minimum int
    var value int
    var BestRoute int
  for k:=0;k<len(l);k++ {
  var CostUber int
  CostUber = 0;
  var DurationUber int
  DurationUber = 0;
  var DistanceUber float64
  DistanceUber = 0;
  StartingLocid := u.Starting_from_location_id
  
  for m:=0;m<len(in)+1;m++ {
        
    oid := bson.ObjectIdHex(StartingLocid)
    uc.session.DB("mongodatabase").C("CMPE273").FindId(oid).One(&StartResult)      
    
    
    var EndLocationId string
    
    if m<len(in) {
    
    EndLocationId= l[k][m]
    }
    
    if m==len(in){
    
    EndLocationId = StartingLocid 	
    }
  
 
  endoid := bson.ObjectIdHex(EndLocationId)
  uc.session.DB("mongodatabase").C("CMPE273").FindId(endoid).One(&EndResult)      //  if err != nil {
  fmt.Println("long is", EndResult.Coordinate.Longitude)
  fmt.Println("lat is", EndResult.Coordinate.Latitude)
 

  var StartLat float64
  var StartLong float64
  var EndLat float64
  var EndLong float64
  
  StartLat = StartResult.Coordinate.Latitude
  StartLong = StartResult.Coordinate.Longitude
  EndLat = EndResult.Coordinate.Latitude
  EndLong = EndResult.Coordinate.Longitude
 
 	fmt.Println("welcome to uber");
 

	options1 := uber.RequestOptions{
	ServerToken: "GyJ-Y_iqijuAFgEd5rQBN0XV9Ojy6ybKWgvET6jd",
    ClientId:  "VS8-KN-nVJ4jWyDEYdBQnZU-muBpHicd",
    ClientSecret: "aK9O5H02Uuq6n34x74SJaM6QgMy7UAkIJqzjMeRd",
    AppName: "mongodatabaseDeepa",
}

 
	client := uber.Create(&options1)
 
	pe := &uber.PriceEstimates{}
	pe.StartLatitude = StartLat
	pe.StartLongitude = StartLong
	pe.EndLatitude = EndLat
	pe.EndLongitude = EndLong
	if e := client.Get(pe); e != nil {
		log.Fatal(e)
	}

var PriceUber int
PriceUber=0
var DisUber float64
DisUber=0
var DurUber int
DurUber=0

	fmt.Println("\nHere are the Uber price estimates: \n")
	for _, price := range pe.Prices {
	if(price.DisplayName=="uberX"){
	PriceUber=price.LowEstimate
	
	DurUber=price.Duration
	DisUber=price.Distance
		fmt.Println(price.DisplayName + ": " + strconv.Itoa(price.LowEstimate))
	
	}
	
	}
	
			CostUber=CostUber+PriceUber
			DurationUber=DurationUber+DurUber
			DistanceUber=DistanceUber+DisUber
			
		if(m<len(in)){
		
			StartingLocid=l[k][m]
			 fmt.Println("swapped start loc id is ****************",StartingLocid)
  }
		fmt.Println("Cost uber is",CostUber )
			
		
 } //inner loop ends here
 
 
 FinalCostArray[k] = CostUber
 FinalDurationArray[k]=DurationUber
        FinalDistanceArray[k]=DistanceUber//float 64 to int
      
   		       }
      
       fmt.Println("final cost array contains",FinalCostArray)
       fmt.Println("final duration array contains",FinalDurationArray)
       fmt.Println("final distance array contains",FinalDistanceArray)
               
         minimum=FinalCostArray[0]
         fmt.Println("minimum value us",minimum)
         fmt.Println("len is", len(FinalCostArray))
         for i:=1;i<len(FinalCostArray);i++ {
         fmt.Println("i is",i)
         if (value > FinalCostArray[i]){
         
         
               value = FinalCostArray[i]
               fmt.Println("value is ", value)
               minimum = value
               BestRoute = i
                  
              fmt.Println("minimum value  inside if loop is",minimum)
                  }
         
         }
       			fmt.Println("total uber cost is",minimum)
       			fmt.Println("minimum is on location", BestRoute)
       
              	fmt.Println("best location",l[BestRoute])
               fmt.Println("best location distance",FinalDistanceArray[BestRoute])
               fmt.Println("best location duration",FinalDurationArray[BestRoute])
              
       users.Id = bson.NewObjectId()
		users.Status="planning"
		users. Starting_from_location_id=u.Starting_from_location_id
		users.Best_route_location_ids=l[BestRoute]
		users.Total_uber_costs=FinalCostArray[BestRoute]
		users.Total_uber_duration=FinalDurationArray[BestRoute]
		users.Total_distance=FinalDistanceArray[BestRoute]

	        //Write the user to mongolab
    uc.session.DB("mongodatabase").C("CMPE273Assignment3").Insert(users)
     
          
    // Marshal provided interface into JSON structure
    uj, _ := json.Marshal(users)
        
   	rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(201)
    fmt.Fprintf(rw, "%s", uj)
       
       
       
      
      }

