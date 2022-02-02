package main

func main() {
	client := InitConnection()

	go client.ListenForMessages()
	client.SendMessages()
}
