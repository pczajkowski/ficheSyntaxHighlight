package main

import (
	"bufio"
	"strings"
)

var checks map[string][]string

func init() {
	checks = make(map[string][]string)
	checks["c"] = []string{"printf(", "malloc(", "realloc(", "free(", "#include", "#define", "sizeof", "typeof"}
	checks["go"] = []string{"fmt.", "package ", ":= range", "make(map[", "if err != nil"}
	checks["csharp"] = []string{"using", "namespace", ".Where(", ".Select(", "public ", "private ", "readonly ", "List<", " async ", "await "}
	checks["html"] = []string{"html>", "head>", "body>", "title>", "script>", "div>"}
}

func removeString(array []string, index int) []string {
	return append(array[:index], array[index+1:]...)
}

func check(content string, keywords []string) float64 {
	keywordsToCheck := make([]string, len(keywords))
	copy(keywordsToCheck, keywords)
	certainty := 0.0

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		if certainty >= 1.0 || len(keywordsToCheck) == 0 {
			break
		}

		line := scanner.Text()
		for index, keyword := range keywordsToCheck {
			if strings.Contains(line, keyword) {
				certainty += 0.1
				keywordsToCheck = removeString(keywordsToCheck, index)
			}
		}
	}

	return certainty
}

func detectLanguage(content string) string {
	currentLanguage := ""
	currentBest := 0.0

	for key, value := range checks {
		result := check(content, value)
		if result > currentBest {
			currentLanguage = key
			currentBest = result
		}
	}

	return currentLanguage
}
