package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}

type TemplateData struct {
	ResultAscii string
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		renderTemplate(w, "404", nil)
	} else {
		if req.Method == "GET" {
			renderTemplate(w, "index", nil)
		}
	}
}

func AsciiArtHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var data map[string]string             // Create a map to store data
	err = json.Unmarshal(bodyBytes, &data) // Decode JSON data
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text := data["text"]   // Access "text" field from JSON
	style := data["style"] // Access "style" field from JSON
	result := AsciiProg(text, style)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, result)
}
func DownloadHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	filePath := "static/asciiart.txt"
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found.", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file info
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info.", http.StatusInternalServerError)
		return
	}
	Text := req.FormValue("text")
	style := req.FormValue("style")
	result := AsciiProg(Text, style)
	os.WriteFile(filePath, []byte(result), 0777)
	w.Header().Set("Content-Disposition", "attachment; filename=asciiart.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	http.ServeContent(w, req, fileInfo.Name(), fileInfo.ModTime(), file)

}
