gen-proto:
	protoc \
		--proto_path=./api \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./api \
		--go-grpc_out=paths=source_relative:./api \
		./api/*.proto