package ServerModels

import "bufio"

type ServerCommand struct {
	Name        string
	ExecCommand func(s *Server, reader *bufio.Reader) error
}
