package ServerModels

type ClientCommand struct {
	CommandName string
	MinArgs     int
	ExecFunc    func(client *ChatClient, server *Server, args []string) error
}
