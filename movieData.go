package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Put your IMDB API Key in a .env file with a parameter called API_KEY
// or change how you get your API Key to make this func work!
func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	apiKey := os.Getenv("API_KEY")

	return apiKey
}

func getMovieData(movieName string) []byte {
	apiKey := getApiKey()

	url := fmt.Sprintf("https://imdb8.p.rapidapi.com/auto-complete?q=%s", movieName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", "imdb8.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	return body
}

func getMovieName(movieName string) string {
	jsonData := getMovieData(movieName)

	var data struct {
		D []struct {
			MovieName string `json:"l"`
		} `json:"d"`
	}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		panic(err)
	}

	return data.D[0].MovieName
}

func getMoviePoster(movieName string) string {
	jsonData := getMovieData(movieName)

	var data struct {
		D []struct {
			I struct {
				ImageURL string `json:"imageUrl"`
			} `json:"i"`
		} `json:"d"`
	}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		panic(err)
	}

	return data.D[0].I.ImageURL
}

func getMovieYear(movieName string) int {
	jsonData := getMovieData(movieName)

	var data struct {
		D []struct {
			MovieYear int `json:"y"`
		} `json:"d"`
	}

	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		panic(err)
	}

	return data.D[0].MovieYear
}
