package Shared

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func RemoveSendingCharacters(mesg string) string {
	return strings.Replace(mesg, "\r\n", "", -1)
}

func FormatCommand(mesg string) string {
	return strings.ToLower(RemoveSendingCharacters(mesg))
}

func LogError(err *error) {
	if *err != nil {
		fmt.Println(*err)
	}
}

func EnterCommand(reader bufio.Reader, PromptMessage string) (string, error) {
	fmt.Println("For exit enter press enter!")
	fmt.Print(PromptMessage)
	command, error := reader.ReadString('\n')

	if error != nil {
		return "", error
	}

	command = RemoveSendingCharacters(command)

	if command == "" {
		return "", errors.New("Cancle command")
	}

	return command, nil
}
