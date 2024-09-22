gen-proto:
	protoc \
		--proto_path=./api \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./api \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./api \
		./api/*.proto

gen-db:
	gentool -c "gen.yml"

build:
	docker build -t base-service .

build1:
	 docker build --build-arg proxy="https://goproxy.cn" -t base-service .
