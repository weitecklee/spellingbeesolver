package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func getLetters() (string, error) {
	req, err := http.NewRequest("GET", "https://www.nytimes.com/puzzles/spelling-bee", nil)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v\n", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making request: %v\n", err)

	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Request failed with status: %s\n", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v\n", err)
	}

	re := regexp.MustCompile(`"validLetters":\[(.*?)\]`)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) > 1 {
		letters := regexp.MustCompile(`"(\w)"`).FindAllStringSubmatch(matches[1], -1)
		var result string
		for _, letter := range letters {
			result += letter[1]
		}
		return result, nil
	} else {
		return "", fmt.Errorf("Failed to extract letters\n")
	}
}
