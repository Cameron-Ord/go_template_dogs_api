package rendering

import (
	"fmt"
	"html/template"
	"net/http"
)

func DefaultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Setting")
	w.Header().Set("Content-Type", "text/html")
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Go Server-Side Rendering",
		Message: "This is a dynamic message rendered with Go templates!",
	}

	tmpl, err := template.New("index").Parse(`
	<!DOCTYPE html>
        <html>
        <head>
            <title>{{.Title}}</title>
			<link rel="stylesheet" type="text/css" href="/assets/styles.css">
        </head>
        <body>
            <h1 class="title_tag">{{.Title}}</h1>
            <p class="message_tag">{{.Message}}</p>
			<button id="click_me">click me!</button>
			<script type="module" src="js/api_calls.js"></script>
        </body>
        </html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
