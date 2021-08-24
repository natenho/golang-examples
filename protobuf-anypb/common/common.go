package common

import (
	"log"

	"google.golang.org/protobuf/reflect/protoreflect"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func MakeStringValue(value string) *anypb.Any {
	protoMessage := &wrapperspb.StringValue{Value: value}

	any, err := anypb.New(protoMessage)
	if err != nil {
		log.Fatalf("could not create message: %v", err)
	}
	return any
}

func MakeInt32Value(value int32) *anypb.Any {
	protoMessage := &wrapperspb.Int32Value{Value: value}

	any, err := anypb.New(protoMessage)
	if err != nil {
		log.Fatalf("could not create message %v", err)
	}
	return any
}

func MakeCustomValue(src protoreflect.ProtoMessage) *anypb.Any {
	any, err := anypb.New(src)
	if err != nil {
		log.Fatalf("could not create message %v", err)
	}
	return any
}