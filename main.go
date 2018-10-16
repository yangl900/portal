package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type urlRequest struct {
	URL string `json:"url"`
}

type resource struct {
	ID string `json:"id"`
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(bytes)

	if isJSON(input) {
		var res resource
		if err := json.Unmarshal([]byte(input), &res); err != nil {
			return
		}

		if isResID(res.ID) {
			printLink(res.ID)
			tryOpenLink(res.ID)
		}
	} else {
		if shouldUnwrap, str := isJSONString(input); shouldUnwrap {
			input = str
		}

		input = normalize(input)

		if isResID(input) {
			printLink(input)
			tryOpenLink(input)
		} else if isMultiLine(input) {
			lines := strings.Split(input, "\n")
			for _, line := range lines {
				if isResID(line) {
					printLink(line)
				}
			}
		}
	}
}

func normalize(s string) string {
	return strings.Trim(s, "\n ")
}

func isJSONString(s string) (bool, string) {
	var js string
	return json.Unmarshal([]byte(s), &js) == nil, js

}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isResID(id string) bool {
	return strings.HasPrefix(id, "/subscriptions/") && !strings.Contains(id, "\n")
}

func isMultiLine(id string) bool {
	return strings.Contains(id, "\n")
}

func printLink(id string) {
	fmt.Printf(getLink(id))
}

func getLink(id string) string {
	return fmt.Sprintf("https://portal.azure.com/#resource%s\n", id)
}

func tryOpenLink(id string) {
	if termID, ok := os.LookupEnv("ACC_TERM_ID"); ok {
		url := fmt.Sprintf("http://localhost:8888/openLink/%s", termID)

		client := &http.Client{}
		reqBody, _ := json.Marshal(urlRequest{URL: getLink(id)})
		resp, err := client.Post(url, "application/json", bytes.NewReader(reqBody))

		if err != nil {
			fmt.Printf("Failed to open links in Cloud Shell: %s\n", err.Error())
		}

		if resp.StatusCode != 200 {
			content, _ := ioutil.ReadAll(resp.Body)
			fmt.Printf("Failed to open links in Cloud Shell: %s\n%s\n", resp.Status, content)
		}

		fmt.Println("Opening portal in new tab...")
	}
}
