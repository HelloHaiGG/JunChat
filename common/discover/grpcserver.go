package common

import (
	"google.golang.org/grpc"
	"net"
)


func (p *RegisterSvr) RunRpcServer(port string, register func(s *grpc.Server)) error {
	listen, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	register(server)

	err = server.Serve(listen)

	return err
}
