package ApiClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
// Error Type

type EmptyResult struct {
	Name string
}
func (e *EmptyResult) Error() string { return e.Name + ": No Result" }
var ErrorEmpty *EmptyResult

type RespApi struct {
	Error 	interface{}
	Info 	struct{
		QueryId	int
		Limit 	int
		Page	int
		Count	string
	}
	Result	[]interface{}
}


type Client struct{
	Version 	int
	Token 		string
	Url 		string
}

func NewClient(token string) *Client {
	return &Client{Version: 1,Token: token,Url: "https://newapi.critsend.com"}
}
func (c *Client) Get(param string,model interface{}) (interface{},error){
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
	var resObject RespApi
	json.Unmarshal(responseData, &resObject)
	if err != nil {
		return "", ErrorEmpty
	}
	if resObject.Info.Count=="0"{
		return "", ErrorEmpty
	}
	fmt.Printf("%+v\n", resObject.Info)
	return resObject.Result,nil
}

func  (c *Client) CheckResponse(resp interface{}) error{
	//Return Empty if no value



	return nil

}