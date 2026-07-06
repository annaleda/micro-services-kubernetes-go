package handlers

import (
"html/template"
"net/http"

"frontend/internal/service"
)

func Home(w http.ResponseWriter, r *http.Request) {
data := service.LoadPageData()

tmpl := template.Must(template.ParseFiles("internal/templates/index.html"))
tmpl.Execute(w, data)
}
