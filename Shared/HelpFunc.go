package Shared

import (
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
