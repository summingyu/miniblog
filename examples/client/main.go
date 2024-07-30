package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/summingyu/miniblog/internal/pkg/log"
	pb "github.com/summingyu/miniblog/pkg/proto/miniblog/v1"
)

var (
	addr  = flag.String("addr", "127.0.0.1:8090", "The address to connect to.")
	limit = flag.Int64("limit", 10, "Limit to list users.")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(1024*1024*20)), // 20M
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Minute,
			Timeout:             time.Minute,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatalw("Failed to connect", "err", err)
	}
	defer conn.Close()

	c := pb.NewMiniBlogClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var offsetValue int64 = 0
	r, err := c.ListUser(ctx, &pb.ListUserRequest{Offset: &offsetValue, Limit: limit})
	if err != nil {
		log.Fatalw("could not greet: %v", err)
	}

	fmt.Println("TotalCount:", r.TotalCount)
	for _, u := range r.Users {
		d, _ := json.Marshal(u)
		fmt.Println(string(d))
	}
}
