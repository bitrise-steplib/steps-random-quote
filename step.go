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
)

const (
	urlString = "http://api.icndb.com/jokes/random"
)

func fail(format string, v ...interface{}) {
	errorMsg := fmt.Sprintf(format, v...)
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", errorMsg)
	os.Exit(1)
}

func warn(format string, v ...interface{}) {
	errorMsg := fmt.Sprintf(format, v...)
	fmt.Printf("\x1b[33;1m%s\x1b[0m\n", errorMsg)
}

func info(format string, v ...interface{}) {
	fmt.Println()
	errorMsg := fmt.Sprintf(format, v...)
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", errorMsg)
}

func details(format string, v ...interface{}) {
	errorMsg := fmt.Sprintf(format, v...)
	fmt.Printf("  %s\n", errorMsg)
}

func done(format string, v ...interface{}) {
	errorMsg := fmt.Sprintf(format, v...)
	fmt.Printf("  \x1b[32;1m%s\x1b[0m\n", errorMsg)
}

func exportEnvironmentWithEnvman(keyStr, valueStr string) error {
	envman := exec.Command("envman", "add", "--key", keyStr)
	envman.Stdin = strings.NewReader(valueStr)
	envman.Stdout = os.Stdout
	envman.Stderr = os.Stderr
	return envman.Run()
}

func main() {
	// request
	request, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		fail("Failed to create requestuest, err: %s", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// perform request
	info("Getting random quote from: %s", urlString)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fail("Failed to perform request, err: %s", err)
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			warn("Failed to close response body, error: %s", err)
		}
	}()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fail("Failed to read request body, err: %s", err)
	}

	if response.StatusCode == 200 {
		var data map[string]interface{}
		if err = json.Unmarshal(bodyBytes, &data); err != nil {
			fail("Failed to unmarshal (%s), err: %s", string(bodyBytes), err)
		}

		value := data["value"]
		valueMap, isKind := value.(map[string]interface{})
		if isKind == false {
			fail("Failed to convert response: %s", value)
		}

		joke := valueMap["joke"].(string)
		joke, err = url.QueryUnescape(joke)
		if err != nil {
			fail("Failed to url decode response (%s), err: %s", joke, err)
		}

		if err := exportEnvironmentWithEnvman("RANDOM_QUOTE", joke); err != nil {
			fail("Failed to add output to envman, err: %s", err)
		}

		done(joke)
	} else {
		fail("Status code: %d Body: %s", response.StatusCode, response.Body)
	}
}
