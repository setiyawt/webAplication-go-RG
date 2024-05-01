package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	//client := http.Client{}

	// Hit API https://animechan.xyz/api/quotes/anime?title=naruto with method GET:
	resp, err := http.Get("https://animechan.xyz/api/quotes/anime?title=naruto")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode JSON respon ke []Animechan
	var quotes []Animechan
	err = json.Unmarshal(body, &quotes)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})

	// Send POST request
	resp, err := http.Post("https://postman-echo.com/post", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return Postman{}, err
	}
	defer resp.Body.Close()

	// baca response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Postman{}, err
	}

	var postman Postman
	err = json.Unmarshal(body, &postman)
	if err != nil {
		return Postman{}, err
	}

	return postman, nil
	// Hit API https://postman-echo.com/post with method POST:
	// TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
