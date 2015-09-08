package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"./markdownlog"
)

func errorMessageToOutput(msg string) {
	message := "Quote generation failed!\n"
	message = message + "Error message:\n"
	message = message + msg
	markdownlog.ErrorSectionToOutput(message)
}

func successMessageToOutput(msg string) {
	message := "Quote successfully generated!\n"
	message = message + "Quote:\n"
	message = message + msg + "\n"
	markdownlog.SectionToOutput(message)
}

// RunPipedEnvmanAdd ...
func RunPipedEnvmanAdd(key, value string) error {
	args := []string{"add", "-k", key}
	envman := exec.Command("envman", args...)
	envman.Stdin = strings.NewReader(value)
	envman.Stdout = os.Stdout
	envman.Stderr = os.Stderr
	return envman.Run()
}

func main() {
	// init / cleanup the formatted output
	pth := os.Getenv("BITRISE_STEP_FORMATTED_OUTPUT_FILE_PATH")
	markdownlog.Setup(pth)
	err := markdownlog.ClearLogFile()
	if err != nil {
		fmt.Errorf("Failed to clear log file, err: %s", err)
	}

	// request
	urlString := "http://api.icndb.com/jokes/random"

	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		errorMessageToOutput(fmt.Sprintf("Failed to create requestuest, err: %s", err))
		os.Exit(1)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// perform request
	client := &http.Client{}
	response, err := client.Do(request)

	var data map[string]interface{}
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errorMessageToOutput(fmt.Sprintf("Failed to read request body, err: %s", err))
		os.Exit(1)
	}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		errorMessageToOutput(fmt.Sprintf("Json unmarshal failed, err: %s", err))
		os.Exit(1)
	}

	if response.StatusCode == 200 {
		value := data["value"]
		valueMap, isKind := value.(map[string]interface{})
		if isKind == false {
			errorMessageToOutput(fmt.Sprintf("Failed to convert response, err: %s", err))
			os.Exit(1)
		}

		joke := valueMap["joke"].(string)
		joke, err = url.QueryUnescape(joke)
		if err != nil {
			errorMessageToOutput(fmt.Sprintf("Failed to url decode response, err: %s", err))
			os.Exit(1)
		}

		err := RunPipedEnvmanAdd("quote", joke)
		if err != nil {
			errorMessageToOutput(fmt.Sprintf("Failed to add output to envman, err: %s", err))
			os.Exit(1)
		}

		successMessageToOutput(fmt.Sprintf("%v", valueMap["joke"]))
	} else {
		errorMsg := fmt.Sprintf("Status code: %d Body: %s", response.StatusCode, response.Body)
		errorMessageToOutput(errorMsg)

		os.Exit(1)
	}
}
