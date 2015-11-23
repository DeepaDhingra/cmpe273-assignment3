# cmpe273-assignment3


Trip Planner

The trip planner is a feature that will take a set of locations from the database and will then check against UBER’s price estimates API to suggest the best possible route in terms of costs and duration.

Execution Steps:

Import following packages and the run the go files using go run command: 
"github.com/julienschmidt/httprouter"
"github.com/anweiss/uber-api-golang/uber"



GET        /trips/{trip_id} # Check the trip details and status
PUT        /trips/{trip_id}/request 

1) Execution Post Operation :

a) As explain in below github Read me, execute POST operation for locations to get the locations id's

https://github.com/DeepaDhingra/cmpe273-assignment2/blob/master/README.md

Pre-requisite : Execute Post opearation as shown below to generate trip Id's:

POST Opreation JSON in :

POST        /trips   # Plan a trip

Request
{

    "starting_from_location_id: "999999",

    "location_ids" : [ "10000", "10001", "20004", "30003" ] 

}

 Response: HTTP 201


{

     "id" : "1122",

     “status” : “planning”,

     "starting_from_location_id: "999999",

     "best_route_location_ids" : [ "30003", "10001", "10000", "20004" ],

  "total_uber_costs" : 125,

  "total_uber_duration" : 640,

  "total_distance" : 25.05 

}

2 GET        /trips/{trip_id} # Check the trip details and status
        

        Request:  GET             /trips/1122


Response:


{

     "id" : "1122",

     "status" : "planning",

     "starting_from_location_id: "999999",

     "best_route_location_ids" : [ "30003", "10001", "10000", "20004" ],

  "total_uber_costs" : 125,

  "total_uber_duration" : 640,

  "total_distance" : 25.05 

}



  PUT        /trips/{trip_id}/request # Start the trip by requesting UBER for the first destination. You will call UBER request API to request a car from starting point to the next destination.
        

        UBER Request API:  PUT /v1/sandbox/requests/{request_id}

        

        Once a destination is reached, the subsequent call the API will request a car for the next destination.


        Request:  PUT             /trips/1122/request


Response:


{

     "id" : "1122",

     "status" : "requesting",

     "starting_from_location_id”: "999999",

     "next_destination_location_id”: "30003",

     "best_route_location_ids" : [ "30003", "10001", "10000", "20004" ],

  "total_uber_costs" : 125,

  "total_uber_duration" : 640,

  "total_distance" : 25.05,

  "uber_wait_time_eta" : 5 

}
