package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/mux"
)

const (
    githubAPI = "https://api.github.com/search/repositories"
)

type GitHubRepo struct {
    FullName        string `json:"full_name"`
    HTMLURL         string `json:"html_url"`
    StargazersCount int    `json:"stargazers_count"`
    ForksCount      int    `json:"forks_count"`
}

type SearchResult struct {
    Items []GitHubRepo `json:"items"`
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", topReposHandler).Methods("GET")
    fmt.Println("Server listening on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func topReposHandler(w http.ResponseWriter, r *http.Request) {
    client := &http.Client{Timeout: 10 * time.Second}
    // Fetch repositories from the last 7 days
    todayDate := time.Now().Format("2006-01-02")
    url := fmt.Sprintf("%s?q=created:>=%s&sort=stars&order=desc&per_page=10", githubAPI, todayDate)


    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        http.Error(w, "Failed to create request: "+err.Error(), http.StatusInternalServerError)
        return
    }

    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Failed to fetch data: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var searchResult SearchResult
    if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
        http.Error(w, "Failed to decode data: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Widget-Title", "Top GitHub Repositories of the Week")
    w.Header().Set("Widget-Content-Type", "html")
    w.Header().Set("Content-Type", "text/html")

    htmlContent := "<ul class='list list-gap-14 collapsible-container' data-collapse-after='5'>"
    for _, repo := range searchResult.Items {
        htmlContent += fmt.Sprintf(`<li>
            <a class='size-h3 color-primary-if-not-visited' href='%s' target='_blank'>%s</a>
            <ul class='list-horizontal-text'>
                <li>Stars: %d</li>
                <li>Forks: %d</li>
            </ul>
        </li>`, repo.HTMLURL, repo.FullName, repo.StargazersCount, repo.ForksCount)
    }
    htmlContent += "</ul>"
    fmt.Fprint(w, htmlContent)
}

