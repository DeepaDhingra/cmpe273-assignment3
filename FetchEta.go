
/* The Aceess Token is valid for 30 days. It is issued on 11/20/2015*/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
		)

type UberResponse struct 
{
	Status string
	Eta int
}


func CallUberAPI(start_latitude, start_longitude, end_latitude, end_longitude string) int {
	url := "https://sandbox-api.uber.com/v1/requests"
	fmt.Println("URL:", url)
	
	jsonString := "{\"start_latitude\":"+ start_latitude +",\"start_longitude\":" + start_longitude +",\"end_latitude\":" + end_latitude +",\"end_longitude\":" + end_longitude +",\"product_id\":" + "\"2832a1f5-cfc0-48bb-ab76-7ea7a62060e7\",\"scope\":\"request\"}"

	var jsonStr = []byte(jsonString)
	
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzY29wZXMiOlsicmVxdWVzdCJdLCJzdWIiOiJiNTJkZDhjMi1lYzkzLTRiOWUtYThhYy02NDI0MmQ3NzQwNmMiLCJpc3MiOiJ1YmVyLXVzMSIsImp0aSI6IjkxY2Y0MmQ1LTE3NTktNDNlZC04YTY5LWQwZjEzMmYwMjc2MyIsImV4cCI6MTQ1MDcyNTM3MywiaWF0IjoxNDQ4MTMzMzczLCJ1YWN0IjoiOUxxS2dER3h0MW56R1AwdXNsRHU2U1RzT09SQjNkIiwibmJmIjoxNDQ4MTMzMjgzLCJhdWQiOiJWUzgtS04tblZKNGpXeURFWWRCUW5aVS1tdUJwSGljZCJ9.EAs8qgL8x_kZYNLU-kZGk02GVxiMCJ3eKgsLNlsdeTaUB_aNfEY2_ab-zW0C9CP0YnUxlml8hGDgcUTiM3P982gDyo1WBWnnpaqmqqDZbbkrssxf2zGPpfuRCRB9tdpJw_DxGq4TYh7bmanrEedMOQwVVgQhjTBtZnR0tsFnQL4_rWetcH-ImbUK-6mYGs602DEDv-gz3zZzT1636K0ia5OL0m4O61GOteTJ8AMwtcqHmUVaizU8OsxOpO1r5aLqv3f2NXp4Hybxu65ATz1dwS1G8XVmPHAidS-EdaMn-EopHe8hDmDLa-nCV9LnWu2Z3jFpHNzz5ozmyA0f7dLsMw")
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	
		var UberResponseobj UberResponse
	
	err = json.Unmarshal(body, &UberResponseobj)
	if err != nil {
	
	fmt.Println("Err:", err)
	}
	
	fmt.Println("ETA*********:", UberResponseobj.Eta)
	return UberResponseobj.Eta
}