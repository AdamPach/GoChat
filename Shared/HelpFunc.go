package Shared

import "strings"

func RemoveSendingCharacters(mesg string) string {
	return strings.Replace(mesg, "\r\n", "", -1)
}
