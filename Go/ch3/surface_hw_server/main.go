package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// обработка ключей
		keys := r.URL.Query()
		cellsKey := keys.Get("cells")
		cells, err := strconv.Atoi(cellsKey)
		if err != nil {
			log.Printf("Ошибка cells: %v", err)
		}
		log.Printf("%d\n", cells)

		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w, cells)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
