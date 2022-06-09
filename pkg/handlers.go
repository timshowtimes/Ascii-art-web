package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	FONTS               = []string{"standard", "shadow", "thinkertoy"}
	Templates, TemplErr = template.ParseGlob("ui/templates/*.html")
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		log.Println("Status: Not Found (404)")
		Templates.ExecuteTemplate(w, "error.html", 404)
		return
	}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}

	Templates.ExecuteTemplate(w, "index.html", nil)
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	var Result string
	r.ParseForm()
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		log.Printf("Status: 405. %v Method is Not Allowed", r.Method)
		Templates.ExecuteTemplate(w, "error.html", http.StatusMethodNotAllowed)
		return
	}
	if TemplErr != nil {
		log.Fatal(TemplErr)
	}
	texts, textOk := r.Form["text"]
	fonts, fontOk := r.Form["fontType"]
	button, buttonOk := r.Form["submit"]
	if !textOk || !fontOk || !buttonOk || !Contains(FONTS, fonts) || IsNotAscii(texts[0]) {
		w.WriteHeader(400)
		log.Println("Status: Bad Request (400)")
		Templates.ExecuteTemplate(w, "error.html", http.StatusBadRequest)
		return
	}
	text := strings.Join(texts, "") // takes value from textfield
	font := strings.Join(fonts, "") // takes value from radio-button

	fontType := "fontstyles/" + font + ".txt" // path to .txt
	filePath, err := ioutil.ReadFile(fontType)
	if err != nil || !DHashSum(filePath) {
		w.WriteHeader(500)
		log.Println("Status: Internal Server Error (500).")
		Templates.ExecuteTemplate(w, "error.html", http.StatusInternalServerError)
		return
	}

	symbol := GetMap(string(filePath))
	text1 := strings.Split(text, "\n")

	for _, val := range text1 {
		Result += GetStr(val, symbol)
	}

	if button[0] == "Submit" {
		Templates.ExecuteTemplate(w, "index.html", Result)
		log.Println("Status: OK (200)")

		fmt.Fprint(w, "<div id='maincontent'><pre>"+Result+"</pre></div>")
	} else if button[0] == "Download" {
		w.Header().Set("Content-Disposition", "attachment; filename=ascii-art.txt")
		w.Header().Set("Content-Type", r.Header.Get("application/x-www-form-urlencoded"))
		w.Write([]byte(Result))
	}
}
