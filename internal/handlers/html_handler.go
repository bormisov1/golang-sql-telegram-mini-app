// internal/handlers/html_handler.go
package handlers

import (
    "html/template"
    "net/http"
)

type HTMLHandler struct{}

func (h *HTMLHandler) ServeMiniApp(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/miniapp.html")
    if err != nil {
        http.Error(w, "Failed to load mini app", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
