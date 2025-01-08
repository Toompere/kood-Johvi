package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob("templates/*.html"))
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	//Check for POST method
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		HandlePostData(r.PostForm)

		if err := templates.ExecuteTemplate(w, "base.html", filterData); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else { // GET method
		filterData = data
		if err := templates.ExecuteTemplate(w, "base.html", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Depending on which variables are recieved from post execute functions
func HandlePostData(pd map[string][]string) {
	if pd["search"] != nil {
		filterData = SearchResult(pd["search"][0])
	} else if pd["cmin"] != nil {
		FilterData(pd["cmin"][0], pd["cmax"][0], pd["famin"][0], pd["famax"][0], pd["location"][0], pd["members"])
	} else if pd["band"] != nil {
		id, _ := strconv.Atoi(pd["band"][0])
		filterData.Details = artists[id-1]
	}
}
