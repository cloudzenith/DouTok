gen-proto:
	protoc \
		--proto_path=./api \
		--proto_path=./third_party \
		--go_out=paths=source_relative:./api \
		--go-http_out=paths=source_relative:./api \
		--openapi_out=fq_schema_naming=true,default_response=false:. \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:./api \
		./api/imapi/*.proto

gen-db:
	gentool -c "gen.yml"

build:
	docker build -t im-service .
