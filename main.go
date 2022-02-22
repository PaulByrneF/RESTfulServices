package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

// Struct for CatFact
type CatFact struct {
	Fact   string
	Length int
}

type UserPicture struct {
	Large     string `json:"large,omitempty"`
	Medium    string `json:"medium,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type UserName struct {
	Title string `json:"title,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

type UserResult struct {
	Name    UserName    `json:"name,omitempty"`
	Email   string      `json:"email,omitempty"`
	Picture UserPicture `json:"picture,omitempty"`
}

type RandomUser struct {
	Results []UserResult `json:"results,omitempty"`
}

// Retrieve Random User
func GetRandomUser() *RandomUser {
	url := "https://randomuser.me/api/?inc=name,email,picture"
	var user RandomUser
	err := GetJson(url, &user)
	if err != nil {
		fmt.Printf("Error getting Json Body: %s\n", err.Error())
	} else {
		fmt.Printf("User: %s %s %s\nEmail: %s\nThumbnail: %s\n",
			user.Results[0].Name.Title,
			user.Results[0].Name.First,
			user.Results[0].Name.Last,
			user.Results[0].Email,
			user.Results[0].Picture.Thumbnail,
		)
	}

	return &user
}

// Retrieves cat fact from catfact.ninja and stores in struct
func GetCatFact() {
	url := "https://catfact.ninja/fact"
	var catFact CatFact
	err := GetJson(url, &catFact)
	if err != nil {
		fmt.Printf("Error getting cat fact: %s\n", err.Error())
		return
	} else {
		fmt.Printf("A super interesting cat fact: %s\n", catFact.Fact)
	}
}

// Decodes json to struct
func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	GetCatFact()
	user := GetRandomUser()
	PrintUserJson(user)
	PrintUser(user)
}

func PrintUser(user *RandomUser) {
	fmt.Println(user)
}

func PrintUserJson(userPointer *RandomUser) {
	jsonStr, err := json.Marshal(userPointer)
	if err != nil {
		fmt.Printf("Error converting json to string", err.Error())
	} else {
		fmt.Printf("Json payload is : %s\n", jsonStr)
	}
}
