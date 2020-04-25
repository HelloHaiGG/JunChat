package servers

import (
	"github.com/gorilla/websocket"
	"net"
	"net/http"
)

func Start(port string) error {
	listen, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		return err
	}

	chatConn, err := listen.Accept()
	if err != nil {
		return err
	}
}



