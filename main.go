package ApiClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
// Error Type

type EmptyResult struct {
	Name string
}
func (e *EmptyResult) Error() string { return e.Name + ": No Result" }
var ErrorEmpty *EmptyResult


type Client struct{
	Version 	int
	Token 		string
	Url 		string
}

func NewClient(token string) *Client {
	return &Client{Version: 1,Token: token,Url: "https://newapi.critsend.com"}
}
func (c *Client) Get(param string) (interface{},error){
	client := http.Client{}
	req , err := http.NewRequest("GET", "https://newapi.critsend.com/"+param, nil)
	if err != nil {
		//Handle Error
	}
	req.Header = http.Header{
		"Authorization": []string{c.Token},
	}
	response, err := client.Do(req)
	//Check if response work
	if err != nil {
		fmt.Printf("%+v\n", err)
		return "", ErrorEmpty
	}
	responseData, err := ioutil.ReadAll(response.Body)
	var resObject interface{}
	json.Unmarshal(responseData, &resObject)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resObject)
	return resObject,nil
}

func  (c *Client) CheckResponse(resp interface{}) error{
	//Return Empty if no value

	return nil

}