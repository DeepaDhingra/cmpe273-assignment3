# cmpe273-assignment3


Part II - Trip Planner

The trip planner is a feature that will take a set of locations from the database and will then check against UBER’s price estimates API to suggest the best possible route in terms of costs and duration.

Execution Steps:

Compile and run the go files using

go run file1.go file2.go

POST        /trips   # Plan a trip
GET        /trips/{trip_id} # Check the trip details and status
PUT        /trips/{trip_id}/request 

Execution Post Operation:

Pre-requisite : Execute Post opearation as shown below to generate Location Id's:






POST Opreation JSON in :

POST        /trips   # Plan a trip

        Request:

     Example: Inv   

{

    "starting_from_location_id: "999999",

    "location_ids" : [ "10000", "10001", "20004", "30003" ] 

}

