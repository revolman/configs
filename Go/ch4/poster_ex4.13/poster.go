package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// APIKey ...
const APIKey = "b414d24b"

// APIURL ...
const APIURL = "http://www.omdbapi.com/"

// Movie ...
type Movie struct {
	Title      string
	Year       string
	Released   string
	Director   string
	Country    string
	Poster     string
	Ratings    []*Ratings
	IMDBRating string
	Type       string
	Actors     string
	Plot       string
	imdbID     string
}

// Ratings ...
type Ratings struct {
	Source string
	Value  string
}

// SearchingResult ...
type SearchingResult struct {
	Search       []*Search
	totalResults string
	Response     string
}

// Search ...
type Search struct {
	Title  string
	Year   string
	imdbID string
	Type   string
	Poster string
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("Использование:\nВведите название фильма, который ищете")
		os.Exit(1)
	}
	// Определение поискового запроса и запуск поиска
	query := strings.Join(os.Args[1:], " ")
	result, err := searching(query)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	for _, item := range result {
		fmt.Println("Найдено:")
		fmt.Printf("imdbID: %s\t%s\n", )
	}

}

// searching ...
func searching(query string) ([]SearchingResult, error) {
	// q := strings.Join([]string{APIURL, "?apikey=" + APIKey + "&s=" + query}, "/")
	q := url.QueryEscape(query)
	fmt.Println(q)
	resp, err := http.Get(APIURL + "?apikey=" + APIKey + "&s=" + q)
	if err != nil {
		return nil, fmt.Errorf("Не удаётся установить соединение")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result []SearchingResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Ошибка маршалинга: %v", err)
	}

	return result, nil
}
