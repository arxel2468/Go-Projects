package main

import (
    "html/template"
    "log"
    "net/http"
    "sync"
)

// Post represents a blog post
type Post struct {
    Title   string
    Content string
}

var posts []Post
var mutex = &sync.Mutex{}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/new", newPostHandler)
    http.HandleFunc("/save", savePostHandler)

    log.Println("🚀 Server starting on http://localhost:8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("❗ Server failed to start:", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    tmpl.Execute(w, posts)
}

func newPostHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/new.html")
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

func savePostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    title := r.FormValue("title")
    content := r.FormValue("content")

    if title == "" || content == "" {
        http.Redirect(w, r, "/new", http.StatusSeeOther)
        return
    }

    mutex.Lock()
    posts = append(posts, Post{Title: title, Content: content})
    mutex.Unlock()

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
