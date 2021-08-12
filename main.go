package ApiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Error Type

type EmptyResult struct {
	Name string
}

func (e *EmptyResult) Error() string { return e.Name + ": No Result" }

var ErrorEmpty *EmptyResult

type RespApi struct {
	Error interface{}
	Info  struct {
		QueryId int
		Limit   int
		Page    int
		Count   string
	}
	Result []map[string]interface{}
}
type OptionApi struct {
	Error   interface{}
	Info    interface{}
	Title   string
	History []struct {
		Code string
		Info string
	}
	Description string

	Parameters interface{}
}

type Client struct {
	Version int
	Token   string
	Url     string
}

func NewClient(token string) *Client {
	return &Client{Version: 1, Token: token, Url: "https://newapi.critsend.com"}
}

func (c *Client) Insert(param string, data string) (i []map[string]string, e error) {
	client := http.Client{}
	b := bytes.NewBuffer([]byte(data))
	req, err := http.NewRequest("POST", "https://newapi.critsend.com/"+param, b)
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
		return i, ErrorEmpty
	}
	responseData, err := ioutil.ReadAll(response.Body)
	var resObject RespApi
	json.Unmarshal(responseData, &resObject)
	if err != nil {
		return i, ErrorEmpty
	}
	if len(resObject.Result) == 0 {
		return i, ErrorEmpty
	}
	for _, row := range resObject.Result {
		dd := make(map[string]string)
		for field, value := range row {
			dd[field] = fmt.Sprint(reflect.ValueOf(value))
		}
		i = append(i, dd)
	}
	return i, nil
}
func (c *Client) Update(param string, id int, data string) (e error) {
	return e
}
func (c *Client) Delete(param string, id int) (e error) {
	return e
}
func (c *Client) Info(param string) (i OptionApi, e error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://newapi.critsend.com/"+param, nil)
	if err != nil {
		//Handle Error
	}
	req.Header = http.Header{
		"Authorization": []string{c.Token},
	}
	response, err := client.Do(req)
	//Check if response work
	if err != nil {
		fmt.Printf("Error")
		return i, ErrorEmpty
	}

	responseData, err := ioutil.ReadAll(response.Body)

	var respOption OptionApi
	json.Unmarshal(responseData, &respOption)
	return respOption, e
}
func (c *Client) Get(param string) (i []map[string]string, e error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://newapi.critsend.com/"+param, nil)
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
		return i, ErrorEmpty
	}
	responseData, err := ioutil.ReadAll(response.Body)
	fmt.Println(responseData)
	var resObject RespApi
	json.Unmarshal(responseData, &resObject)
	if err != nil {
		return i, ErrorEmpty
	}
	if resObject.Info.Count == "0" {
		return i, ErrorEmpty
	}
	//i=resObject.Result
	for _, row := range resObject.Result {
		dd := make(map[string]string)
		for field, value := range row {
			dd[field] = fmt.Sprint(reflect.ValueOf(value))
		}
		i = append(i, dd)
	}
	return i, nil
}

func (c *Client) CheckResponse(resp interface{}) error {
	//Return Empty if no value
	return nil
}
