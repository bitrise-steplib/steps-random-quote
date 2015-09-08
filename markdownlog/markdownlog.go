package markdownlog

import (
	"fmt"
	"os"
	"strings"
)

var pth string

func Setup(logPath string) {
	pth = logPath
}

func ClearLogFile() error {
	if pth != "" {
		err := os.Remove(pth)
		if err != nil {
			return err
		}

		fmt.Println("Log file cleared")
	} else {
		fmt.Errorf("No log path defined!")
	}

	return nil
}

func ErrorMessageToOutput(msg string) {
	lines := strings.Split(msg, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}

	if pth != "" {
		f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Errorf("Failed to open log file:", err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Errorf("Failed to close log file:", err)
			}
		}()

		_, err = f.Write([]byte(msg))
		if err != nil {
			fmt.Errorf("Failed to write log:", err)
		}
	} else {
		fmt.Errorf("No log path defined!")
	}
}

func ErrorSectionToOutput(section string) {
	msg := "\n" + section + "\n"
	ErrorMessageToOutput(msg)
}

func ErrorSectionStartToOutput(section string) {
	msg := section + "\n"
	ErrorMessageToOutput(msg)
}

func MessageToOutput(msg string) {
	lines := strings.Split(msg, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}

	if pth != "" {
		f, err := os.OpenFile(pth, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Errorf("Failed to open log file:", err)
		}
		defer func() {
			err := f.Close()
			if err != nil {
				fmt.Errorf("Failed to close log file:", err)
			}
		}()

		_, err = f.Write([]byte(msg))
		if err != nil {
			fmt.Errorf("Failed to write log:", err)
		}
	} else {
		fmt.Errorf("No log path defined!")
	}
}

func SectionToOutput(section string) {
	msg := "\n" + section + "\n"
	MessageToOutput(msg)
}

func SectionStartToOutput(section string) {
	msg := section + "\n"
	MessageToOutput(msg)
}
