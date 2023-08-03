generate:
	mkdir -p internal/film
	protoc --go_out=. --go_opt=paths=source_relative \
           api/proto/films.proto \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           api/proto/films.proto