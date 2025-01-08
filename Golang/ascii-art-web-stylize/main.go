package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

// Create a struct that holds information to be displayed in our HTML file
type Outstrct struct {
	Asciitext []string
	Place     string
}

func main() {
	// Declare struct which later takes info from POST and sends to html
	asciiInfo := Outstrct{}
	
	//Parse index.html
	templates := template.Must(template.ParseFiles("templates/index.html"))

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//This method takes in the URL path "/" and a function that takes in a response writer, and a http request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check for correct URL path
		if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
			http.Error(w, "404 not found.", http.StatusNotFound)
			return
		}
		//Check for POST method and assign values to textInput and banner
		if r.Method == "POST" {
			textInput := r.FormValue("textInput")
			banner := r.FormValue("banner")
			if textInput != "" && banner != "" {
				asciiInfo.Asciitext = toArt(banner, textInput)
				asciiInfo.Place = textInput
			}
			//If errors show an internal server error message
			// Also passing asciiInfo struct to index.html
			if err := templates.ExecuteTemplate(w, "index.html", asciiInfo); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else { // Executes index.html for GET method, sends empty asciiInfo struct
			asciiInfo = Outstrct{}
			if err := templates.ExecuteTemplate(w, "index.html", asciiInfo); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}

	})

	// Print link to site
	fmt.Println("Open http://localhost:8000")
	//Start the web server, set the port to listen to 8000. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	fmt.Println(http.ListenAndServe(":8000", nil))

}

func toArt(banner, word string) []string {
	// Open banner file
	file, err := os.ReadFile("banners/" + banner + ".txt")
	checkErr(err)

	// Remove non ASCII characters
	re := regexp.MustCompile(`[^\s!-~]*`)
	word = re.ReplaceAllString(word, "$1")
	// if input starts and ends with "", remove them
	re = regexp.MustCompile(`^"(.*)"$`)
	word = re.ReplaceAllString(word, "$1")

	// Split word and ascii art characters into slice
	re = regexp.MustCompile(`[[:cntrl:]]`)
	words := re.Split(word, -1)
	// regexp expression is different in Windows OS for all banner files, and thinkertoy in Linux
	if runtime.GOOS == "windows" || banner == "thinkertoy" {
		re = regexp.MustCompile(`\r\n`)
	}
	letters := re.Split(string(file), -1)

	// Declare []string where lines to be printed out will be appended
	art := []string{}
	// Print out the words
	for _, v := range words {
		if len(v) > 0 {
			for i := 0; i < 9; i++ {
				line := ""
				for _, vi := range v {
					line += letters[(int(vi)-32)*9+i]
				}
				art = append(art, line)
			}
		}
	}
	return art
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
