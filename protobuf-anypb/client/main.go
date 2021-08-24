package main

import (
	"context"
	"encoding/json"
	"github.com/natenho/golang-examples/protobuf-anypb/common"
	"github.com/natenho/golang-examples/protobuf-anypb/proto/cache"
	"github.com/natenho/golang-examples/protobuf-anypb/proto/custom"
	"log"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cache.NewCacherClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	items := make(map[string]*anypb.Any)
	items["mystringkey"] = common.MakeStringValue("Hello " + name)
	items["myintkey"] = common.MakeInt32Value(7777)
	items["myobjkey"] = common.MakeCustomValue(&custom.SearchRequest{
		Query:         "Any given query",
		PageNumber:    int32(rand.Uint32()),
		ResultPerPage: int32(rand.Uint32()),
		Corpus:        custom.SearchRequest_VIDEO,
	})

	request := &cache.SetCacheRequest{Items: items}

	_, err = c.Set(ctx, request)
	if err != nil {
		log.Fatalf("could not set cache: %v", err)
	}

	cacheResponse, err := c.Get(ctx, &cache.GetCacheRequest{Key: "teste"})
	if err != nil {
		log.Fatalf("could not get cache: %v", err)
	}

	for k, v := range cacheResponse.Items {
		o, _ := v.UnmarshalNew()
		j, _ := json.Marshal(o)
		log.Printf("%s = %T with value %s = %s", k, v, v.TypeUrl, j)
	}

	ctx.Done()
}
