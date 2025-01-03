// This is the implementation of a TribClient that we have written for you as
// an example. This code also serves as a good reference for understanding
// how RPC works in Go. DO NOT MODIFY!

package tribclient

import (
	"context"
	"fmt"
	"log"

	"github.com/dylanfeehan/tribbler/api/tribrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type tribHandle struct {
	client tribrpc.TribClient
	ctx    context.Context
	opts   grpc.CallOption
}

func NewTribHandle(serverHost string, serverPort int) (*tribHandle, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	target := fmt.Sprintf("%v:%v", serverHost, serverPort)
	conn, err := grpc.NewClient(target, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	rpcClient := tribrpc.NewTribClient(conn)
	ctx := context.Background()
	var callOptions grpc.CallOption = grpc.EmptyCallOption{}
	return &tribHandle{
		client: rpcClient,
		ctx:    ctx,
		opts:   callOptions,
	}, nil
}

func (th *tribHandle) Close() error {
	return th.Close()
}
