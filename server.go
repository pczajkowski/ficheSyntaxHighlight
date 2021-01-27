package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var host = flag.String("h", "localhost", "Host on which to serve")
var port = flag.String("p", "9090", "Port on which to serve")
var mainFolder = flag.String("f", "", "Absolute path to where files are kept")
var logFilePath = flag.String("l", "", "Path to log file")

type paths struct {
	textFile string
	htmlFile string
}

func getPaths(folder string) (paths, error) {
	destination := filepath.Join(*mainFolder, folder)
	if !strings.HasPrefix(destination, *mainFolder) {
		return paths{}, fmt.Errorf("Wrong destination: %s", destination)
	}

	if _, err := os.Stat(destination); os.IsNotExist(err) {
		return paths{}, fmt.Errorf("Destination folder doesn't exist at %s!", destination)
	}

	textFile := filepath.Join(destination, "index.txt")
	if _, err := os.Stat(textFile); os.IsNotExist(err) {
		return paths{}, fmt.Errorf("Text file doesn't exist at %s!", textFile)
	}

	return paths{textFile, filepath.Join(destination, "index.html")}, nil
}

type params struct {
	folder string
	option string
}

func getParameters(path string) (params, error) {
	if path == "" {
		return params{}, fmt.Errorf("Empty string!")
	}

	parts := strings.Split(path, "/")
	partsCount := len(parts)
	if partsCount > 2 {
		return params{}, fmt.Errorf("Too many parameters: %s", parts)
	}

	if partsCount == 2 {
		return params{parts[0], parts[1]}, nil
	}

	return params{parts[0], ""}, nil
}

func serve(w http.ResponseWriter, r *http.Request) {
	pathFromURL := strings.TrimPrefix(r.URL.Path, "/")
	t := template.Must(template.ParseFiles("./html/result.html"))

	parameters, err := getParameters(pathFromURL)
	if err != nil {
		log.Println(err)
		t.Execute(w, "Bad parameters!")
		return
	}

	paths, err := getPaths(parameters.folder)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, "Not found!")
		return
	}

	if parameters.option == "t" {
		http.ServeFile(w, r, paths.textFile)
		return
	}

	if _, err := os.Stat(paths.htmlFile); os.IsNotExist(err) {
		err := convert(paths)
		if err != nil {
			log.Println(err)
			http.ServeFile(w, r, paths.textFile)
			return
		}
	}

	http.ServeFile(w, r, paths.htmlFile)
}

func setLog() *os.File {
	if *logFilePath == "" {
		return nil
	}

	file, err := os.OpenFile(*logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Can't open logs file: %s", err)
	}

	log.SetOutput(file)
	return file
}

func main() {
	flag.Parse()

	logFile := setLog()
	if logFile != nil {
		defer logFile.Close()
	}

	hostname := *host + ":" + *port
	http.HandleFunc("/", serve)
	err := http.ListenAndServe(hostname, nil)
	if err != nil {
		fmt.Println(err)
	}
}
