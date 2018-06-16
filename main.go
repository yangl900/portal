package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type resource struct {
	ID string `json:"id"`
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Println(err)
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
		}
	} else {
		if shouldUnwrap, str := isJSONString(input); shouldUnwrap {
			input = str
		}

		if isResID(input) {
			printLink(input)
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
	fmt.Printf("https://portal.azure.com/#resource%s\n", id)
}
