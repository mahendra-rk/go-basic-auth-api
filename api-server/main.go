package main

import (
	"encoding/json"
	"fmt"
"net/http"
)

var (
	username = "abc"
	password = "123"
)

type Data struct {
	Meta Meta
	Resource Resource
}
type Resource map[string]string
type Meta map[string]string
type Payload struct {
	Stuff Data
}
func main() {
	handler := http.HandlerFunc(handleRequest)
	http.Handle("/example", handler)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	u, p, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		w.WriteHeader(401)
		return
	}
	if u != username {
		fmt.Printf("Username provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	if p != password {
		fmt.Printf("Password provided is correct: %s\n", u)
		w.WriteHeader(401)
		return
	}
	fmt.Printf("Username: %s\n", u)
	fmt.Printf("Password: %s\n", p)
	w.WriteHeader(200)

	jData, err := getJsonResponse()
	if err != nil {
		// handle error
	}
	w.Write(jData)
	return
}

func getJsonResponse()([]byte, error) {
	meta := make(map[string]string)
	meta["host"] = "vdb.host.com"
	meta["url"] = "blah"

	resource := make(map[string]string)
	resource["VDB"] = "TESTVDB"
	resource["timeStamp"] = "Today"

	d := Data{meta, resource}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}
