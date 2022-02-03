package ServerModels

import (
	"net"
)

type ChatClient struct {
	Name       string
	Room       *Room
	Connection net.Conn
}
