package tribserver

import (
	"log"
	"net"
	"time"

	"github.com/dylanfeehan/tribbler/api/tribrpc"
	"google.golang.org/grpc"
)

type tribServer struct {
	Tribs                           map[string]*tribrpc.Tribble
	tribrpc.UnimplementedTribServer // struct embedding?
}

func newTribServerImpl() *tribServer {
	s := &tribServer{
		Tribs: map[string]*tribrpc.Tribble{},
	}
	return s
}

func NewHandle(host, port string) {
	address := net.JoinHostPort(host, port)
	lis, err := net.Listen("tcp", address) // lis, Listener, is essentially a wrapper around a TCP socket.
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, grpc.ConnectionTimeout(120*time.Second))

	grpcServer := grpc.NewServer(serverOptions...) // Server is an interface we can register grpc server implementations to
	// tell the grpc server about our RPC implementation, register it
	tribrpc.RegisterTribServer(grpcServer, newTribServerImpl()) // warpper around grpcServer.RegisterService(trib desc, trib impl)

	grpcServer.Serve(lis)
}
