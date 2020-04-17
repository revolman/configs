package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) < 1 {
		exitWithUsage()
	}

	option := os.Args[1]
	args := os.Args[2:]

	// Для поиска используется
	if option == "search" {
		fmt.Printf("Поиск тем по критериям: %s\n", strings.Join(args, " "))
		search(args)
		os.Exit(0)
	}

	owner, repo := args[0], args[1]

	switch option {
	// Получить спиок тем по репозиторию
	case "getall":
		if len(os.Args[1:]) < 2 {
			exitWithUsage()
		}
		fmt.Println("Получение списка тем в репозитории.")
		getAll(owner, repo)
	case "create":
		if len(os.Args[1:]) < 2 {
			exitWithUsage()
		}
		fmt.Println("Создание новой темы.")
		create(owner, repo)
	}

}

func exitWithUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n"+
		"search QUERY\n"+
		"getll|create OWNER REPO\n"+
		"(.......) OWNER REPO NUMBER\n")
	os.Exit(1)
}

// getAll - вывод списка тем найденых в репозитории указанного владельца
func getAll(owner string, repo string) {
	issues, err := GetAnIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d тем найдено по запросу revolman configs:\n", len(issues))
	for _, item := range issues {
		fmt.Printf("#%-5d %9.9s %-55.55s %-10.10s\n",
			item.Number, item.User.Login, item.Title,
			item.CreatedAt.Format(time.RFC3339))
	}
}

// search - поиск по issues. Запускается во всех случаях, когда не указано другое действие
// синтаксис поискового запроса в соответствии с API Github
// пример: repo:golan/go is:open json decoder
func search(args []string) {
	result, err := SearchAnIssues(args) // Обрабатывает больше аргументов, чем другие функции
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues: \n", len(result.Items))
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %-55.55s %-10.10s\n",
			item.Number, item.User.Login, item.Title,
			item.CreatedAt.Format(time.RFC3339))
	}
}

//create - создаёт новую тему
func create(owner string, repo string) {
	data := ParseFile()
	issue, err := CreateAnIssue(owner, repo, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новый вопрос создан успешно.")
	fmt.Printf("#%-5d %9.9s %-55.55s %-10.10s\n",
		issue.Number, issue.User.Login, issue.Body,
		issue.CreatedAt.Format(time.RFC3339))

}
