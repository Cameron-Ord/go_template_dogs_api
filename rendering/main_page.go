package rendering

import (
	"fmt"
	"html/template"
	"net/http"

	"main.go/requests"
)

func RenderBase(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Setting")
	w.Header().Set("Content-Type", "text/html")

	dogs, err := requests.Req_Dogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("html/template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, dogs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
