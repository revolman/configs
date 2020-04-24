package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Comics - структура для хранения информации о публикации
type Comics struct {
	Day        string
	Month      string
	Year       string
	Num        int
	Title      string
	Transcript string
}

const xkcdURL string = "https://xkcd.com"

func main() {
	_, err := os.Stat("comics.json")

	if os.IsNotExist(err) {
		fmt.Println("Архива не обнаружен. Подготовка архива:")

		template, _ := json.MarshalIndent([]Comics{}, "", " ")
		if err := ioutil.WriteFile("comics.json", template, 0644); err != nil {
			log.Fatal("Ошибка при создании шаблона файла: ", err)
		}
		getComs(1, lastOnSite())
		os.Exit(0)
	}

	getComs(lastIndex()+1, lastOnSite())
	fmt.Println("Архив а актуальном состоянии.")

}

func getComs(num int, end int) {
	// var comics []*Comics
	var alreadyInFile []*Comics

	rfile, err := os.OpenFile("comics.json", os.O_RDONLY, 0644) // дескриптор указывает на то что файл открыт только на чтение
	if err != nil {
		log.Fatalf("Открыие файла на чтение: %v\n", err)
	}
	defer rfile.Close()

	wfile, err := os.OpenFile("comics.json", os.O_WRONLY, 0644) // дескриптор указывает на то что файл открыт на перезапись
	if err != nil {
		log.Fatalf("Открыие файла на чтение: %v\n", err)
	}
	defer wfile.Close()

	if err := json.NewDecoder(rfile).Decode(&alreadyInFile); err != nil {
		log.Fatal("Ошибка декодирования:", err)
	}

	// for num := 1; num <= countComs(); num++ {
	for num <= end {
		q := strings.Join([]string{xkcdURL, strconv.Itoa(num), "info.0.json"}, "/")
		num++
		resp, err := http.Get(q)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Статус не ОК")
			continue
		}

		var result Comics
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Fatal("Ошибка декодирования:", err)
		}

		// этот сроез вообще нужен?
		// comics = append(comics, &result)
		fmt.Println(result.Num)

		alreadyInFile = append(alreadyInFile, &result)
	}

	// декодировать содержимое файла в струкруру, добавить новые данные, закодировать обратно

	marshaled, _ := json.MarshalIndent(alreadyInFile, "", " ")
	wfile.WriteString(string(marshaled) + "\n")
}

func lastIndex() int {
	file, err := os.OpenFile("comics.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []Comics
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		log.Fatalf("Ошибка маршалинга: %v\n", err)
	}

	index := result[len(result)-1].Num
	fmt.Printf("lastIndex считает, что номер последнего комикса %d\n", index)

	return index
}

func lastOnSite() int {
	resp, err := http.Get(xkcdURL + "/" + "info.0.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result Comics
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Ошибка маршалинга: %v\n", err)
	}

	return result.Num
}
