package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

// ReposAPI - ссылка на API
const ReposAPI = "https://api.github.com/repos/"

// IssuesURL - адрес API для работы с поиском issues
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult - структура для хранения результата выполненного поиска
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issues
}

// Issues - структура каждой темы
type Issues struct {
	Number    int
	User      *User
	Title     string
	Body      string
	State     string
	CreatedAt time.Time `json:"created_at"`
}

// User - структура хранения инфы о пользователе
type User struct {
	Login string
}

// SearchAnIssues - делает http.Get запрос к API issues Github,
// записывает результат в структуру IssuesSearchResult
// и возвращает *IssuesSearchResult или error
func SearchAnIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAnIssues делает запрос к API и возвращает срез всех Issues
func GetAnIssues(owner string, repo string) ([]Issues, error) {
	q := strings.Join([]string{owner, repo}, "/")
	resp, err := http.Get(ReposAPI + q + "/issues")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Сбой подключения: %v", resp.Status)
	}
	defer resp.Body.Close()

	var result []Issues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateAnIssue создаёт новую тему (Issue)
func CreateAnIssue(owner string, repo string, data map[string]string) (*Issues, error) {
	// преобразование данных в формат JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(jsonData)

	// аутентификация на github
	client := &http.Client{}                                             // подготовка клиента
	username, password := credentials()                                  // получение логина и пароля
	q := ReposAPI + strings.Join([]string{owner, repo}, "/") + "/issues" // подготовка url

	req, err := http.NewRequest("POST", q, buf) // подготовка POST запроса
	req.SetBasicAuth(username, password)        // в запрос будет добавлен авторизационный заголовок

	resp, err := client.Do(req) // клиент делает запрос, получает ответ или ошибку
	if err != nil {

		return nil, fmt.Errorf("Ошибка: %v", err)
	}
	if resp.StatusCode != http.StatusCreated { // проверка статус кода
		return nil, fmt.Errorf("Статус: %s", resp.Status)
	}
	defer resp.Body.Close()

	var result Issues                                                  // определяю переменную для хранения результата
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // расшифровка тела ответа в структуру Issues
		return nil, err
	}
	return &result, nil // возвращаю указатель на результат
}

// Получаю username и password
// Для скрытого ввода пароля используется библиотека golang.org/x/crypto/ssh/terminal
func credentials() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите логин: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err)
	}

	fmt.Print("Введите пароль: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err)
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password)
}
