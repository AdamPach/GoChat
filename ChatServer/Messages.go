package main

import "strings"

func ItsCommand(mesg string) bool {
	return mesg[0] == '/'
}

func RemoveSendingCharacters(mesg string) string {
	return strings.Replace(mesg, "\r\n", "", -1)
}
