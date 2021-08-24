package main

import (
	"context"
	"fmt"
	"github.com/natenho/golang-examples/protobuf-anypb/common"
	"github.com/natenho/golang-examples/protobuf-anypb/proto/cache"
	"github.com/natenho/golang-examples/protobuf-anypb/proto/custom"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	port = ":50051"
)

type server struct {
	cache.UnimplementedCacherServer
}

func (s *server) Get(context.Context, *cache.GetCacheRequest) (*cache.GetCacheResponse, error) {

	items := make(map[string]*anypb.Any)

	for i := 0; i < 3; i++ {
		items[fmt.Sprintf("A key %v", i)] = common.MakeStringValue(fmt.Sprintf("value %v", i))
	}

	for i := 4; i < 7; i++ {
		items[fmt.Sprintf("B key %v", i)] = common.MakeInt32Value(int32(i))
	}

	for i := 8; i < 10; i++ {
		items[fmt.Sprintf("C key %v", i)] = common.MakeCustomValue(&custom.SearchRequest{
			Query:         "Any given query in response",
			PageNumber:    int32(rand.Int31()),
			ResultPerPage: int32(rand.Int31()),
			Corpus:        custom.SearchRequest_VIDEO,
		})
	}

	return &cache.GetCacheResponse{Items: items}, nil
}

func (s *server) Set(_ context.Context, in *cache.SetCacheRequest) (*emptypb.Empty, error) {
	for k, v := range in.Items {

		log.Printf("Received %v", k)
		m, err := v.UnmarshalNew()
		if err != nil {
			log.Fatalf("Unmarshal new failed: %v", err)
		}

		switch x := m.(type) {
		case *wrapperspb.StringValue:
			log.Printf("Received string! :) %v", x)
		default:
			log.Printf("Received string! :) %v", x)
		}

	}

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cache.RegisterCacherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
