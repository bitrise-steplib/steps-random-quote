package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// request
	url := "http://api.icndb.com/jokes/random"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create requestuest:", err)
		os.Exit(1)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// perform request
	client := &http.Client{}
	response, err := client.Do(request)

	var data map[string]interface{}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Println("Json unmarshal failed:", err)
		os.Exit(1)
	}

	if response.StatusCode == 200 {
		value := data["value"]
		valueMap := value.(map[string]interface{})
		fmt.Println("Joke:", valueMap["joke"])

		cmd := fmt.Sprintf("envman add --key RANDOM_QUOTE --value %v", valueMap["joke"])
		c := exec.Command("bash", "-c", cmd)
		err := c.Run()
		if err != nil {
			fmt.Println("Failed to add random quote to envman", err)
			os.Exit(1)
		}
	} else {
		errorMsg := fmt.Sprintf("Status code: %s Body: %s", response.StatusCode, response.Body)
		fmt.Println("Request failed:", errorMsg)

		os.Exit(1)
	}
}
