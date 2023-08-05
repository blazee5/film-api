generate:
	protoc --go_out=. --go_opt=paths=source_relative \
           api/proto/v1/films.proto \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           api/proto/v1/films.proto