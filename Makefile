generate:
	mkdir -p internal/film
	protoc --go_out=internal/film --go_opt=paths=source_relative \
           api/film_v1/film.proto \
           --go-grpc_out=internal/film --go-grpc_opt=paths=source_relative \
           api/film_v1/film.proto