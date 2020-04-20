package main

import (
	"fmt"
	"io/ioutil"
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

	if len(args) < 2 {
		exitWithUsage()
	}

	owner, repo := args[0], args[1]

	switch option {
	// Получить спиок тем по репозиторию
	case "getall":
		if len(args) != 2 {
			exitWithUsage()
		}
		fmt.Println("Получение списка тем в репозитории.")
		getAll(owner, repo)
	// Создать новую тему
	case "create":
		if len(args) != 2 {
			exitWithUsage()
		}
		fmt.Println("Создание новой темы.")
		create(owner, repo)
	// Обновление темы
	case "update":
		if len(args) != 3 {
			exitWithUsage()
		}
		fmt.Println("Обновление темы.")
		number := args[2]
		update(owner, repo, number)
	// Получить одну тему
	case "get":
		if len(args) != 3 {
			exitWithUsage()
		}
		fmt.Printf("Запрос темы #%s\n", args[2])
		number := args[2]
		GetAnIssue(owner, repo, number)
	}
}

func exitWithUsage() {
	fmt.Fprintf(os.Stderr, "Usage:\n"+
		"search QUERY\n"+
		"getll|create OWNER REPO\n"+
		"update|get OWNER REPO NUMBER\n")
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
	const method string = "POST"

	// отпределяю расположение временного файла
	fpath := os.TempDir() + "/github_issues.tmp"
	// записываю во временный файл шаблон создания issue
	if err := ioutil.WriteFile(fpath, []byte("Title: \nBody: "), 0644); err != nil {
		log.Fatalf("Ошибка при создании шаблона файла: %v", err)
	}
	// вызываю эдитор
	Edit("vim", fpath)
	// преобразаую ввод юзера в отображение
	data := ParseFile(fpath)
	query := ReposAPI + strings.Join([]string{owner, repo, "issues"}, "/")

	result, err := UpdateAnIssue(method, query, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Новый вопрос создан успешно.")
	fmt.Printf("#%-5d %9.9s %-55.55s %-10.10s\n",
		result.Number, result.User.Login, result.Title,
		result.CreatedAt.Format(time.RFC3339))
}

func update(owner string, repo string, number string) {
	const method string = "PATCH"

	fpath := os.TempDir() + "/github_issues.tmp"
	issue, err := GetAnIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(fpath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Title: " + issue.Title +
		"\nBody: " + issue.Body +
		"\nState: " + issue.State)
	defer file.Close()
	// вызываю эдитор
	Edit("vim", fpath)
	// преобразаую ввод юзера в отображение
	data := ParseFile(fpath)

	// ####### далее тест #######
	query := ReposAPI + strings.Join([]string{owner, repo, "issues", number}, "/")

	result, err := UpdateAnIssue(method, query, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Тема обновлена")
	fmt.Printf("#%-5d %9.9s %-30.30s %-25.25s %-10.10s\n",
		result.Number, result.User.Login, result.Title, result.State,
		result.CreatedAt.Format(time.RFC3339))
}
