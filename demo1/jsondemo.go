package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Users []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}
type Todo []struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder(resp.Body)
	dataStruct := Todo{}
	jsonDecoder.Decode(&dataStruct)
	resp.Body.Close()
	fmt.Println(len(dataStruct))
	//fmt.Println(dataStruct)
	dataStruct[0].Title = "xxxxxx"
	jsonEncoder := json.NewEncoder(os.Stdout)
	jsonEncoder.Encode(dataStruct)

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	/*resp1, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return
	}
	jsonDecoder1 := json.NewDecoder(resp1.Body)
	dataStruct1 := Users{}
	jsonDecoder1.Decode(&dataStruct1)
	//jsonDecoder. Decode(&dataStruct2)
	resp1.Body.Close()
	fmt.Println(len(dataStruct1))
	dataStruct1[0].Name = "TestName"
	jsonEncoder1 := json.NewEncoder(os.Stdout)
	jsonEncoder1.Encode(dataStruct1)*/

}
