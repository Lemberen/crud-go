package controllers

import (
	vm "crud/controllers/viewmodels"
	"crud/model/db/dao"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const controllerRoute string = "/book"

func init() {
	http.Handle("/", http.FileServer(http.Dir("./views/crud/dist/crud/")))

	http.HandleFunc(createRoute("/delete"), func(w http.ResponseWriter, r *http.Request) {
		authorize(&w)

		if r.Method == "DELETE" {
			var body struct {
				Id int
			}

			_ = json.NewDecoder(r.Body).Decode(&body)

			_ = dao.Delete(body.Id)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
	})

	http.HandleFunc(createRoute("/get"), func(w http.ResponseWriter, r *http.Request) {
		authorize(&w)

		q := dao.Get()
		var result []vm.Book

		for q.Next() {
			var book vm.Book

			err := q.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Language, &book.Pages, &book.Edition, &book.Year, &book.Isbn)
			if err != nil {
				log.Fatal("SCAN FAILED:", err)
			}

			result = append(result, book)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	})

	http.HandleFunc(createRoute("/post"), func(w http.ResponseWriter, r *http.Request) {
		authorize(&w)

		var body vm.Book

		if r.Method == "POST" {
			_ = json.NewDecoder(r.Body).Decode(&body)

			_ = dao.Post(body.Id, body.Title, body.Author, body.Publisher, body.Language, body.Pages, body.Edition, body.Year, body.Isbn)
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "")
	})

	http.HandleFunc(createRoute("/put"), func(w http.ResponseWriter, r *http.Request) {
		authorize(&w)

		var body vm.Book

		if r.Method == "PUT" {
			_ = json.NewDecoder(r.Body).Decode(&body)

			_ = dao.Put(body.Id, body.Title, body.Author, body.Publisher, body.Language, body.Pages, body.Edition, body.Year, body.Isbn)
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
	})
}

func authorize(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

func createRoute(urlRoute string) string {
	return fmt.Sprintf("%s%s", controllerRoute, urlRoute)
}
