package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"news-webapp/news"
)

var htmlTemplate = template.Must(template.ParseFiles("templates/index.html"))

type Search struct {
	Query      string
	NextPage   int
	TotalPages int
	Results    *news.Results
}

func (s *Search) PreviousPage() int {
	return s.CurrentPage() - 1
}

func (s *Search) CurrentPage() int {
	if s.NextPage == 1 {
		return s.NextPage
	}

	return s.NextPage - 1
}

func (s *Search) IsLastPage() bool {
	return s.NextPage >= s.TotalPages
}

func livenessHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Container is alive!")
}

func readinessHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Application is ready!")
}

func indexHandler(writer http.ResponseWriter, _ *http.Request) {
	buf := &bytes.Buffer{}
	err := htmlTemplate.Execute(buf, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(writer)
}

func searchHandler(apiClient *news.Client) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		u, err := url.Parse(request.URL.String())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		searchQuery := params.Get("q")
		page := params.Get("page")
		if page == "" {
			page = "1"
		}
		results, err := apiClient.FetchEverything(searchQuery, page)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		nextPage, err := strconv.Atoi(page)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      searchQuery,
			NextPage:   nextPage,
			TotalPages: int(math.Ceil(float64(results.TotalResults) / float64(apiClient.PageSize))),
			Results:    results,
		}

		if ok := !search.IsLastPage(); ok {
			search.NextPage++
		}

		buf := &bytes.Buffer{}
		err = htmlTemplate.Execute(buf, search)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(writer)
	}
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiKey, err := os.ReadFile("/secrets/news_api_key.secret.example")
	if err != nil {
		log.Fatal("ApiKey must be set")
		os.Exit(1)
	}

	myClient := &http.Client{Timeout: 10 * time.Second}
	apiClient := news.NewClient(myClient, string(apiKey), 20)

	// Initializing the webserver.
	mux := http.NewServeMux()

	// Serving static files.
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler(apiClient))
	mux.HandleFunc("/health", livenessHandler) // Endpoint for liveness probe
 mux.HandleFunc("/ready", readinessHandler) // Endpoint for readiness probe

	http.ListenAndServe(":"+port, mux)
}
