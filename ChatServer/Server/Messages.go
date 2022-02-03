package Server

func ItsCommand(mesg string) bool {
	return mesg[0] == '/'
}
