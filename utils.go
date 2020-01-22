package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//For config.json
type Configuration struct {
	Port         string
	ReadTimeout  int64
	WriteTimeout int64
	IdleTimeout  int64
	Static       string
	App          string
	SiteName     string
}

var config Configuration
var logger *log.Logger

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

//pass true for log to file, or false for log to stdout
func initLogger(toFile bool) {
	if toFile {
		logname := strings.ToLower(config.App) + ".log"
		file, err := os.OpenFile(logname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open log file", err)
		}
		logger = log.New(file, config.App+":", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		logger = log.New(os.Stdout, config.App+":", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func init() {
	loadConfig()
	initLogger(false)
}

//template helper with fmaps and data, generates html to writer
func generateHTML(writer http.ResponseWriter, data interface{}, fmap template.FuncMap, filenames ...string) {

	var files []string
	filenames = append([]string{"layout.html"}, filenames...) //prepend layout and pass variadic
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s", file))
	}

	templates := template.New("layout")
	templates = templates.Funcs(fmap) //can be nil
	templates = template.Must(templates.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

//example fmap
func formatRfc822(t time.Time) template.HTML {
	return template.HTML(t.Format(time.RFC822))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
