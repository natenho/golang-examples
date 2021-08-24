run:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/cache/cache.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/custom/custom.proto
	go run server/main.go &
	go run client/main.go "Any given message 1"
	go run client/main.go "Any given message 2"
