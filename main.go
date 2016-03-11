package main

import (
	"net/http"
	"encoding/json"
	"time"
	)

type Payload struct {
	ServerInfo ServerInfo
}

type ServerInfo struct {
	Timestamp time.Time
	IpAddress string
	Name string
}

func main(){
	http.HandleFunc("/", serverTest)
	http.ListenAndServe(":8080", nil)
}

func serverTest(w http.ResponseWriter, req *http.Request) {
	response, err := getJsonResponse(req)
	if err != nil {
		panic(err)
	}

	w.Write(response)
}

func getJsonResponse(req *http.Request) ([]byte, error) {
	var tmstp time.Time = time.Now()
	var ipAddress, name string = req.Host, "rate is coming"

	s := ServerInfo{tmstp, ipAddress, name}
	p := Payload{s}

	return json.MarshalIndent(p, "", "  ")
}
