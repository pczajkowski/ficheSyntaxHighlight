package main

import (
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"io/ioutil"
	"log"
	"os"
)

func readFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("Error reading file %s: %s", path, err)
	}

	return string(content), nil
}

func createHTML(content string, file *os.File) error {
	language := detectLanguage(content)
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	style := styles.Get("dracula")
	if style == nil {
		style = styles.Fallback
	}

	formatter := html.New(html.Standalone(true), html.WithLineNumbers(true), html.LinkableLineNumbers(true, ""), html.LineNumbersInTable(true))
	iterator, err := lexer.Tokenise(nil, content)
	if err != nil {
		return fmt.Errorf("Error tokenizing content HTML: %s", err)
	}

	err = formatter.Format(file, style, iterator)
	if err != nil {
		return fmt.Errorf("Error writing HTML: %s", err)
	}

	return nil
}

func convert(files paths) error {
	content, err := readFile(files.textFile)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %s", files.textFile, err)
	}

	file, err := os.Create(files.htmlFile)
	if err != nil {
		return fmt.Errorf("Error creating file: %s", err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Error closing file: %s", err)
		}
	}()

	err = createHTML(content, file)
	if err != nil {
		return err
	}

	return nil
}
