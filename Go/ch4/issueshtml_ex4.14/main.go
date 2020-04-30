package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var issuesTemplate = template.Must(template.New("issuelist").Parse(`
	<h1>{{.Issues | len}} тем</h1>
	<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>Состояние</th>
		<th>Пользователь</th>
		<th>Заголовок</th>
	</tr>
	{{range .Issues}}
	<tr>
		<td><a href='{{.CacheURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
`))

var issueTemplate = template.Must(template.New("issueinfo").Parse(`
	<h1>{{.Title}}</h1>
	<samll>by <a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></small>
	<p>{{.Body}}</p>
`))

// IssuesCache - структура, в которую сохраняются все полученные темы
// позволяет вытягивать темы по номеру, да создания кэша тем
type IssuesCache struct {
	Issues         []Issue
	IssuesByNumber map[int]Issue
	Comments       []Comment
}

// type IssuesCache struct {
// 	Issues         []Issue
// 	IssuesByNumber map[int]Issue
// 	Comments	[]Comment
// }

func (ic IssuesCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := strings.SplitN(r.URL.Path, "/", -1)
	if len(query) < 3 || query[2] == "" {
		if err := issuesTemplate.Execute(w, ic); err != nil {
			log.Print(err)
		}
		return
	}
	numStr := query[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Номер задан не числом, вот этой фигнёй: '%s'", numStr)))
		if err != nil {
			log.Printf("Не удалось создать запрос для %v: %v", r, err)
		}
		return
	}
	issue, ok := ic.IssuesByNumber[num]
	if !ok {
		_, err := w.Write([]byte(fmt.Sprintf("Не существует темы с номером %d", num)))
		if err != nil {
			log.Printf("Не удалось создать запрос для %v: %v", r, err)
		}
		return
	}
	if err := issueTemplate.Execute(w, issue); err != nil {
		log.Print(err)
	}

}

func getNewCache(owner string, repo string) (ic IssuesCache, err error) {
	issues, err := GetIssues(owner, repo)
	if err != nil {
		log.Fatal(err)
	}
	ic.Issues = issues
	ic.IssuesByNumber = make(map[int]Issue, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
	}

	return
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Обязательно нужно указать аргументы: owner repo")
		os.Exit(0)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	issuesCache, err := getNewCache(owner, repo)
	if err != nil {
		log.Print(err)
	}

	http.Handle("/", issuesCache)                         // вызов типа IssueCache запускает метод ServeHTTP
	log.Fatal(http.ListenAndServe("localhost:8000", nil)) // сервер
}
