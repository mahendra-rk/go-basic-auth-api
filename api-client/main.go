package main



import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}

	url := "http://localhost:8080/example"
	fmt.Println("HTTP JSON POST URL:", url)

	var jsonData = []byte(`{
		"DbName": "OST",
		"TrouxID": "529751",
		"PrimaryVDBOwner": "servicepvnp",
		"SecondaryVDBOwner": "pierre_morel"
	}`)
	request, error := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.SetBasicAuth("abc", "123")


	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

}
